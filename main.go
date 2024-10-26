package main

import (
	"log"
	"time"

	"github.com/fuskovic/go-pro-sdk/internal/ble"
)

func main() {
	adapter, err := ble.NewAdapter()
	if err != nil {
		log.Fatal(err)
	}
	defer adapter.Close()

	log.Println("done")
	time.Sleep(time.Hour)
}
