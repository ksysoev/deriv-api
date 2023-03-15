package main

import (
	"fmt"
	"net/url"
	"strconv"

	"golang.org/x/net/websocket"
)

type DerivAPI struct {
	Origin   *url.URL
	Endpoint *url.URL
	AppID    int
	Lang     string
}

func NewDerivAPI(endpoint string, appID int, lang string, origin string) (*DerivAPI, error) {
	urlEnpoint, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	urlOrigin, err := url.Parse(origin)
	if err != nil {
		return nil, err
	}

	if urlEnpoint.Scheme != "wss" {
		return nil, fmt.Errorf("Invalid endpoint scheme")
	}

	if appID < 1 {
		return nil, fmt.Errorf("Invalid app id")
	}

	if lang == "" || len(lang) != 2 {
		return nil, fmt.Errorf("Invalid language")
	}

	query := urlEnpoint.Query()
	query.Add("app_id", strconv.Itoa(appID))
	query.Add("l", lang)

	urlEnpoint.RawQuery = query.Encode()

	api := DerivAPI{
		Origin:   urlOrigin,
		Endpoint: urlEnpoint,
		AppID:    appID,
		Lang:     lang,
	}
	return &api, nil
}

func (api *DerivAPI) Connect() (ws *websocket.Conn, err error) {
	ws, err = websocket.Dial(api.Endpoint.String(), "", api.Origin.String())
	if err != nil {
		return nil, err
	}
	return ws, nil
}

func (api *DerivAPI) Disconnect() {

}
