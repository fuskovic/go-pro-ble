package ble

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"slices"
	"strings"

	"tinygo.org/x/bluetooth"
)

var (
	ErrPermissionDenied       = errors.New("permission denied")
	ErrCharacteristicNotFound = errors.New("characteristic not found")
)

type Adapter interface {
	Write(Characteristic, []byte) (int, error)
	Read(Characteristic, []byte) (int, error)
	Close() error
}

type notification struct {
	*bytes.Buffer
	characteristic Characteristic
}

type adapter struct {
	notifications   chan *notification
	device          *bluetooth.Device
	characteristics map[Characteristic]*bluetooth.DeviceCharacteristic
}

func (a *adapter) Write(c Characteristic, b []byte) (int, error) {
	if !c.Writeable() {
		return -1, ErrPermissionDenied
	}
	char, ok := a.characteristics[c]
	if !ok {
		return -1, ErrCharacteristicNotFound
	}
	return char.Write(b)
}

func (a *adapter) Read(c Characteristic, b []byte) (int, error) {
	if !c.Readable() {
		return -1, ErrPermissionDenied
	}
	char, ok := a.characteristics[c]
	if !ok {
		return -1, ErrCharacteristicNotFound
	}
	return char.Read(b)
}

func (a *adapter) Close() error {
	return a.device.Disconnect()
}

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
		notifications:   make(chan *notification),
		characteristics: make(map[Characteristic]*bluetooth.DeviceCharacteristic),
		device:          &d,
	}

	log.Println("discovering services + characteristics")
	srvcs, err := a.device.DiscoverServices(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to descover services: %s", err)
	}
	log.Printf("discovered %d services\n", len(srvcs))

	for _, srvc := range srvcs {
		s := Service(srvc.UUID().String())

		chars, err := srvc.DiscoverCharacteristics(nil)
		if err != nil {
			log.Printf("failed to discover service characteristics: %s\n", err)
			continue
		}

		if slices.Contains(Services, s) {
			log.Println("- service", s.Name())
		}

		for i, char := range chars {
			uuid := char.UUID().String()
			c := Characteristic(uuid)
			if !slices.Contains(Characterstics, c) {
				continue
			}

			a.characteristics[c] = &char
			log.Printf("provisioned %s characeristic", c.Name())
			log.Printf("-- characteristic #%d: %s[%s]\n", i+1, c.Name(), uuid)
			log.Println("    readable=", c.Readable())
			log.Println("    writable=", c.Writeable())
			log.Println("    notifiable=", c.Notifiable())

			if c.Notifiable() {
				notificationHandler := func(buf []byte) {
					a.notifications <- &notification{
						Buffer:         bytes.NewBuffer(buf),
						characteristic: c,
					}
				}
				if err := char.EnableNotifications(notificationHandler); err != nil {
					log.Printf("failed to handle notification for %s characteristic: %v\n", c.Name(), err)
				}
			}
		}
	}
	return a, nil
}
