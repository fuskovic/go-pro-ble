package ble

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"slices"
	"strings"
	"time"

	"tinygo.org/x/bluetooth"
)

var (
	// ErrPermissionDenied is returned when an operation is attempted on a ble
	// characteristic that does not have support for the operation.
	ErrPermissionDenied = errors.New("permission denied")

	// ErrCharacteristicNotFound is returned when the target characteristic
	// is not supported by the ble adapter.
	ErrCharacteristicNotFound = errors.New("characteristic not found")
)

// Adapter is a bluetooth interface that supports reading and writing to
// ble service characteristics.
type Adapter interface {
	Write(Characteristic, []byte) (int, error)
	Read(Characteristic, []byte) (int, error)
	ReadString(Characteristic) (string, error)
	HandleNotifications(func(Notification) error) error
	Close() error
}

type adapter struct {
	pkts            chan []byte
	device          *bluetooth.Device
	characteristics map[Characteristic]*bluetooth.DeviceCharacteristic
	log             *slog.Logger
}

type AdapterConfig struct {
	// Toggle debug logging.
	Debug bool
}

// NewAdapter initializes a new bluetooth interface for reading and writing to various
// ble service characteristics as well as listening for notifications.
func NewAdapter(config *AdapterConfig) (Adapter, error) {
	a := &adapter{
		pkts:            make(chan []byte),
		characteristics: make(map[Characteristic]*bluetooth.DeviceCharacteristic),
		log:             newLogger(config.Debug),
	}

	tinyGoAdapter := bluetooth.DefaultAdapter
	if err := tinyGoAdapter.Enable(); err != nil {
		return nil, fmt.Errorf("failed to enable BLE stack: %v", err)
	}
	a.log.Info("BLE stack successfully enabled")

	ch := make(chan bluetooth.ScanResult, 1)
	a.log.Info("scanning for devices...")
	err := tinyGoAdapter.Scan(func(adapter *bluetooth.Adapter, result bluetooth.ScanResult) {
		if strings.Contains(result.LocalName(), "GoPro") {
			a.log.Debug("target device found", "name", result.LocalName())
			adapter.StopScan()
			ch <- result
		}
	})
	if err != nil {
		return nil, fmt.Errorf("scan error: %v", err)
	}

	// block until scan is complete
	scanResult := <-ch

	d, err := tinyGoAdapter.Connect(scanResult.Address, bluetooth.ConnectionParams{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s", scanResult.LocalName())
	}
	a.log.Info("connected",
		"local-name", scanResult.LocalName(),
		"address", scanResult.Address.String(),
	)
	a.device = &d

	a.log.Debug("discovering services + characteristics")
	srvcs, err := a.device.DiscoverServices(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to descover services: %s", err)
	}

	for _, srvc := range srvcs {
		var s Service
		if !slices.ContainsFunc(Services, func(svc Service) bool {
			if svc.uuid == srvc.UUID().String() {
				s = svc
				return true
			}
			return false
		}) {
			continue
		}

		chars, err := srvc.DiscoverCharacteristics(nil)
		if err != nil {
			a.log.Error("failed to discover service characteristics", "error", err)
			continue
		}

		for i, char := range chars {
			uuid := char.UUID().String()
			var c Characteristic
			if !slices.ContainsFunc(Characterstics, func(characteristic Characteristic) bool {
				if characteristic.uuid == uuid {
					c = characteristic
					return true
				}
				return false
			}) {
				continue
			}

			a.characteristics[c] = &char
			a.log.Debug("discovered",
				"service", s.Name(),
				"characteristic-number", i+1,
				"characteristic", c.name,
				"readable", c.readable,
				"writable", c.writeable,
				"notifiable", c.notifiable,
			)

			if c.notifiable {
				if err := char.EnableNotifications(func(b []byte) { a.pkts <- b }); err != nil {
					a.log.Error("failed to handle notification",
						"characteristic-name", c.name,
						"error", err,
					)
				}
			}
		}
	}
	return a, nil
}

func (a *adapter) Write(c Characteristic, b []byte) (int, error) {
	if !c.writeable {
		return -1, ErrPermissionDenied
	}
	char, ok := a.characteristics[c]
	if !ok {
		return -1, ErrCharacteristicNotFound
	}
	return char.Write(b)
}

func (a *adapter) Read(c Characteristic, b []byte) (int, error) {
	if !c.readable {
		return -1, ErrPermissionDenied
	}
	char, ok := a.characteristics[c]
	if !ok {
		return -1, ErrCharacteristicNotFound
	}
	return char.Read(b)
}

func (a *adapter) ReadString(c Characteristic) (string, error) {
	b := make([]byte, 255)
	n, err := a.Read(c, b)
	if err != nil {
		return "", fmt.Errorf("failed to read characteristic: %v", err)
	}
	log.Printf("read %d bytes", n)
	return string(b[:n]), nil
}

func (a *adapter) HandleNotifications(handler func(Notification) error) error {
	var (
		timer    = time.NewTimer(time.Second)
		firstPkt []byte
		n        *notification
	)
	for {
		select {
		case <-timer.C:
			n.cmdID = COMMAND_ID(firstPkt[0+n.payload.offset])
			n.status = TLV_RESPONSE_STATUS(firstPkt[1+n.payload.offset])
			return handler(n)
		case pkt := <-a.pkts:
			if n == nil {
				firstPkt = pkt
				n = &notification{payload: new(payload)}
			}
			n.payload.Accumulate(pkt)
		}
	}
}

func (a *adapter) Close() error {
	return a.device.Disconnect()
}

func newLogger(debug bool) *slog.Logger {
	logLvl := slog.LevelInfo
	if debug {
		logLvl = slog.LevelDebug
	}
	return slog.New(slog.NewTextHandler(os.Stdout,
		&slog.HandlerOptions{
			Level: logLvl,
		},
	))
}
