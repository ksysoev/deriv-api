package deriv

import (
	"golang.org/x/net/websocket"
	"io"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewDerivAPI(t *testing.T) {
	// Valid endpoint and origin
	endpoint := "wss://example.com/ws"
	origin := "https://example.com"
	appID := 123
	lang := "en"
	api, err := NewDerivAPI(endpoint, appID, lang, origin)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if api.Endpoint.String() != endpoint+"?app_id=123&l=en" {
		t.Errorf("Unexpected endpoint: got %v, want %v", api.Endpoint.String(), endpoint)
	}
	if api.Origin.String() != origin {
		t.Errorf("Unexpected origin: got %v, want %v", api.Origin.String(), origin)
	}
	if api.AppID != appID {
		t.Errorf("Unexpected app ID: got %v, want %v", api.AppID, appID)
	}
	if api.Lang != lang {
		t.Errorf("Unexpected language: got %v, want %v", api.Lang, lang)
	}

	// Invalid endpoint scheme
	endpoint = "https://example.com/ws"
	origin = "https://example.com"
	appID = 123
	lang = "en"
	api, err = NewDerivAPI(endpoint, appID, lang, origin)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Invalid app ID
	endpoint = "wss://example.com/ws"
	origin = "https://example.com"
	appID = -1
	lang = "en"
	api, err = NewDerivAPI(endpoint, appID, lang, origin)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Invalid language
	endpoint = "wss://example.com/ws"
	origin = "https://example.com"
	appID = 123
	lang = "eng"
	api, err = NewDerivAPI(endpoint, appID, lang, origin)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Invalid endpoint URL
	endpoint = ":invalid:"
	origin = "https://example.com"
	appID = 123
	lang = "en"
	api, err = NewDerivAPI(endpoint, appID, lang, origin)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Invalid origin URL
	endpoint = "wss://example.com/ws"
	origin = ":invalid:"
	appID = 123
	lang = "en"
	api, err = NewDerivAPI(endpoint, appID, lang, origin)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestGetNextRequestID(t *testing.T) {
	api := &DerivAPI{lastRequestID: 0}
	requestIDs := make(map[int]bool)
	orderedRequestIDs := make([]int, 0)
	numRequests := 5

	for i := 0; i < numRequests; i++ {
		requestID := api.getNextRequestID()
		if _, ok := requestIDs[requestID]; ok {
			t.Errorf("Request ID %d already used", requestID)
		}
		requestIDs[requestID] = true
		orderedRequestIDs = append(orderedRequestIDs, requestID)
	}

	if len(requestIDs) != numRequests {
		t.Errorf("Expected %d unique request IDs, but got %d", numRequests, len(requestIDs))
	}

	lastID := 0
	for _, id := range orderedRequestIDs {
		if id <= lastID {
			t.Errorf("Request IDs not increasing, lastID=%d currentID=%d", lastID, id)
		}
		lastID = id
	}
}

func TestConnect(t *testing.T) {
	server := httptest.NewServer(websocket.Handler(func(ws *websocket.Conn) { io.Copy(ws, ws) }))
	url := "ws://" + server.Listener.Addr().String()

	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")

	err := api.Connect()
	if err != nil {
		t.Fatalf("Failed to connect to mocked WebSocket server: %v", err)
	}

	if api.ws == nil {
		t.Fatalf("WebSocket connection not established")
	}

	ws1 := api.ws

	err = api.Connect()
	if err != nil {
		t.Fatalf("Failed to connect to mocked WebSocket server: %v", err)
	}

	if api.ws != ws1 {
		t.Fatalf("Expected WebSocket connection to be the same, but got a different connection")
	}

	// Close the WebSocket connection and test that we can't connect again
	server.Close()

	api, _ = NewDerivAPI(url, 123, "en", "http://example.com")

	err = api.Connect()
	if err == nil {
		t.Fatalf("Expected fail to connect but didn't get any error")
	}

	if api.ws != nil {
		t.Fatalf("Expected fail to connect, but WebSocket connection was established")
	}
}

func TestDisconnect(t *testing.T) {
	server := httptest.NewServer(websocket.Handler(func(ws *websocket.Conn) { io.Copy(ws, ws) }))
	url := "ws://" + server.Listener.Addr().String()

	defer server.Close()

	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")

	api.Disconnect()

	if api.ws != nil {
		t.Fatalf("Expected WebSocket connection to be nil, but got a connection")
	}

	err := api.Connect()
	if err != nil {
		t.Fatalf("Failed to connect to mocked WebSocket server: %v", err)
	}

	if api.ws == nil {
		t.Fatalf("WebSocket connection not established")
	}

	api.Disconnect()

	if api.ws != nil {
		t.Fatalf("Expected WebSocket connection to be nil, but got a connection")
	}
}

func TestSend(t *testing.T) {
	server := httptest.NewServer(websocket.Handler(func(ws *websocket.Conn) { io.Copy(ws, ws) }))
	url := "ws://" + server.Listener.Addr().String()

	defer server.Close()

	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")

	err := api.Connect()
	if err != nil {
		t.Fatalf("Failed to connect to mocked WebSocket server: %v", err)
	}

	if api.ws == nil {
		t.Fatalf("WebSocket connection not established")
	}

	respChan, err := api.Send(1, struct {
		ReqID int `json:"req_id"`
	}{1})
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	if respChan == nil {
		t.Fatalf("Response channel is nil")
	}

	msg := <-respChan
	testMsg := "{\"req_id\":1}"
	if msg != testMsg {
		t.Fatalf("Expected message to be %s, but got %s", testMsg, msg)
	}
}

func TestSendToDisconnectedConnection(t *testing.T) {
	server := httptest.NewServer(websocket.Handler(func(ws *websocket.Conn) { io.Copy(ws, ws) }))
	url := "ws://" + server.Listener.Addr().String()
	server.Close()

	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")

	respChan, err := api.Send(1, struct {
		ReqID int `json:"req_id"`
	}{1})

	if err == nil {
		t.Fatalf("Expected error, got nil")
	}

	if respChan != nil {
		t.Fatalf("Expected response channel to be nil, but got a channel")
	}
}

func TestSendReqWhichNobodyWaits(t *testing.T) {
	server := httptest.NewServer(websocket.Handler(func(ws *websocket.Conn) { io.Copy(ws, ws) }))
	url := "ws://" + server.Listener.Addr().String()

	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")

	respChan, err := api.Send(2, struct {
		ReqID int `json:"req_id"`
	}{1})

	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}
	select {
	case <-respChan:
		t.Fatalf("Expected no response, but got a response")
	case <-time.After(time.Millisecond):
	}
}

func TestSendRequest(t *testing.T) {
	testResp := `{
		"echo_req": {
		  "ping": 1,
		  "req_id": 1
		},
		"msg_type": "ping",
		"ping": "pong",
		"req_id": 1
	  }`
	server := httptest.NewServer(websocket.Handler(func(ws *websocket.Conn) { ws.Write([]byte(testResp)) }))
	url := "ws://" + server.Listener.Addr().String()
	defer server.Close()

	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")

	_ = api.Connect()

	reqID := 1
	req := Ping{Ping: 1, ReqId: &reqID}
	var resp PingResp
	err := api.SendRequest(reqID, req, &resp)

	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	if *resp.Ping != "pong" {
		t.Fatalf("Expected response to be pong, but got %s", *resp.Ping)
	}
}

func TestSubscribeRequest(t *testing.T) {
	testResp := `{
		"echo_req": {
		  "subscribe": 1,
		  "ticks": "R_50"
		},
		"req_id": 1,
		"msg_type": "tick",
		"subscription": {
		  "id": "9ed45a5e-8f87-c735-2b63-36108719eadd"
		},
		"tick": {
		  "ask": 186.9688,
		  "bid": 186.9488,
		  "epoch": 1679722832,
		  "id": "9ed45a5e-8f87-c735-2b63-36108719eadd",
		  "pip_size": 4,
		  "quote": 186.9588,
		  "symbol": "R_50"
		}
	  }`
	server := httptest.NewServer(websocket.Handler(func(ws *websocket.Conn) {
		ws.Write([]byte(testResp))
		ws.Write([]byte(testResp))
	}))
	url := "ws://" + server.Listener.Addr().String()
	defer server.Close()

	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")

	_ = api.Connect()

	reqID := 1
	var f TicksSubscribe = 1
	req := Ticks{Ticks: "R50", Subscribe: &f, ReqId: &reqID}
	respChan, err := api.SubscribeRequest(reqID, req)

	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	if respChan == nil {
		t.Fatalf("Expected to get response channel, but got nil")
	}

	// First message
	select {
	case msg := <-respChan:
		if msg != testResp {
			t.Fatalf("Expected message to be %s, but got %s", testResp, msg)
		}
	case <-time.After(time.Millisecond):
		t.Fatalf("Expected to get a response, but got nothing")
	}

	// Second message
	select {
	case msg := <-respChan:
		if msg != testResp {
			t.Fatalf("Expected message to be %s, but got %s", testResp, msg)
		}
	case <-time.After(time.Millisecond):
		t.Fatalf("Expected to get a response, but got nothing")
	}

}
