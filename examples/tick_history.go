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

	var startTime deriv.TicksHistoryAdjustStartTime = 1
	start := 1
	resp, sub, err := api.SubscribeTicksHistory(deriv.TicksHistory{
		TicksHistory:    "R_50",
		AdjustStartTime: &startTime,
		End:             "latest",
		Start:           &start,
		Style:           "ticks",
		Count:           10})

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
