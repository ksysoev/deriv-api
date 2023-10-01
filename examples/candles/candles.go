package main

import (
	"fmt"
	"log"

	"github.com/ksysoev/deriv-api"
	"github.com/ksysoev/deriv-api/schema"
)

func main() {
	api, err := deriv.NewDerivAPI("wss://ws.binaryws.com/websockets/v3", 1089, "en", "https://localhost/")

	if err != nil {
		log.Fatal(err)
	}

	defer api.Disconnect()

	var startTime schema.TicksHistoryAdjustStartTime = 1
	start := 1
	resp, sub, err := api.SubscribeCandlesHistory(schema.TicksHistory{
		TicksHistory:    "R_50",
		AdjustStartTime: &startTime,
		End:             "latest",
		Start:           &start,
		Count:           10})

	if err != nil {
		log.Fatal(err)
		return
	}

	for _, candle := range resp.Candles {
		fmt.Println(
			"Symbol: ", "R_50",
			"Open: ", *candle.Open,
			"Close: ", *candle.Close,
			"High: ", *candle.High,
			"Low: ", *candle.Low,
			"Epoch: ", *candle.Epoch)
	}

	for tick := range sub.Stream {
		fmt.Println(
			"Symbol: ", *tick.Ohlc.Symbol,
			"Open: ", *tick.Ohlc.Open,
			"Close: ", *tick.Ohlc.Close,
			"High: ", *tick.Ohlc.High,
			"Low: ", *tick.Ohlc.Low,
			"Epoch: ", *tick.Ohlc.Epoch)
	}
}
