package deriv

import (
	"log"
)

func (api *DerivAPI) SubscribeTicks(symbol string) (chan TicksResponse, error) {
	reqID := api.getNextRequestID()
	var subscibe TicksRequestSubscribe
	subscibe = 1
	request := TicksRequest{Ticks: symbol, ReqId: &reqID, Subscribe: &subscibe}

	respChan, err := api.SubscribeRequest(reqID, request)

	parsedChan := make(chan TicksResponse)

	go func(outChan chan TicksResponse, inChan chan string) {
		for {
			responseString := <-inChan
			var response TicksResponse

			err := ParseError(responseString)
			if err != nil {
				log.Fatal(err)
				return
			}

			err = response.UnmarshalJSON([]byte(responseString))
			if err != nil {
				log.Fatal(err)
				return
			}
			outChan <- response
		}
	}(parsedChan, respChan)

	return parsedChan, err
}
