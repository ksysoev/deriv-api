package deriv

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"sync/atomic"

	"golang.org/x/net/websocket"
)

type DerivAPI struct {
	Origin        *url.URL
	Endpoint      *url.URL
	AppID         int
	Lang          string
	ws            *websocket.Conn
	lastRequestID int64
	responseMap   map[int]chan string
}

type ApiReqest interface {
}
type ApiObjectResponse interface {
	UnmarshalJSON([]byte) error
}

type APIResponse struct {
	ReqID   int    `json:"req_id"`
	MsgType string `json:"msg_type"`
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
		Origin:        urlOrigin,
		Endpoint:      urlEnpoint,
		AppID:         appID,
		Lang:          lang,
		lastRequestID: 0,
		responseMap:   make(map[int]chan string),
	}
	return &api, nil
}

// Connect establishes a WebSocket connection with the Deriv API endpoint.
// Returns an error if the dialing process fails.
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

		var response APIResponse

		err = json.Unmarshal([]byte(msg), &response)
		if err != nil {
			log.Fatal(err)
		}

		channel, ok := api.responseMap[response.ReqID]

		if ok {
			channel <- msg
		}
	}
}

func (api *DerivAPI) Send(reqID int, request ApiReqest) (chan string, error) {
	var err error

	if api.ws == nil {
		err = api.Connect()

		if err != nil {
			return nil, err
		}
	}

	msg, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	respChan := make(chan string)
	api.responseMap[reqID] = respChan

	err = websocket.Message.Send(api.ws, msg)
	if err != nil {
		delete(api.responseMap, reqID)
		close(respChan)
		// TODO: Disconnect from the API. Need to handle this error properly
		return nil, err
	}

	return respChan, nil
}

// SendRequest sends a request to the Deriv API and returns the response
func (api *DerivAPI) SendRequest(reqID int, request ApiReqest, response ApiObjectResponse) (err error) {
	respChan, err := api.Send(reqID, request)

	if err != nil {
		return err
	}

	defer close(respChan)
	defer delete(api.responseMap, reqID)

	responseJSON := <-respChan

	if err = ParseError(responseJSON); err != nil {
		return err
	}

	return response.UnmarshalJSON([]byte(responseJSON))
}

// SubscribeRequest sends a request to the Deriv API and returns a channel that will receive responses
func (api *DerivAPI) SubscribeRequest(reqID int, request ApiReqest) (chan string, error) {
	respChan, err := api.Send(reqID, request)

	if err != nil {
		return nil, err
	}

	return respChan, nil
}

func (api *DerivAPI) getNextRequestID() int {
	return int(atomic.AddInt64(&api.lastRequestID, 1))
}
