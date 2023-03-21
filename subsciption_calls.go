package deriv

func (api *DerivAPI) SubscribeTicks(request Ticks) (*Subsciption[TicksResp], error) {
	reqID := api.getNextRequestID()
	var subscibeFlag TicksSubscribe
	subscibeFlag = 1
	request.ReqId = &reqID
	request.Subscribe = &subscibeFlag

	sub := NewSubcription[TicksResp](api)

	err := sub.Start(reqID, request)

	if err != nil {
		return nil, err
	}

	return sub, nil
}
