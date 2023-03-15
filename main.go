package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	api, err := NewDerivAPI("wss://ws.binaryws.com/websockets/v3", 1, "en", "https://www.binary.com")

	if err != nil {
		fmt.Println(err)
		return
	}

	ws, err := api.Connect()

	if err != nil {

		fmt.Println("Fail to connect:", api.Endpoint.String(), err)
		return
	}

	err = websocket.Message.Send(ws, "{\"ticks\":\"R_50\",\"subscribe\":1}")
	if err != nil {
		log.Fatal(err)
	}

	go func(ws *websocket.Conn) {
		for {
			var msg string
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(msg)
		}
	}(ws)

	for {
		err = websocket.Message.Send(ws, "{\"time\": 1}")
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(5 * time.Second)

	}

	defer ws.Close()
}
