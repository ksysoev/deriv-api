package deriv

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/url"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/coder/websocket"
	"github.com/ksysoev/deriv-api/schema"
)

const (
	// DefaultKeepAliveInterval is the default interval for sending ping requests to keep the connection alive.
	defaultKeepAliveInterval = 25 * time.Second
	defaultTimeout           = 30 * time.Second
	defaultMaxMessageSize    = 1024 * 1024
)

// DerivAPI is the main struct for the DerivAPI client
type Client struct {
	ctx               context.Context
	Endpoint          *url.URL
	Origin            *url.URL
	ws                *websocket.Conn
	closingChan       chan int
	reqChan           chan APIReqest
	cancel            context.CancelFunc
	keepAliveInterval time.Duration
	lastRequestID     int64
	connectionLock    sync.Mutex
	debugEnabled      bool
}

// APIReqest is an interface for all API requests.
type APIReqest struct {
	respChan chan []byte
	msg      []byte
	id       int
}

type APIResponseReqID struct {
	ReqID int `json:"req_id"`
}

type APIOption func(api *Client)

// NewDerivAPI creates a new instance of DerivAPI by parsing and validating the given
// endpoint URL, appID, language, and origin URL. It returns a pointer to a DerivAPI object
// or an error if any of the validation checks fail.
//
// Parameters:
// - endpoint: string - The WebSocket endpoint URL for the DerivAPI server
// - appID: int - The app ID for the DerivAPI server
// - lang: string - The language code (ISO 639-1) for the DerivAPI server
// - origin: string - The origin URL for the DerivAPI server
// - opts: APIOption - A variadic list of APIOption functions to configure the DerivAPI object (optional)
//   - KeepAlive: A APIOption function to keep the connection alive by sending ping requests.
//   - Debug: A APIOption function to enable debug messages.
//
// Returns:
//   - *Client: A pointer to a new instance of DerivAPI with the validated endpoint, appID,
//     language, and origin values.
//   - error: An error if any of the validation checks fail.
//
// Example:
//
//	api, err := NewDerivAPI("wss://trade.deriv.com/websockets/v3", 12345, "en", "https://myapp.com")
//	if err != nil {
//		log.Fatal(err)
//	}
func NewDerivAPI(endpoint string, appID int, lang, origin string, opts ...APIOption) (*Client, error) {
	urlEnpoint, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	urlOrigin, err := url.Parse(origin)
	if err != nil {
		return nil, err
	}

	if urlEnpoint.Scheme != "wss" && urlEnpoint.Scheme != "ws" {
		return nil, ErrInvalidSchema
	}

	if appID < 1 {
		return nil, ErrInvalidAppID
	}

	if lang == "" || len(lang) != 2 {
		return nil, ErrInvalidLanguage
	}

	query := urlEnpoint.Query()
	query.Add("app_id", strconv.Itoa(appID))
	query.Add("l", lang)

	urlEnpoint.RawQuery = query.Encode()

	api := Client{
		Origin:            urlOrigin,
		Endpoint:          urlEnpoint,
		lastRequestID:     0,
		connectionLock:    sync.Mutex{},
		closingChan:       make(chan int),
		keepAliveInterval: 0,
		ctx:               context.Background(),
	}

	for _, opt := range opts {
		opt(&api)
	}

	api.ctx, api.cancel = context.WithCancel(api.ctx)

	return &api, nil
}

// KeepAlive option which keeps the connection alive by sending ping requests.
// By default the websocket connection is closed after 30 seconds of inactivity.
func KeepAlive(api *Client) {
	api.keepAliveInterval = defaultKeepAliveInterval
}

// Debug option which enables debug messages.
func Debug(api *Client) {
	api.debugEnabled = true
}

// logDebugf prints a debug message if debug is enabled.
func (api *Client) logDebugf(format string, v ...interface{}) {
	if api.debugEnabled {
		log.Printf(format, v...)
	}
}

// Connect establishes a WebSocket connection with the Deriv API endpoint.
// Returns an error if it fails to connect to WebSoket server.
func (api *Client) Connect() error {
	api.connectionLock.Lock()
	defer api.connectionLock.Unlock()

	if api.ws != nil {
		return nil
	}

	api.logDebugf("Connecting to %s", api.Endpoint.String())

	ws, resp, err := websocket.Dial(context.TODO(), api.Endpoint.String(), &websocket.DialOptions{
		HTTPHeader: map[string][]string{
			"Origin": {api.Origin.String()},
		},
	})
	if err != nil {
		api.logDebugf("Failed to establish WS connection: %s", err.Error())
		return err
	}

	ws.SetReadLimit(defaultMaxMessageSize)

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	api.logDebugf("Connected to %s", api.Endpoint.String())

	api.ws = ws

	api.reqChan = make(chan APIReqest)
	respChan := make(chan []byte)
	outputChan := make(chan []byte)

	go api.handleResponses(ws, respChan)
	go api.requestSender(ws, outputChan)
	go api.requestMapper(respChan, outputChan, api.reqChan, api.closingChan)

	if api.keepAliveInterval > 0 {
		go func() {
			for {
				select {
				case <-time.After(api.keepAliveInterval):
					ctx, cancel := context.WithTimeout(api.ctx, defaultTimeout)
					_, err := api.Ping(ctx, schema.Ping{Ping: 1})

					cancel()

					if err != nil {
						return
					}
				case <-api.ctx.Done():
					return
				}
			}
		}()
	}

	return nil
}

// Disconnect gracefully disconnects from the Deriv API WebSocket server
// and cleans up any resources associated with the connection.
// This function should only be called after
// the WebSocket connection has been established using the Connect method.
func (api *Client) Disconnect() {
	api.connectionLock.Lock()
	defer api.connectionLock.Unlock()

	if api.ws == nil {
		return
	}

	api.logDebugf("Disconnecting from %s", api.Endpoint.String())

	api.cancel()

	api.ws.Close(websocket.StatusNormalClosure, "disconnecting")
	api.ws = nil
}

// requestSender sends requests to the Deriv API WebSocket server over the provided WebSocket connection.
// It reads requests from the reqChan channel and sends them using the websocket.Message.Send method.
// If an error occurs while sending a request, it calls the Disconnect method to gracefully disconnect from the WebSocket server.
func (api *Client) requestSender(wsConn *websocket.Conn, reqChan chan []byte) {
	defer func() {
		api.Disconnect()
	}()

	for {
		select {
		case <-api.ctx.Done():
			return
		case req, ok := <-reqChan:
			if !ok {
				return
			}

			api.logDebugf("Sending request: %s", req)

			err := wsConn.Write(api.ctx, websocket.MessageText, req)
			if err != nil {
				api.logDebugf("Failed to send request: %s", err.Error())
				return
			}
		}
	}
}

// handleResponses reads responses from the Deriv API WebSocket server over the provided WebSocket connection.
// It reads responses using the websocket.Message.Receive method and sends them to the respChan channel.
// If an error occurs while receiving a response, it calls the Disconnect method to gracefully disconnect from the WebSocket server.
// The function returns when the WebSocket connection is closed or when an error occurs.
func (api *Client) handleResponses(wsConn *websocket.Conn, respChan chan []byte) {
	defer func() {
		close(respChan)

		api.Disconnect()
	}()

	for api.ctx.Err() == nil {
		msgType, reader, err := wsConn.Reader(api.ctx)
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

		select {
		case <-api.ctx.Done():
			return
		case respChan <- buffer.Bytes():
		}
	}
}

// requestMapper forward requests to the Deriv API server and
// responses from the WebSocket server to the appropriate channels.
func (api *Client) requestMapper(respChan, outputChan chan []byte, reqChan chan APIReqest, closingChan chan int) {
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

			if err := json.Unmarshal(rawResp, &response); err != nil {
				continue
			}

			if channel, ok := responseMap[response.ReqID]; ok {
				channel <- rawResp
			}
		case req, ok := <-reqChan:
			if !ok {
				return
			}

			responseMap[req.id] = req.respChan
			outputChan <- req.msg
		case reqID := <-closingChan:
			if channel, ok := responseMap[reqID]; ok {
				close(channel)
				delete(responseMap, reqID)
			}
		case <-api.ctx.Done():
			return
		}
	}
}

// Send sends a request to the Deriv API and returns a channel that will receive the response
func (api *Client) Send(ctx context.Context, reqID int, request any) (chan []byte, error) {
	err := api.Connect()

	if err != nil {
		return nil, err
	}

	msg, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	respChan := make(chan []byte)

	req := APIReqest{
		id:       reqID,
		msg:      msg,
		respChan: respChan,
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-api.ctx.Done():
		return nil, ErrConnectionClosed
	case api.reqChan <- req:
		return respChan, nil
	}
}

// SendRequest sends a request to the Deriv API and returns the response
func (api *Client) SendRequest(ctx context.Context, reqID int, request, response any) error {
	respChan, err := api.Send(ctx, reqID, request)

	if err != nil {
		return err
	}

	defer api.closeRequestChannel(reqID)

	select {
	case <-api.ctx.Done():
		return ErrConnectionClosed
	case <-ctx.Done():
		return ctx.Err()
	case responseJSON, ok := <-respChan:
		if !ok {
			api.logDebugf("Connection closed while waiting for response for request %d", reqID)
			return ErrConnectionClosed
		}

		if err := parseError(responseJSON); err != nil {
			return err
		}

		if err := json.Unmarshal(responseJSON, response); err != nil {
			api.logDebugf("Failed to parse response for request %d: %s", reqID, err.Error())
			return err
		}

		return nil
	}
}

// getNextRequestID returns the next request ID
func (api *Client) getNextRequestID() int {
	return int(atomic.AddInt64(&api.lastRequestID, 1))
}

// closeRequestChannel closes the channel that receives the response for a request
func (api *Client) closeRequestChannel(reqID int) {
	select {
	case api.closingChan <- reqID:
	case <-api.ctx.Done():
	}
}
