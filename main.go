package main

import (
	"log"
	"time"

	"github.com/fuskovic/go-pro-sdk/internal/ble"
)

func main() {
	adapter, err := ble.NewAdapter()
	if err != nil {
		log.Printf("failed to init ble adapter: %v\n", err)
		return
	}
	defer adapter.Close()

	wifiSsid, err := adapter.ReadString(ble.WifiApSsidUuid)
	if err != nil {
		log.Printf("failed to read wifi ssid: %v\n", err)
		return
	}

	wifiPw, err := adapter.ReadString(ble.WifiApPasswordUuid)
	if err != nil {
		log.Printf("failed to read wifi password: %v\n", err)
		return
	}

	// wifiPower, err := adapter.ReadString(ble.WifiApPowerUuid)
	// if err != nil {
	// 	log.Printf("failed to read wifi power: %v\n", err)
	// 	return
	// }

	wifiState, err := adapter.ReadString(ble.WifiApStateUuid)
	if err != nil {
		log.Printf("failed to read wifi state: %v\n", err)
		return
	}

	log.Printf("ssid: %s\n", wifiSsid)
	log.Printf("password: %s\n", wifiPw)
	// log.Printf("power: %s\n", wifiPower)
	log.Printf("state:  %s\n", wifiState)

	go adapter.HandleNotifications(func(c ble.Characteristic, b []byte) error {
		log.Printf("%q", string(b))
		// TODO: decode into human readable formt 
		// 2024/10/26 16:19:09 received notification from command-response
		// 2024/10/26 16:19:09 "\x02\x17\x00"
		return nil
	})

	n, err := adapter.Write(ble.CmdRequestUuid, ble.WifiApControlEnable)
	if err != nil {
		log.Printf("failed to enable wifi access point: %v\n", err)
		return
	}
	log.Printf("wrote %d bytes", n)


	log.Println("done")
	time.Sleep(time.Hour)
}
