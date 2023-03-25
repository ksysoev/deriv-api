package deriv

import (
	"errors"
	"fmt"
	"golang.org/x/net/websocket"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestParseSubscription_ValidInput(t *testing.T) {
	input := `{"subscription": {"id": "123"}}`
	expected := SubscriptionResponse{Subscription: SubscriptionIDResponse{ID: "123"}}
	result, err := parseSubsciption(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, but got %+v", expected, result)
	}
}

func TestParseSubscription_InvalidJSONInput(t *testing.T) {
	input := `{"subscription": {"id": "123", "status": "active"`
	_, err := parseSubsciption(input)
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

func TestParseSubscription_InvalidSubscriptionData(t *testing.T) {
	input := `{"subscription": {"id": "123", "status": "active"}, "error": {"code": "invalid_subscription"}}`
	expectedErr := &APIError{Code: "invalid_subscription"}
	_, err := parseSubsciption(input)
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
	if !reflect.DeepEqual(err, expectedErr) {
		t.Errorf("Expected %+v, but got %+v", expectedErr, err)
	}
}

func TestParseSubscription_EmptyInput(t *testing.T) {
	_, err := parseSubsciption("")
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

func TestParseSubscription_EmptySubscriptionData(t *testing.T) {
	input := `{}`
	expectedErr := fmt.Errorf("subscription ID is empty")
	_, err := parseSubsciption(input)
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
	if errors.Is(err, expectedErr) {
		t.Errorf("Expected %+v, but got %+v", expectedErr, err)
	}
}

func TestNewNewSubcription(t *testing.T) {
	api, _ := NewDerivAPI("ws://example.com", 123, "en", "http://example.com")
	sub := NewSubcription[TicksResp](api)
	if sub == nil {
		t.Errorf("Expected a subscription, but got nil")
	}
}

func TestStart(t *testing.T) {
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
	fmt.Println(url)
	defer server.Close()
	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")
	err := api.Connect()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	sub := NewSubcription[TicksResp](api)

	if sub == nil {
		t.Errorf("Expected a subscription, but got nil")
	}

	reqID := 1
	var f TicksSubscribe = 1
	req := Ticks{Ticks: "R50", Subscribe: &f, ReqId: &reqID}
	err = sub.Start(reqID, req)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if sub.Stream == nil {
		t.Errorf("Expected a stream, but got nil")
	}

	// First message
	select {
	case tick := <-sub.Stream:
		if *tick.Tick.Quote != 186.9588 {
			t.Fatalf("Expected message to be %v, but got %v", 186.9588, *tick.Tick.Quote)
		}
	case <-time.After(time.Millisecond):
		t.Fatalf("Expected to get a response, but got nothing")
	}

	// Second message
	select {
	case tick := <-sub.Stream:
		if *tick.Tick.Quote != 186.9588 {
			t.Fatalf("Expected message to be %v, but got %v", 186.9588, tick.Tick.Quote)
		}
	case <-time.After(time.Millisecond):
		t.Fatalf("Expected to get a response, but got nothing")
	}
}

func TestForget(t *testing.T) {
	responses := make(chan string, 2)

	responses <- `{
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
		for resp := range responses {
			ws.Write([]byte(resp))
		}
	}))
	url := "ws://" + server.Listener.Addr().String()
	defer server.Close()

	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")
	err := api.Connect()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	sub := NewSubcription[TicksResp](api)

	if sub == nil {
		t.Errorf("Expected a subscription, but got nil")
	}

	reqID := api.getNextRequestID()
	fmt.Println(reqID)
	var f TicksSubscribe = 1
	req := Ticks{Ticks: "R50", Subscribe: &f, ReqId: &reqID}
	err = sub.Start(reqID, req)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if sub.Stream == nil {
		t.Errorf("Expected a stream, but got nil")
	}

	if sub.IsActive != true {
		t.Errorf("Expected subscription to be active, but got false")
	}

	go func() {
		time.Sleep(time.Millisecond * 5)
		responses <- `{
			"echo_req": {
			  "forget": "9ed45a5e-8f87-c735-2b63-36108719eadd"
			},
			"req_id": 2,
			"forget": 1,
			"msg_type": "forget"
		  }`
	}()
	sub.Forget()

	if sub.IsActive == true {
		t.Errorf("Expected subscription to be deactivated, but got true")
	}

	sub.Forget()

	if sub.IsActive == true {
		t.Errorf("Expected subscription to be deactivated, but got true")
	}
}

func TestStartAPIError(t *testing.T) {
	testResp := `{
		"echo_req": {
		  "subscribe": 1,
		  "ticks": "R_50"
		},
		"req_id": 1,
		"msg_type": "tick",
		"error": {
			"code": "WrongRequest",
			"message": "Invalid request"
		}
	  }`
	server := httptest.NewServer(websocket.Handler(func(ws *websocket.Conn) {
		ws.Write([]byte(testResp))
	}))
	url := "ws://" + server.Listener.Addr().String()
	fmt.Println(url)
	defer server.Close()
	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")
	err := api.Connect()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	sub := NewSubcription[TicksResp](api)

	if sub == nil {
		t.Errorf("Expected a subscription, but got nil")
	}

	reqID := 1
	var f TicksSubscribe = 1
	req := Ticks{Ticks: "R50", Subscribe: &f, ReqId: &reqID}
	err = sub.Start(reqID, req)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

func TestStartInvalidResponse(t *testing.T) {
	testResp := `{
		"req_id": 1,
		"msg_type": "tick",
		"subscription": {
		  "id": "9ed45a5e-8f87-c735-2b63-36108719eadd"
		},
		"tick": 1
	  }`
	server := httptest.NewServer(websocket.Handler(func(ws *websocket.Conn) {
		ws.Write([]byte(testResp))
	}))
	url := "ws://" + server.Listener.Addr().String()
	fmt.Println(url)
	defer server.Close()
	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")
	err := api.Connect()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	sub := NewSubcription[TicksResp](api)

	if sub == nil {
		t.Errorf("Expected a subscription, but got nil")
	}

	reqID := 1
	var f TicksSubscribe = 1
	req := Ticks{Ticks: "R50", Subscribe: &f, ReqId: &reqID}
	err = sub.Start(reqID, req)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}
