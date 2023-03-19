package deriv

func (api *DerivAPI) SendTime() (TimeResponse, error) {
	var response TimeResponse

	reqID := api.getNextRequestID()

	request := TimeRequest{Time: 1, ReqId: &reqID}

	err := api.SendRequest(reqID, request, &response)

	return response, err
}

func (api *DerivAPI) SendAuthorize(apiToken string) (AuthorizeResponse, error) {
	var response AuthorizeResponse

	reqID := api.getNextRequestID()

	request := AuthorizeRequest{Authorize: apiToken, ReqId: &reqID}

	err := api.SendRequest(reqID, request, &response)

	return response, err
}
