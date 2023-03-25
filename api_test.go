package deriv

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"net/http/httptest"
	"testing"
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

	fmt.Println("Server URL:", url)

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
