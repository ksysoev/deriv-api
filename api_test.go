package deriv

import (
	"bufio"
	"context"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/coder/websocket"
	"github.com/ksysoev/deriv-api/schema"
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

	if _, err = NewDerivAPI(endpoint, appID, lang, origin); err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Invalid app ID
	endpoint = "wss://example.com/ws"
	origin = "https://example.com"
	appID = -1
	lang = "en"

	if _, err = NewDerivAPI(endpoint, appID, lang, origin); err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Invalid language
	endpoint = "wss://example.com/ws"
	origin = "https://example.com"
	appID = 123
	lang = "eng"

	if _, err = NewDerivAPI(endpoint, appID, lang, origin); err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Invalid endpoint URL
	endpoint = ":invalid:"
	origin = "https://example.com"
	appID = 123
	lang = "en"

	if _, err = NewDerivAPI(endpoint, appID, lang, origin); err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Invalid origin URL
	endpoint = "wss://example.com/ws"
	origin = ":invalid:"
	appID = 123
	lang = "en"

	if _, err = NewDerivAPI(endpoint, appID, lang, origin); err == nil {
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
	server := newMockWSServer(echoHandler)
	url := "ws://" + server.Listener.Addr().String()

	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")

	err := api.Connect()
	if err != nil {
		t.Errorf("Failed to connect to mocked WebSocket server: %v", err)
	}

	if api.ws == nil {
		t.Errorf("WebSocket connection not established")
	}

	ws1 := api.ws

	err = api.Connect()
	if err != nil {
		t.Errorf("Failed to connect to mocked WebSocket server: %v", err)
	}

	if api.ws != ws1 {
		t.Errorf("Expected WebSocket connection to be the same, but got a different connection")
	}

	// Close the WebSocket connection and test that we can't connect again
	server.Close()

	api, _ = NewDerivAPI(url, 123, "en", "http://example.com")

	err = api.Connect()
	if err == nil {
		t.Errorf("Expected fail to connect but didn't get any error")
	}

	if api.ws != nil {
		t.Errorf("Expected fail to connect, but WebSocket connection was established")
	}
}

func TestDisconnect(t *testing.T) {
	server := newMockWSServer(echoHandler)
	url := "ws://" + server.Listener.Addr().String()

	defer server.Close()

	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")

	api.Disconnect()

	if api.ws != nil {
		t.Errorf("Expected WebSocket connection to be nil, but got a connection")
	}

	err := api.Connect()
	if err != nil {
		t.Errorf("Failed to connect to mocked WebSocket server: %v", err)
	}

	if api.ws == nil {
		t.Errorf("WebSocket connection not established")
	}

	api.Disconnect()

	if api.ws != nil {
		t.Errorf("Expected WebSocket connection to be nil, but got a connection")
	}
}

func TestSend(t *testing.T) {
	server := newMockWSServer(echoHandler)
	url := "ws://" + server.Listener.Addr().String()

	defer server.Close()

	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")

	err := api.Connect()
	if err != nil {
		t.Errorf("Failed to connect to mocked WebSocket server: %v", err)
	}

	if api.ws == nil {
		t.Errorf("WebSocket connection not established")
	}

	respChan, err := api.Send(1, struct {
		ReqID int `json:"req_id"`
	}{1})
	if err != nil {
		t.Errorf("Failed to send message: %v", err)
	}

	if respChan == nil {
		t.Errorf("Response channel is nil")
	}

	msg := <-respChan

	testMsg := "{\"req_id\":1}"
	if string(msg) != testMsg {
		t.Errorf("Expected message to be %s, but got %s", testMsg, msg)
	}
}

func TestSendToDisconnectedConnection(t *testing.T) {
	server := newMockWSServer(echoHandler)
	url := "ws://" + server.Listener.Addr().String()
	server.Close()

	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")

	respChan, err := api.Send(1, struct {
		ReqID int `json:"req_id"`
	}{1})

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if respChan != nil {
		t.Errorf("Expected response channel to be nil, but got a channel")
	}
}

func TestSendReqWhichNobodyWaits(t *testing.T) {
	server := newMockWSServer(echoHandler)
	url := "ws://" + server.Listener.Addr().String()

	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")

	respChan, err := api.Send(2, struct {
		ReqID int `json:"req_id"`
	}{1})

	if err != nil {
		t.Errorf("Failed to send message: %v", err)
	}
	select {
	case <-respChan:
		t.Errorf("Expected no response, but got a response")
	case <-time.After(time.Millisecond):
	}
}

func TestSendRequestTimeout(t *testing.T) {
	server := newMockWSServer(
		onMessageHanler(func(_ *websocket.Conn, _ websocket.MessageType, _ []byte) {
			time.Sleep(time.Second)
		}),
	)
	defer server.Close()
	url := "ws://" + server.Listener.Addr().String()

	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")
	api.TimeOut = time.Millisecond

	reqID := 1
	req := schema.Ping{Ping: 1, ReqId: &reqID}

	var resp schema.PingResp

	err := api.SendRequest(reqID, req, &resp)

	if err != nil && err.Error() != "timeout" {
		t.Errorf("Expected timeout error, got %v", err)
	}
}

func TestSendRequestAndGotInvalidJSON(t *testing.T) {
	testResp := `{Corrupted JSON}`
	server := newMockWSServer(
		onMessageHanler(func(ws *websocket.Conn, _ websocket.MessageType, _ []byte) {
			err := ws.Write(context.Background(), websocket.MessageText, []byte(testResp))
			if err != nil {
				return
			}
			time.Sleep(time.Second) // to keep the connection open
		}))

	defer server.Close()

	url := "ws://" + server.Listener.Addr().String()
	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")

	_ = api.Connect()

	respChan, err := api.Send(2, struct {
		ReqID int `json:"req_id"`
	}{1})

	if err != nil {
		t.Errorf("Failed to send message: %v", err)
	}
	select {
	case _, ok := <-respChan:
		if ok {
			t.Errorf("Expected no response, but got a response")
		}
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
	server := newMockWSServer(
		onMessageHanler(func(ws *websocket.Conn, _ websocket.MessageType, _ []byte) {
			err := ws.Write(context.Background(), websocket.MessageText, []byte(testResp))
			if err != nil {
				return
			}
			time.Sleep(time.Second) // to keep the connection open
		}))

	defer server.Close()

	url := "ws://" + server.Listener.Addr().String()

	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")

	err := api.Connect()

	if err != nil {
		t.Errorf("Failed to connect to mocked WebSocket server: %v", err)
	}

	reqID := 1
	req := schema.Ping{Ping: 1, ReqId: &reqID}

	var resp schema.PingResp

	err = api.SendRequest(reqID, req, &resp)

	if err != nil {
		t.Errorf("Failed to send message: %v", err)
	}

	if *resp.Ping != "pong" {
		t.Errorf("Expected response to be pong, but got %s", *resp.Ping)
	}
}

func TestSendRequestFailed(t *testing.T) {
	server := newMockWSServer(
		onMessageHanler(func(ws *websocket.Conn, _ websocket.MessageType, _ []byte) {
			err := ws.Write(context.Background(), websocket.MessageText, []byte(""))
			if err != nil {
				return
			}
			time.Sleep(time.Second) // to keep the connection open
		}))
	url := "ws://" + server.Listener.Addr().String()
	server.Close()

	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")

	reqID := 1
	req := schema.Ping{Ping: 1, ReqId: &reqID}

	var resp schema.PingResp

	err := api.SendRequest(reqID, req, &resp)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestKeepConnectionAlive(t *testing.T) {
	resChan := make(chan string)
	testResp := `{
			"echo_req": {
			  "ping": 1,
			  "req_id": 1
			},
			"msg_type": "ping",
			"ping": "pong",
			"req_id": 1
		  }`

	server := newMockWSServer(
		onMessageHanler(func(ws *websocket.Conn, _ websocket.MessageType, msg []byte) {
			resChan <- string(msg)

			err := ws.Write(context.Background(), websocket.MessageText, []byte(testResp))
			if err != nil {
				return
			}
			time.Sleep(time.Second) // to keep the connection open
		}))

	url := "ws://" + server.Listener.Addr().String()
	defer server.Close()

	api, _ := NewDerivAPI(url, 123, "en", "http://example.com", KeepAlive)
	api.keepAliveInterval = time.Millisecond

	if err := api.Connect(); err != nil {
		t.Errorf("Failed to connect to mocked WebSocket server: %v", err)
	}

	select {
	case msg := <-resChan:
		if msg == "" {
			t.Errorf("Expected to receive ping request, but got nothing")
		}
	case <-time.After(time.Millisecond * 2):
		t.Errorf("Expected to receive ping request, but got nothing")
	}

	api.Disconnect()
	api.keepAlive = false

	if err := api.Connect(); err != nil {
		t.Errorf("Failed to connect to mocked WebSocket server: %v", err)
	}

	select {
	case msg, ok := <-resChan:
		if ok && msg != "" {
			t.Errorf("Expected to not receive ping request, but got n%sn>", msg)
		}
	case <-time.After(time.Millisecond * 2):
	}
}

func TestDebugLogs(t *testing.T) {
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Errorf("Failed to create pipe: %v", err)
	}

	defer reader.Close()

	log.SetOutput(writer)

	scanner := bufio.NewScanner(reader)
	server := newMockWSServer(echoHandler)
	url := "ws://" + server.Listener.Addr().String()
	api, _ := NewDerivAPI(url, 123, "en", "http://example.com", Debug)

	_ = api.Connect()

	scanner.Scan()
	got := scanner.Text()
	want := "Connecting to " + url

	if !strings.Contains(got, want) {
		t.Errorf("Expected log to contain %s, but got %s", want, got)
	}
}
