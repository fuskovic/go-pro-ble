package main

import (
	"log"
	"sync"

	"github.com/fuskovic/go-pro-ble"
)

func main() {
	adapter, err := ble.NewAdapter()
	if err != nil {
		log.Fatalf("failed to init ble adapter: %v\n", err)
	}
	defer adapter.Close()

	var wg sync.WaitGroup
	wg.Add(1)
	go adapter.HandleNotifications(func(c ble.Characteristic, b []byte) error {
		if len(b) >= 3 {
			defer wg.Done()
			// Second byte is the command ID. Third byte is the status.
			// https://gopro.github.io/OpenGoPro/tutorials/parse-ble-responses#responses-with-payload
			cmdID, status := b[1], b[2]
			if cmdID == ble.WIFI_AP_TOGGLE_COMMAND_ID.Byte() {
				log.Println("received response from wifi-access-point-toggle")
				if status == byte(ble.TLV_RESPONSE_SUCCESS) {
					log.Println("successfully enabled wifi-access-point")
				} else {
					log.Println("failed to enable wifi-access-point")
				}
			}
		}
		return nil
	})

	log.Println("enabling wifi-access-point")
	if _, err := adapter.Write(ble.CmdRequest, ble.WIFI_AP_CONTROL_ENABLE); err != nil {
		log.Fatalf("failed to enable wifi access point: %v\n", err)
	}

	wg.Wait()

	wifiSsid, err := adapter.ReadString(ble.WifiApSsid)
	if err != nil {
		log.Fatalf("failed to read wifi ssid: %v\n", err)
	}

	wifiPw, err := adapter.ReadString(ble.WifiApPassword)
	if err != nil {
		log.Fatalf("failed to read wifi password: %v\n", err)
	}

	log.Println("you can now connect to your GoPro's wifi-access-point using the following credentials")
	log.Printf("ssid: %s\n", wifiSsid)
	log.Printf("password: %s\n", wifiPw)
}
