package ble

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"slices"
	"strings"
	"time"

	"github.com/fuskovic/go-pro-ble/internal/packet"
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
	rawNotifications chan *rawNotification
	device           *bluetooth.Device
	characteristics  map[Characteristic]*bluetooth.DeviceCharacteristic
}

// NewAdapter initializes a new bluetooth interface for reading and writing to various
// ble service characteristics as well as listening for notifications.
func NewAdapter() (Adapter, error) {
	tinyGoAdapter := bluetooth.DefaultAdapter
	if err := tinyGoAdapter.Enable(); err != nil {
		return nil, fmt.Errorf("failed to enable BLE stack: %v", err)
	}
	log.Println("BLE stack successfully enabled")

	ch := make(chan bluetooth.ScanResult, 1)
	log.Println("scanning for devices...")
	err := tinyGoAdapter.Scan(func(adapter *bluetooth.Adapter, result bluetooth.ScanResult) {
		if strings.Contains(result.LocalName(), "GoPro") {
			log.Printf("found %s\n", result.LocalName())
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
	log.Printf("connected to %s[%s]\n", scanResult.LocalName(), scanResult.Address.String())

	a := &adapter{
		rawNotifications: make(chan *rawNotification),
		characteristics:  make(map[Characteristic]*bluetooth.DeviceCharacteristic),
		device:           &d,
	}

	log.Println("discovering services + characteristics")
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
			log.Printf("failed to discover service characteristics: %v\n", err)
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
			log.Println("- service", s.Name())
			log.Printf("-- characteristic #%d: %s[%s]\n", i+1, c.name, uuid)
			log.Println("    readable=", c.readable)
			log.Println("    writable=", c.writeable)
			log.Println("    notifiable=", c.notifiable)

			if c.notifiable {
				notificationHandler := func(buf []byte) {
					a.rawNotifications <- &rawNotification{
						buf: bytes.NewBuffer(buf),
					}
				}
				if err := char.EnableNotifications(notificationHandler); err != nil {
					log.Printf("failed to handle notification for %s characteristic: %v\n", c.name, err)
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
		timer        = time.NewTimer(time.Second)
		firstPkt     []byte
		notification *humanReadableNotification
	)
	for {
		select {
		case <-timer.C:
			offset := notification.payload.Offset()
			notification.cmdID = COMMAND_ID(firstPkt[0+offset])
			notification.status = TLV_RESPONSE_STATUS(firstPkt[1+offset])
			return handler(notification)
		case rn := <-a.rawNotifications:
			b := rn.buf.Bytes()
			if notification == nil {
				firstPkt = b
				notification = &humanReadableNotification{
					payload: new(packet.Payload),
				}
			}
			notification.payload.Accumulate(b)
		}
	}
}

func (a *adapter) Close() error {
	return a.device.Disconnect()
}
