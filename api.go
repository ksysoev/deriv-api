package deriv

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"

	"golang.org/x/net/websocket"
)

type DerivAPI struct {
	Origin        *url.URL
	Endpoint      *url.URL
	AppID         int
	Lang          string
	ws            *websocket.Conn
	lastRequestID int
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

func (api *DerivAPI) SendRequest(reqID int, request ApiReqest, response ApiObjectResponse) (err error) {

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

	respChan := make(chan string)

	api.responseMap[reqID] = respChan
	defer delete(api.responseMap, reqID)

	err = websocket.Message.Send(api.ws, requestJSON)
	if err != nil {
		return err
	}

	responseJSON := <-respChan

	err = ParseError(responseJSON)

	if err != nil {
		return err
	}

	err = response.UnmarshalJSON([]byte(responseJSON))

	return err
}

func (api *DerivAPI) SubscribeRequest(reqID int, request ApiReqest) (chan string, error) {
	respChan := make(chan string)
	if api.ws == nil {
		err := api.Connect()

		if err != nil {
			return respChan, err
		}
	}

	requestJSON, err := json.Marshal(request)
	if err != nil {
		return respChan, err
	}

	api.responseMap[reqID] = respChan

	err = websocket.Message.Send(api.ws, requestJSON)
	if err != nil {
		defer delete(api.responseMap, reqID)
		return respChan, err
	}

	return respChan, nil
}

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
			fmt.Println(responseString)
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

func (api *DerivAPI) getNextRequestID() int {
	api.lastRequestID++
	return api.lastRequestID
}
