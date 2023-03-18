package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	api, err := NewDerivAPI("wss://ws.binaryws.com/websockets/v3", 1, "en", "https://www.binary.com")

	defer api.Disconnect()

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		timeRequest := TimeRequest{Time: 1}
		err = api.SendRequest(timeRequest)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(2 * time.Second)
	}
}
