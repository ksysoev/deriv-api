package deriv

func (api *DerivAPI) SubscribeTicks(symbol string) (*Subsciption[TicksResponse], error) {
	reqID := api.getNextRequestID()
	var subscibe TicksRequestSubscribe
	subscibe = 1
	request := TicksRequest{Ticks: symbol, ReqId: &reqID, Subscribe: &subscibe}

	sub := NewSubcription[TicksResponse](api)

	err := sub.Start(reqID, request)

	if err != nil {
		return nil, err
	}

	return sub, nil
}
