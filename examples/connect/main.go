package main

import (
	"log"
	"sync"

	ble "github.com/fuskovic/go-pro-ble"
)

func main() {
	adapter, err := ble.NewAdapter(&ble.AdapterConfig{Debug: true})
	if err != nil {
		log.Fatalf("failed to init ble adapter: %v\n", err)
	}
	defer adapter.Close()

	var wg sync.WaitGroup
	wg.Add(1)
	go adapter.HandleNotifications(func(n ble.Notification) error {
		defer wg.Done()
		if n.Match(ble.WIFI_AP_TOGGLE_COMMAND_ID, ble.TLV_RESPONSE_SUCCESS) {
			log.Println("successfully enabled wifi-access-point")
			return nil
		}
		return nil
	})

	log.Println("enabling wifi-access-point")
	if _, err := adapter.Write(ble.CmdRequest, ble.WIFI_AP_CONTROL_ENABLE); err != nil {
		log.Fatalf("failed to enable wifi access point: %v\n", err)
	}

	wg.Wait()

	wifiSsid, err := adapter.GetCharacteristicValue(ble.WifiApSsid)
	if err != nil {
		log.Fatalf("failed to read wifi ssid: %v\n", err)
	}

	wifiPw, err := adapter.GetCharacteristicValue(ble.WifiApPassword)
	if err != nil {
		log.Fatalf("failed to read wifi password: %v\n", err)
	}

	log.Println("you can now connect to your GoPro's wifi-access-point using the following credentials")
	log.Printf("ssid: %s\n", wifiSsid)
	log.Printf("password: %s\n", wifiPw)

	if err := ble.ConnectToWifiAccessPoint(wifiSsid, wifiPw); err != nil {
		log.Fatalf("failed to connect to wifi access point: %v\n", err)
	}
	log.Printf("connected to %s\n", wifiSsid)
}
