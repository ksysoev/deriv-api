package deriv

import (
	"log"
)

func (api *DerivAPI) SubscribeTicks(symbol string) (chan TicksResponse, error) {
	reqID := api.getNextRequestID()
	var subscibe TicksRequestSubscribe
	subscibe = 1
	request := TicksRequest{Ticks: symbol, ReqId: &reqID, Subscribe: &subscibe}

	inChan, err := api.Send(reqID, request)

	if err != nil {
		return nil, err
	}

	initResponse := <-inChan

	if err = ParseError(initResponse); err != nil {
		close(inChan)
		delete(api.responseMap, reqID)
		return nil, err
	}

	var response TicksResponse
	err = response.UnmarshalJSON([]byte(initResponse))
	if err != nil {
		close(inChan)
		delete(api.responseMap, reqID)
		return nil, err
	}

	outChan := make(chan TicksResponse, 1)
	outChan <- response

	go func(inChan chan string, outChan chan TicksResponse) {
		defer close(outChan)
		for rawResponse := range inChan {
			err := ParseError(rawResponse)
			if err != nil {
				log.Printf("Error in subsciption message: %v", err)
				continue
			}

			var response TicksResponse
			err = response.UnmarshalJSON([]byte(rawResponse))
			if err != nil {
				log.Printf("Error in subsciption message: %v", err)
				continue
			}
			outChan <- response
		}
	}(inChan, outChan)

	return outChan, nil
}
