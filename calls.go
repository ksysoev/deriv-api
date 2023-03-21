package deriv

func (api *DerivAPI) Time() (TimeResponse, error) {
	var response TimeResponse

	reqID := api.getNextRequestID()

	request := TimeRequest{Time: 1, ReqId: &reqID}

	err := api.SendRequest(reqID, request, &response)

	return response, err
}

func (api *DerivAPI) Forget(subscriptionID string) (ForgetResponse, error) {
	var response ForgetResponse

	reqID := api.getNextRequestID()

	request := ForgetRequest{Forget: subscriptionID, ReqId: &reqID}

	err := api.SendRequest(reqID, request, &response)

	return response, err
}

func (api *DerivAPI) Authorize(apiToken string) (AuthorizeResponse, error) {
	var response AuthorizeResponse

	reqID := api.getNextRequestID()

	request := AuthorizeRequest{Authorize: apiToken, ReqId: &reqID}

	err := api.SendRequest(reqID, request, &response)

	return response, err
}
