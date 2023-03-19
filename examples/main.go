package main

import (
	"fmt"
	"time"

	"github.com/ksysoev/deriv-api"
)

func main() {
	api, err := deriv.NewDerivAPI("wss://ws.binaryws.com/websockets/v3", 1, "en", "https://www.binary.com")

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

	fmt.Println(*resp.Time)

	respChan, err := api.SubscribeTicks("R_50")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		select {
		case resp := <-respChan:
			fmt.Println(*resp.Tick.Symbol, *resp.Tick.Quote)
		case <-time.After(5 * time.Second):
			fmt.Println("Timeout")
			return
		}
	}
}
