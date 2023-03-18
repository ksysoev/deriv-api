package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"

	"golang.org/x/net/websocket"
)

type DerivAPI struct {
	Origin   *url.URL
	Endpoint *url.URL
	AppID    int
	Lang     string
	ws       *websocket.Conn
}

type ApiReqest interface {
}

// NewDerivAPI creates a new instance of DerivAPI by parsing and validating the given
// endpoint URL, appID, language, and origin URL. It returns a pointer to a DerivAPI object
// or an error if any of the validation checks fail.
//
// Parameters:
// - endpoint: string - The WebSocket endpoint URL for the DerivAPI server
// - appID: int - The app ID for the DerivAPI server
// - lang: string - The language code (ISO 639-1) for the DerivAPI server
// - origin: string - The origin URL for the DerivAPI server
//
// Returns:
//   - *DerivAPI: A pointer to a new instance of DerivAPI with the validated endpoint, appID,
//     language, and origin values.
//   - error: An error if any of the validation checks fail.
//
// Example:
//
//	api, err := NewDerivAPI("wss://trade.deriv.com/websockets/v3", 12345, "en", "https://myapp.com")
//	if err != nil {
//		log.Fatal(err)
//	}
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

// Connect connects to the Deriv API
func (api *DerivAPI) Connect() error {
	if api.ws != nil {
		return nil
	}

	ws, err := websocket.Dial(api.Endpoint.String(), "", api.Origin.String())
	if err != nil {
		return err
	}

	api.ws = ws
	go api.handleResponses()

	return nil
}

// Disconnect closes the websocket connection to the Deriv API
func (api *DerivAPI) Disconnect() {
	if api.ws == nil {
		return
	}
	api.ws.Close()
	api.ws = nil
}

// handleResponses handles the responses from the Deriv API
func (api *DerivAPI) handleResponses() {
	for {
		if api.ws == nil {
			return
		}

		var msg string
		err := websocket.Message.Receive(api.ws, &msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(msg)
	}
}

func (api *DerivAPI) SendRequest(request ApiReqest) (err error) {

	if api.ws == nil {
		err = api.Connect()

		if err != nil {
			return err
		}
	}

	requestJSON, err := json.Marshal(request)
	if err != nil {
		return err
	}

	err = websocket.Message.Send(api.ws, requestJSON)
	return err
}
