package main

import (
	"log"
	"time"

	"github.com/fuskovic/go-pro-sdk/internal/ble"
)

func main() {
	adapter, err := ble.NewAdapter()
	if err != nil {
		log.Printf("failed to init ble adapter: %\n", err)
		return
	}
	defer adapter.Close()

	wifiSsid, err := adapter.ReadString(ble.WifiApSsidUuid)
	if err != nil {
		log.Printf("failed to read wifi ssid: %\n", err)
		return
	}

	wifiPw, err := adapter.ReadString(ble.WifiApPasswordUuid)
	if err != nil {
		log.Printf("failed to read wifi password: %\n", err)
		return
	}

	log.Printf("ssid: %s\n", wifiSsid)
	log.Printf("password: %s\n", wifiPw)
	log.Println("done")
	time.Sleep(time.Hour)
}
