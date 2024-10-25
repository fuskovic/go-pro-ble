package main

import (
	"log"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/fuskovic/go-pro-sdk/internal/ble"
	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

func main() {
	log.Println("enabling BLE stack...")
	if err := adapter.Enable(); err != nil {
		log.Fatalf("failed to enable BLE stack: %v\n", err)
	}
	log.Println("BLE stack successfully enabled")

	ch := make(chan bluetooth.ScanResult, 1)
	log.Println("scanning for devices...")
	err := adapter.Scan(func(adapter *bluetooth.Adapter, result bluetooth.ScanResult) {
		if strings.Contains(result.LocalName(), "GoPro") {
			log.Printf("found %s\n", result.LocalName())
			adapter.StopScan()
			ch <- result
		}
	})
	if err != nil {
		log.Fatalf("scan error: %v\n", err)
	}

	var device bluetooth.Device
	scanResult := <-ch

	device, err = adapter.Connect(scanResult.Address, bluetooth.ConnectionParams{})
	if err != nil {
		log.Fatalf("failed to connect to %s\n", scanResult.LocalName())
	}
	defer device.Disconnect()
	log.Printf("connected to %s[%s]\n", scanResult.LocalName(), scanResult.Address.String())

	// get services
	log.Println("discovering services/characteristics")
	srvcs, err := device.DiscoverServices(nil)
	if err != nil {
		log.Printf("failed to descover services: %s\n", err)
		return
	}
	log.Printf("discovered %d services\n", len(srvcs))

	// buffer to retrieve characteristic data
	buf := make([]byte, 255)
	for _, srvc := range srvcs {
		s := ble.Service(srvc.UUID().String())

		chars, err := srvc.DiscoverCharacteristics(nil)
		if err != nil {
			log.Printf("failed to discover service characteristics: %s\n", err)
			continue
		}
		if slices.Contains(ble.Services, s) {
			log.Println("- service", s.Name())
		}

		for i, char := range chars {
			uuid := char.UUID().String()
			c := ble.Characteristic(uuid)
			if !slices.Contains(ble.Characterstics, c) {
				continue
			}

			log.Printf("-- characteristic #%d: %s[%s]\n", i+1, c.Name(), uuid)
			log.Println("    readable=", c.Readable())
			log.Println("    writable=", c.Writeable())
			log.Println("    notifiable=", c.Notifiable())

			if c.Readable() {
				n, err := char.Read(buf)
				if err != nil {
					if strings.Contains(err.Error(), "Reading is not permitted.") {
						log.Printf("reading not permitted for char #%d with uuid: %s\n", i+1, char.UUID().String())
						continue
					}
					log.Println("    read error: ", err.Error())
					continue
				}
				log.Println("    data-bytes=", strconv.Itoa(n))
				log.Println("    value=", string(buf[:n]))
			}

			if c.Notifiable() {
				continue
				var notificationHandler func(buf []byte)
				switch c {
				case ble.NetworkMgmtRespUuid:
					notificationHandler = func(buf []byte) {
						// TODO: handle network mgmt response
					}
				case ble.CmdResponseUuid:
					notificationHandler = func(buf []byte) {
						// TODO: handle network mgmt response
					}
				case ble.SettingsRespUuid:
					notificationHandler = func(buf []byte) {
						// TODO: handle settings response
					}
				case ble.QueryRespUuid:
					notificationHandler = func(buf []byte) {
						// TODO: handle query response
					}
				default:
					log.Printf("no notification handler registered for %s[%s]; skipping\n", c.Name(), uuid)
					continue
				}
				if err := char.EnableNotifications(notificationHandler); err != nil {
					log.Printf("failed to handle notification: %v\n", err)
				}
			}
		}
	}

	if err := device.Disconnect(); err != nil {
		log.Println(err)
		return
	}

	log.Println("done")
	time.Sleep(time.Hour)
}
