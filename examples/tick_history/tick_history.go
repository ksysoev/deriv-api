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
	api, err := deriv.NewDerivAPI("wss://ws.derivws.com/websockets/v3", 1089, "en", "https://localhost/")

	if err != nil {
		log.Fatal(err)
	}

	defer api.Disconnect()

	var startTime schema.TicksHistoryAdjustStartTime = 1
	start := 1
	resp, sub, err := api.SubscribeTicksHistory(ctx, schema.TicksHistory{
		TicksHistory:    "R_50",
		AdjustStartTime: &startTime,
		End:             "latest",
		Start:           &start,
		Style:           "ticks",
		Count:           5000})

	if err != nil {
		log.Fatal(err)
		return
	}

	for _, price := range resp.History.Prices {
		fmt.Println("Symbol: ", "R_50", "Quote: ", price)
	}

	for tick := range sub.Stream {
		fmt.Println("Symbol: ", *tick.Tick.Symbol, "Quote: ", *tick.Tick.Quote)
	}
}
