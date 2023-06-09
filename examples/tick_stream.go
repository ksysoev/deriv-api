package main

import (
	"fmt"
	"log"

	"github.com/ksysoev/deriv-api"
)

func main() {
	api, err := deriv.NewDerivAPI("wss://ws.binaryws.com/websockets/v3", 1, "en", "https://www.binary.com")

	if err != nil {
		log.Fatal(err)
	}

	defer api.Disconnect()

	resp, sub, err := api.SubscribeTicks(deriv.Ticks{Ticks: "R_50"})

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Symbol: ", *resp.Tick.Symbol, "Quote: ", *resp.Tick.Quote)

	for tick := range sub.Stream {
		fmt.Println("Symbol: ", *tick.Tick.Symbol, "Quote: ", *tick.Tick.Quote)
	}
}
