package deriv

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/coder/websocket"
	"github.com/ksysoev/deriv-api/schema"
)

func TestParseSubscription_ValidInput(t *testing.T) {
	input := []byte(`{"subscription": {"id": "123"}}`)
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
	input := []byte(`{"subscription": {"id": "123", "status": "active"`)

	if _, err := parseSubsciption(input); err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

func TestParseSubscription_InvalidSubscriptionData(t *testing.T) {
	input := []byte(`{"subscription": {"id": "123", "status": "active"}, "error": {"code": "invalid_subscription"}}`)
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
	_, err := parseSubsciption([]byte(""))
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

func TestParseSubscription_EmptySubscriptionData(t *testing.T) {
	input := []byte(`{}`)
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
	sub := NewSubcription[schema.TicksResp, schema.TicksResp](api)

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
	server := newMockWSServer(
		onMessageHanler(func(ws *websocket.Conn, _ websocket.MessageType, _ []byte) {
			if err := ws.Write(context.Background(), websocket.MessageText, []byte(testResp)); err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if err := ws.Write(context.Background(), websocket.MessageText, []byte(testResp)); err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			time.Sleep(time.Second) // to keep the connection open
		}))

	defer server.Close()

	url := "ws://" + server.Listener.Addr().String()
	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")

	if err := api.Connect(); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	sub := NewSubcription[schema.TicksResp, schema.TicksResp](api)

	if sub == nil {
		t.Errorf("Expected a subscription, but got nil")
	}

	reqID := 1

	var f schema.TicksSubscribe = 1

	req := schema.Ticks{Ticks: "R50", Subscribe: &f, ReqId: &reqID}
	initResp, err := sub.Start(reqID, req)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	// First message
	if *initResp.Tick.Quote != 186.9588 {
		t.Errorf("Expected message to be %v, but got %v", 186.9588, *initResp.Tick.Quote)
	}

	// Second message
	select {
	case tick, ok := <-sub.GetStream():
		if !ok {
			t.Errorf("Expected to get a response, but channel is closed")
		} else if *tick.Tick.Quote != 186.9588 {
			t.Errorf("Expected message to be %v, but got %v", 186.9588, tick.Tick.Quote)
		}
	case <-time.After(time.Millisecond):
		t.Errorf("Expected to get a response, but got nothing")
	}

	if _, err := sub.Start(reqID, req); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if sub.IsActive() == false {
		t.Errorf("Expected subscription to be active, but got inactive")
	}
}

func TestStartFailed(t *testing.T) {
	server := newMockWSServer(
		onMessageHanler(func(ws *websocket.Conn, _ websocket.MessageType, _ []byte) {
			if err := ws.Write(context.Background(), websocket.MessageText, []byte("")); err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			time.Sleep(time.Second) // to keep the connection open
		}))

	server.Close()

	url := "ws://" + server.Listener.Addr().String()
	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")

	sub := NewSubcription[schema.TicksResp, schema.TicksResp](api)

	reqID := 1

	var f schema.TicksSubscribe = 1

	req := schema.Ticks{Ticks: "R50", Subscribe: &f, ReqId: &reqID}
	_, err := sub.Start(reqID, req)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
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

	server := newMockWSServer(
		onMessageHanler(func(ws *websocket.Conn, _ websocket.MessageType, _ []byte) {
			for resp := range responses {
				if err := ws.Write(context.Background(), websocket.MessageText, []byte(resp)); err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			time.Sleep(time.Second) // to keep the connection open
		}))

	defer server.Close()

	url := "ws://" + server.Listener.Addr().String()

	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")

	if err := api.Connect(); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	sub := NewSubcription[schema.TicksResp, schema.TicksResp](api)

	if sub == nil {
		t.Fatal("Expected a subscription, but got nil")
	}

	reqID := api.getNextRequestID()

	var f schema.TicksSubscribe = 1

	req := schema.Ticks{Ticks: "R50", Subscribe: &f, ReqId: &reqID}

	if _, err := sub.Start(reqID, req); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if sub.Stream == nil {
		t.Fatal("Expected a stream, but got nil")
	}

	if sub.IsActive() != true {
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

	if err := sub.Forget(); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if sub.IsActive() == true {
		t.Errorf("Expected subscription to be deactivated, but got true")
	}

	if err := sub.Forget(); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if sub.IsActive() == true {
		t.Errorf("Expected subscription to be deactivated, but got true")
	}
}

func TestForgetFailed(t *testing.T) {
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

	server := newMockWSServer(
		onMessageHanler(func(ws *websocket.Conn, _ websocket.MessageType, _ []byte) {
			for resp := range responses {
				if err := ws.Write(context.Background(), websocket.MessageText, []byte(resp)); err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			time.Sleep(time.Second) // to keep the connection open
		}))

	defer server.Close()

	url := "ws://" + server.Listener.Addr().String()
	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")

	if err := api.Connect(); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	sub := NewSubcription[schema.TicksResp, schema.TicksResp](api)

	if sub == nil {
		t.Fatal("Expected a subscription, but got nil")
	}

	reqID := api.getNextRequestID()

	var f schema.TicksSubscribe = 1

	req := schema.Ticks{Ticks: "R50", Subscribe: &f, ReqId: &reqID}
	_, err := sub.Start(reqID, req)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if sub.Stream == nil {
		t.Fatal("Expected a stream, but got nil")
	}

	if sub.IsActive() != true {
		t.Errorf("Expected subscription to be active, but got false")
	}

	go func() {
		time.Sleep(time.Millisecond * 5)
		responses <- `{
			"echo_req": {
			  "forget": "9ed45a5e-8f87-c735-2b63-36108719eadd"
			},
			"req_id": 2,
			"msg_type": "forget",
			"error": {
				"code": "WrongRequest",
				"message": "Invalid request"
			}
		  }`
	}()

	if err = sub.Forget(); err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if sub.IsActive() == false {
		t.Errorf("Expected subscription to be active, but got false")
	}
}

func TestStartInvalidResponses(t *testing.T) {
	tests := []struct{ name, resp string }{
		{
			name: "InvalidResponse",
			resp: `{
				"req_id": 1,
				"msg_type": "tick",
				"subscription": {
				"id": "9ed45a5e-8f87-c735-2b63-36108719eadd"
				},
				"tick": 1
			}`,
		},
		{
			name: "APIError",
			resp: `{
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
			}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := newMockWSServer(
				onMessageHanler(func(ws *websocket.Conn, _ websocket.MessageType, _ []byte) {
					if err := ws.Write(context.Background(), websocket.MessageText, []byte(test.resp)); err != nil {
						t.Errorf("Unexpected error: %v", err)
					}

					time.Sleep(time.Second) // to keep the connection open
				}))

			defer server.Close()

			url := "ws://" + server.Listener.Addr().String()
			api, _ := NewDerivAPI(url, 123, "en", "http://example.com")

			if err := api.Connect(); err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			sub := NewSubcription[schema.TicksResp, schema.TicksResp](api)

			if sub == nil {
				t.Errorf("Expected a subscription, but got nil")
			}

			reqID := 1

			var f schema.TicksSubscribe = 1

			req := schema.Ticks{Ticks: "R50", Subscribe: &f, ReqId: &reqID}

			if _, err := sub.Start(reqID, req); err == nil {
				t.Errorf("Expected an error, but got nil")
			}
		})
	}
}

func TestStartAPIErrorInSubscription(t *testing.T) {
	responses := []string{`{
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
	  }`,
		`{ 
		"echo_req": {
			"subscribe": 1,
			"ticks": "R_50"
			},
			"req_id": 1,
			"msg_type": "tick",
			"error": {
			"code": "InvalidSymbol",
			"message": "Invalid symbol"
			}
		}`,
		`{ "req_id": 1 }`}

	server := newMockWSServer(
		onMessageHanler(func(ws *websocket.Conn, _ websocket.MessageType, _ []byte) {
			for _, resp := range responses {
				if err := ws.Write(context.Background(), websocket.MessageText, []byte(resp)); err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			time.Sleep(time.Second) // to keep the connection open
		}))

	defer server.Close()

	url := "ws://" + server.Listener.Addr().String()
	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")

	if err := api.Connect(); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	sub := NewSubcription[schema.TicksResp, schema.TicksResp](api)

	if sub == nil {
		t.Errorf("Expected a subscription, but got nil")
	}

	reqID := 1

	var f schema.TicksSubscribe = 1

	req := schema.Ticks{Ticks: "R50", Subscribe: &f, ReqId: &reqID}
	initResp, err := sub.Start(reqID, req)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// First message
	if *initResp.Tick.Quote != 186.9588 {
		t.Errorf("Expected message to be %v, but got %v", 186.9588, *initResp.Tick.Quote)
	}

	// Second message
	select {
	case tick, ok := <-sub.GetStream():
		if ok {
			t.Errorf("Expected to get noting, but got response: %v", tick)
		}
	case <-time.After(time.Millisecond):
	}
}

func TestStartTimeout(t *testing.T) {
	server := newMockWSServer(
		onMessageHanler(func(_ *websocket.Conn, _ websocket.MessageType, _ []byte) {
			time.Sleep(time.Second) // to keep the connection open
		}))
	defer server.Close()
	url := "ws://" + server.Listener.Addr().String()

	api, _ := NewDerivAPI(url, 123, "en", "http://example.com")
	api.TimeOut = time.Millisecond

	sub := NewSubcription[schema.TicksResp, schema.TicksResp](api)

	reqID := 1

	var f schema.TicksSubscribe = 1

	req := schema.Ticks{Ticks: "R50", Subscribe: &f, ReqId: &reqID}
	_, err := sub.Start(reqID, req)

	if err != nil && err.Error() != "timeout" {
		t.Errorf("Expected timeout error, got %v", err)
	}
}
