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
		if n.Match(ble.GET_HARDWARE_INFO_COMMAND_ID, ble.TLV_RESPONSE_SUCCESS) {
			log.Printf("%s", n.Payload())
		}
		return nil
	})

	log.Println("sending get-hardware-info-request")
	if _, err := adapter.Write(ble.CmdRequest, ble.GET_HARDWARE_INFO); err != nil {
		log.Fatalf("failed to send get hardware info request: %v\n", err)
	}
	wg.Wait()
	log.Println("done")
}
