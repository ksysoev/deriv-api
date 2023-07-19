package deriv

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/net/websocket"
)

// DerivAPI is the main struct for the DerivAPI client.
type DerivAPI struct {
	Origin         *url.URL            // The origin URL for the DerivAPI server
	Endpoint       *url.URL            // The WebSocket endpoint URL for the DerivAPI server
	AppID          int                 // The app ID for the DerivAPI server
	Lang           string              // The language code (ISO 639-1) for the DerivAPI server
	ws             *websocket.Conn     // The WebSocket connection to the DerivAPI server
	lastRequestID  int64               // The last request ID used for the DerivAPI server
	responseMap    map[int]chan string // A map of request IDs to response channels for the DerivAPI server
	TimeOut        time.Duration       // The timeout duration for the DerivAPI server api calls
	connectionLock sync.Mutex          // A lock for the DerivAPI server connection
}

// ApiReqest is an interface for all API requests.
type ApiReqest interface {
}

// ApiObjectRequest is an interface for all API requests that return an object.
type ApiResponse interface {
	UnmarshalJSON([]byte) error
}

type APIResponseReqID struct {
	ReqID int `json:"req_id"`
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

	if urlEnpoint.Scheme != "wss" && urlEnpoint.Scheme != "ws" {
		return nil, fmt.Errorf("invalid endpoint scheme")
	}

	if appID < 1 {
		return nil, fmt.Errorf("invalid app id")
	}

	if lang == "" || len(lang) != 2 {
		return nil, fmt.Errorf("Invalid language")
	}

	query := urlEnpoint.Query()
	query.Add("app_id", strconv.Itoa(appID))
	query.Add("l", lang)

	urlEnpoint.RawQuery = query.Encode()

	api := DerivAPI{
		Origin:         urlOrigin,
		Endpoint:       urlEnpoint,
		AppID:          appID,
		Lang:           lang,
		lastRequestID:  0,
		responseMap:    make(map[int]chan string),
		TimeOut:        30 * time.Second,
		connectionLock: sync.Mutex{},
	}
	return &api, nil
}

// Connect establishes a WebSocket connection with the Deriv API endpoint.
// Returns an error if it fails to connect to WebSoket server.
func (api *DerivAPI) Connect() error {
	api.connectionLock.Lock()
	defer api.connectionLock.Unlock()

	if api.ws != nil {
		return nil
	}

	ws, err := websocket.Dial(api.Endpoint.String(), "", api.Origin.String())
	if err != nil {
		return err
	}

	api.ws = ws

	respChan := make(chan string)
	go api.handleResponses(ws, respChan)
	go api.requestMapper(respChan)

	return nil
}

// Disconnect closes the websocket connection to the Deriv API
func (api *DerivAPI) Disconnect() {
	api.connectionLock.Lock()
	defer api.connectionLock.Unlock()

	if api.ws == nil {
		return
	}

	for reqID, channel := range api.responseMap {
		close(channel)
		delete(api.responseMap, reqID)
	}

	api.ws.Close()
	api.ws = nil
}

// handleResponses handles the responses from the Deriv API
func (api *DerivAPI) handleResponses(wsConn *websocket.Conn, respChan chan string) {
	defer close(respChan)

	for {
		var msg string
		err := websocket.Message.Receive(wsConn, &msg)

		if err != nil {
			return
		}

		respChan <- msg
	}
}

// requestMapper handles the responses from the Deriv API
func (api *DerivAPI) requestMapper(respChan chan string) {
	for rawResp := range respChan {
		var response APIResponseReqID
		err := json.Unmarshal([]byte(rawResp), &response)
		if err != nil {
			continue
		}
		api.connectionLock.Lock()
		channel, ok := api.responseMap[response.ReqID]
		api.connectionLock.Unlock()

		if ok {
			channel <- rawResp
		}
	}
}

// Send sends a request to the Deriv API and returns a channel that will receive the response
func (api *DerivAPI) Send(reqID int, request ApiReqest) (chan string, error) {
	err := api.Connect()

	if err != nil {
		return nil, err
	}

	msg, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	err = api.sendMessage(msg)

	if err != nil {
		return nil, err
	}

	respChan := api.createRequestChannel(reqID)

	return respChan, nil
}

func (api *DerivAPI) sendMessage(msg []byte) error {
	api.connectionLock.Lock()

	err := websocket.Message.Send(api.ws, msg)
	if err != nil {
		api.connectionLock.Unlock()
		api.Disconnect()
		return err
	}

	api.connectionLock.Unlock()
	return nil
}

// SendRequest sends a request to the Deriv API and returns the response
func (api *DerivAPI) SendRequest(reqID int, request ApiReqest, response ApiResponse) (err error) {
	respChan, err := api.Send(reqID, request)

	if err != nil {
		return err
	}

	defer api.closeRequestChannel(reqID)

	select {
	case <-time.After(api.TimeOut):
		return fmt.Errorf("timeout")
	case responseJSON, ok := <-respChan:
		if !ok {
			return fmt.Errorf("connection closed")
		}

		if err = parseError(responseJSON); err != nil {
			return err
		}
		return response.UnmarshalJSON([]byte(responseJSON))
	}
}

// SubscribeRequest sends a request to the Deriv API and returns a channel that will receive responses
func (api *DerivAPI) SubscribeRequest(reqID int, request ApiReqest) (chan string, error) {
	respChan, err := api.Send(reqID, request)

	if err != nil {
		return nil, err
	}

	return respChan, nil
}

// getNextRequestID returns the next request ID
func (api *DerivAPI) getNextRequestID() int {
	return int(atomic.AddInt64(&api.lastRequestID, 1))
}

// closeRequestChannel closes the channel that receives the response for a request
func (api *DerivAPI) closeRequestChannel(reqID int) {
	api.connectionLock.Lock()
	defer api.connectionLock.Unlock()
	channel, ok := api.responseMap[reqID]

	if ok {
		close(channel)
		delete(api.responseMap, reqID)
	}
}

// createRequestChannel creates a channel that receives the response for a request
func (api *DerivAPI) createRequestChannel(reqID int) chan string {
	api.connectionLock.Lock()
	defer api.connectionLock.Unlock()

	respChan := make(chan string)
	api.responseMap[reqID] = respChan

	return respChan
}
