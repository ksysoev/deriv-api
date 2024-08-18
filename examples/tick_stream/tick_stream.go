package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ksysoev/deriv-api"
	"github.com/ksysoev/deriv-api/schema"
)

func main() {
	ctx := context.Background()

	api, err := deriv.NewDerivAPI("wss://ws.derivws.com/websockets/v3", 36544, "en", "https://localhost/", deriv.Debug)

	if err != nil {
		log.Fatal(err)
	}

	defer api.Disconnect()

	resp, sub, err := api.SubscribeTicks(ctx, schema.Ticks{Ticks: "R_50"})

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Symbol: ", *resp.Tick.Symbol, "Quote: ", *resp.Tick.Quote)

	for tick := range sub.Stream {
		fmt.Println("Symbol: ", *tick.Tick.Symbol, "Quote: ", *tick.Tick.Quote)
	}
}
