package main

import (
	"fmt"
	"time"
)

func main() {
	api, err := NewDerivAPI("wss://ws.binaryws.com/websockets/v3", 1, "en", "https://www.binary.com")

	defer api.Disconnect()

	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := api.SendTime()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp, err)
	return

	respChan, err := api.SubscribeTicks("R_51")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		select {
		case resp := <-respChan:
			fmt.Println(resp)
		case <-time.After(5 * time.Second):
			fmt.Println("Timeout")
			return
		}
	}
}
