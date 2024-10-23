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
		} else {
			log.Println("- service", srvc.UUID().String())
		}
		log.Printf("	%d characteristics\n", len(chars))

		for i, char := range chars {
			// mtu, err := char.GetMTU()
			// if err != nil {
			// 	log.Println("    mtu: error:", err.Error())
			// 	continue
			// }

			n, err := char.Read(buf)
			if err != nil {
				if strings.Contains(err.Error(), "Reading is not permitted.") {
					log.Printf("reading not permitted for char #%d with uuid: %s\n", i+1, char.UUID().String())
					continue
				}
				log.Println("    read error: ", err.Error())
				continue
			}

			c := ble.Characteristic(char.UUID().String())
			if slices.Contains(ble.Characterstics, c) {
				log.Printf("-- characteristic #%d: %s\n", i+1, c.Name())
				log.Println("    data bytes", strconv.Itoa(n))
				log.Println("    value =", string(buf[:n]))
				log.Println("    notifiable ", c.Notifiable())
			} else {
				log.Printf("-- characteristic #%d: %s\n", i+1, c.String())
				log.Println("    data bytes", strconv.Itoa(n))
				log.Println("    value =", string(buf[:n]))
				log.Println("    notifiable =", c.Notifiable())
			}

			// err = char.EnableNotifications(func(b []byte) {

			// })
			// if err != nil {
			// 	// handle err
			// }

		}
	}

	if err := device.Disconnect(); err != nil {
		log.Println(err)
		return
	}

	log.Println("done")
	time.Sleep(time.Hour)
}
