package main

import (
	"log"
	"sync"

	ble "github.com/fuskovic/go-pro-ble"
)

func main() {
	adapter, err := ble.NewAdapter()
	if err != nil {
		log.Fatalf("failed to init ble adapter: %v\n", err)
	}
	defer adapter.Close()

	var wg sync.WaitGroup
	wg.Add(1)
	go adapter.HandleNotifications(func(n ble.Notification) error {
		defer wg.Done()
		log.Println("reached handle notifications")
		log.Printf("command-id: %v\n", n.CommandID().Byte())
		log.Printf("status: %s\n", n.Status())
		log.Printf("payload: %s\n", n.Payload().Bytes())
		return nil
	})

	log.Println("sending get-hardware-info-request")
	n, err := adapter.Write(ble.CmdRequest, ble.GET_HARDWARE_INFO)
	if err != nil {
		log.Fatalf("failed to send get hardware info request: %v\n", err)
	}
	log.Printf("wrote %d bytes\n", n)
	wg.Wait()
	log.Println("done")
}
