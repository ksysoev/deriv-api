package deriv

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/coder/websocket"
	"github.com/ksysoev/deriv-api/schema"
)

// DerivAPI is the main struct for the DerivAPI client.
type DerivAPI struct {
	Origin                *url.URL        // The origin URL for the DerivAPI server
	Endpoint              *url.URL        // The WebSocket endpoint URL for the DerivAPI server
	AppID                 int             // The app ID for the DerivAPI server
	Lang                  string          // The language code (ISO 639-1) for the DerivAPI server
	ws                    *websocket.Conn // The WebSocket connection to the DerivAPI server
	lastRequestID         int64           // The last request ID used for the DerivAPI server
	TimeOut               time.Duration   // The timeout duration for the DerivAPI server api calls
	connectionLock        sync.Mutex      // A lock for the DerivAPI server connection
	reqChan               chan ApiReqest  // A channel for sending requests to the DerivAPI server
	closingChan           chan int        // A channel for closing the DerivAPI server connection
	keepAlive             bool            // A flag to keep the connection alive
	keepAliveOnDisconnect chan bool       // A channel to keep the connection alive
	keepAliveInterval     time.Duration   // The interval to send ping requests
	debugEnabled          bool            // A flag to print debug messages
}

// ApiReqest is an interface for all API requests.
type ApiReqest struct {
	id       int
	msg      []byte
	respChan chan []byte
}

// ApiObjectRequest is an interface for all API requests that return an object.
type ApiResponse interface {
	UnmarshalJSON([]byte) error
}

type APIResponseReqID struct {
	ReqID int `json:"req_id"`
}

type DerivApiOption func(api *DerivAPI)

// NewDerivAPI creates a new instance of DerivAPI by parsing and validating the given
// endpoint URL, appID, language, and origin URL. It returns a pointer to a DerivAPI object
// or an error if any of the validation checks fail.
//
// Parameters:
// - endpoint: string - The WebSocket endpoint URL for the DerivAPI server
// - appID: int - The app ID for the DerivAPI server
// - lang: string - The language code (ISO 639-1) for the DerivAPI server
// - origin: string - The origin URL for the DerivAPI server
// - opts: DerivApiOption - A variadic list of DerivApiOption functions to configure the DerivAPI object (optional)
//   - KeepAlive: A DerivApiOption function to keep the connection alive by sending ping requests.
//   - Debug: A DerivApiOption function to enable debug messages.
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
func NewDerivAPI(endpoint string, appID int, lang string, origin string, opts ...DerivApiOption) (*DerivAPI, error) {
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
		Origin:            urlOrigin,
		Endpoint:          urlEnpoint,
		AppID:             appID,
		Lang:              lang,
		lastRequestID:     0,
		TimeOut:           30 * time.Second,
		connectionLock:    sync.Mutex{},
		closingChan:       make(chan int),
		keepAliveInterval: 25 * time.Second,
	}

	for _, opt := range opts {
		opt(&api)
	}

	return &api, nil
}

// KeepAlive option which keeps the connection alive by sending ping requests.
// By default the websocket connection is closed after 30 seconds of inactivity.
func KeepAlive(api *DerivAPI) {
	api.keepAlive = true
}

// Debug option which enables debug messages.
func Debug(api *DerivAPI) {
	api.debugEnabled = true
}

// logDebugf prints a debug message if debug is enabled.
func (api *DerivAPI) logDebugf(format string, v ...interface{}) {
	if api.debugEnabled {
		log.Printf(format, v...)
	}
}

// Connect establishes a WebSocket connection with the Deriv API endpoint.
// Returns an error if it fails to connect to WebSoket server.
func (api *DerivAPI) Connect() error {
	api.connectionLock.Lock()
	defer api.connectionLock.Unlock()

	if api.ws != nil {
		return nil
	}

	api.logDebugf("Connecting to %s", api.Endpoint.String())

	ws, resp, err := websocket.Dial(context.TODO(), api.Endpoint.String(), nil)
	if err != nil {
		api.logDebugf("Failed to establish WS connection: %s", err.Error())
		return err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	api.logDebugf("Connected to %s", api.Endpoint.String())

	api.ws = ws

	api.reqChan = make(chan ApiReqest)
	respChan := make(chan []byte)
	outputChan := make(chan []byte)

	go api.handleResponses(ws, respChan)
	go api.requestSender(ws, outputChan)
	go api.requestMapper(respChan, outputChan, api.reqChan, api.closingChan)

	if api.keepAlive {
		api.keepAliveOnDisconnect = make(chan bool, 1)

		go func(interval time.Duration, onDisconnect chan bool) {
			for {
				select {
				case <-time.After(interval):
					_, err := api.Ping(schema.Ping{Ping: 1})
					if err != nil {
						return
					}
				case <-onDisconnect:
					return
				}

			}
		}(api.keepAliveInterval, api.keepAliveOnDisconnect)
	}

	return nil
}

// Disconnect gracefully disconnects from the Deriv API WebSocket server
// and cleans up any resources associated with the connection.
// This function should only be called after
// the WebSocket connection has been established using the Connect method.
func (api *DerivAPI) Disconnect() {
	api.connectionLock.Lock()
	defer api.connectionLock.Unlock()

	if api.ws == nil {
		return
	}

	api.logDebugf("Disconnecting from %s", api.Endpoint.String())

	close(api.reqChan)

	if api.keepAlive {
		api.keepAliveOnDisconnect <- true
		close(api.keepAliveOnDisconnect)
	}

	api.ws.Close(websocket.StatusNormalClosure, "disconnecting")
	api.ws = nil
}

// requestSender sends requests to the Deriv API WebSocket server over the provided WebSocket connection.
// It reads requests from the reqChan channel and sends them using the websocket.Message.Send method.
// If an error occurs while sending a request, it calls the Disconnect method to gracefully disconnect from the WebSocket server.
func (api *DerivAPI) requestSender(wsConn *websocket.Conn, reqChan chan []byte) {
	for req := range reqChan {
		api.logDebugf("Sending request: %s", req)

		err := wsConn.Write(context.TODO(), websocket.MessageText, req)

		if err != nil {
			api.logDebugf("Failed to send request: %s", err.Error())
			api.Disconnect()
			return
		}
	}
}

// handleResponses reads responses from the Deriv API WebSocket server over the provided WebSocket connection.
// It reads responses using the websocket.Message.Receive method and sends them to the respChan channel.
// If an error occurs while receiving a response, it calls the Disconnect method to gracefully disconnect from the WebSocket server.
// The function returns when the WebSocket connection is closed or when an error occurs.
func (api *DerivAPI) handleResponses(wsConn *websocket.Conn, respChan chan []byte) {
	defer func() {
		close(respChan)

		api.Disconnect()
	}()

	for {
		msgType, reader, err := wsConn.Reader(context.TODO())
		if err != nil {
			api.logDebugf("Failed to receive response: %s", err.Error())
			return
		}

		if msgType != websocket.MessageText {
			api.logDebugf("Unexpected message type: %d", msgType)
			continue
		}

		buffer := bytes.NewBuffer(make([]byte, 0))
		_, err = buffer.ReadFrom(reader)

		if err != nil {
			api.logDebugf("Failed to read response: %s", err.Error())
			return
		}

		api.logDebugf("Received response: %s", buffer.String())
		respChan <- buffer.Bytes()
	}
}

// requestMapper forward requests to the Deriv API server and
// responses from the WebSocket server to the appropriate channels.
func (api *DerivAPI) requestMapper(respChan chan []byte, outputChan chan []byte, reqChan chan ApiReqest, closingChan chan int) {
	responseMap := make(map[int]chan []byte)

	defer func() {
		for reqID, channel := range responseMap {
			close(channel)
			delete(responseMap, reqID)
		}
	}()

	for {
		select {
		case rawResp := <-respChan:
			var response APIResponseReqID
			err := json.Unmarshal(rawResp, &response)
			if err != nil {
				continue
			}
			channel, ok := responseMap[response.ReqID]

			if ok {
				channel <- rawResp
			}
		case req, ok := <-reqChan:
			if !ok {
				return
			}
			responseMap[req.id] = req.respChan
			outputChan <- req.msg
		case reqID := <-closingChan:
			channel, okGet := responseMap[reqID]

			if okGet {
				close(channel)
				delete(responseMap, reqID)
			}

		}
	}
}

// Send sends a request to the Deriv API and returns a channel that will receive the response
func (api *DerivAPI) Send(reqID int, request any) (chan []byte, error) {
	err := api.Connect()

	if err != nil {
		return nil, err
	}

	msg, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	respChan := make(chan []byte)

	ApiReqest := ApiReqest{
		id:       reqID,
		msg:      msg,
		respChan: respChan,
	}

	api.reqChan <- ApiReqest

	return respChan, nil
}

// SendRequest sends a request to the Deriv API and returns the response
func (api *DerivAPI) SendRequest(reqID int, request any, response ApiResponse) (err error) {
	respChan, err := api.Send(reqID, request)

	if err != nil {
		return err
	}

	defer api.closeRequestChannel(reqID)

	select {
	case <-time.After(api.TimeOut):
		api.logDebugf("Timeout waiting for response for request %d", reqID)
		return fmt.Errorf("timeout")
	case responseJSON, ok := <-respChan:
		if !ok {
			api.logDebugf("Connection closed while waiting for response for request %d", reqID)
			return fmt.Errorf("connection closed")
		}

		if err = parseError(responseJSON); err != nil {
			return err
		}

		err = response.UnmarshalJSON([]byte(responseJSON))
		if err != nil {
			api.logDebugf("Failed to parse response for request %d: %s", reqID, err.Error())
			return err
		}

		return nil

	}
}

// getNextRequestID returns the next request ID
func (api *DerivAPI) getNextRequestID() int {
	return int(atomic.AddInt64(&api.lastRequestID, 1))
}

// closeRequestChannel closes the channel that receives the response for a request
func (api *DerivAPI) closeRequestChannel(reqID int) {
	api.closingChan <- reqID
}
