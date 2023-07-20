package deriv

// This package provides functionality for working with subscriptions using the Deriv API.

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
)

// Subscription represents a subscription instance.
type Subsciption[initResp any, Resp any] struct {
	API           *DerivAPI
	Stream        chan Resp
	reqID         int
	isActive      bool
	SubsciptionID string
	statusLock    sync.Mutex
}

type SubscriptionResponse struct {
	Error        APIError               `json:"error"`
	Subscription SubscriptionIDResponse `json:"subscription,omitempty"`
}

type SubscriptionIDResponse struct {
	ID string `json:"id"`
}

// parseSubscription parses a subscription-related API response in JSON format.
// It deserializes the response string into a SubscriptionResponse struct and
// returns it along with any error that occurs during the deserialization process.
// If the response contains an error code, it returns a SubscriptionResponse
// struct and an error that wraps the Error field of the struct.
func parseSubsciption(rawResponse string) (SubscriptionResponse, error) {
	var sub SubscriptionResponse

	err := json.Unmarshal([]byte(rawResponse), &sub)
	if err != nil {
		return sub, err
	}

	if sub.Error.Code != "" {
		return sub, &sub.Error
	}

	if sub.Subscription.ID == "" {
		return sub, fmt.Errorf("subscription ID is empty")
	}

	return sub, nil
}

// NewSubscription creates and returns a new Subscription instance with the given DerivAPI client.
// The Subscription instance has a Stream channel that will receive subscription updates, and an
// IsActive boolean that is set to false initially.
func NewSubcription[initResp any, Resp any](api *DerivAPI) *Subsciption[initResp, Resp] {
	return &Subsciption[initResp, Resp]{
		API:      api,
		Stream:   make(chan Resp, 1),
		isActive: false,
	}
}

// Forget cancels an active subscription by sending a Forget request to the API using the DerivAPI.Forget method.
// If the subscription is not currently active, this method does nothing. If an error occurs while sending the Forget
// request, it returns the error.
func (s *Subsciption[initResp, Resp]) Forget() error {
	s.statusLock.Lock()
	defer s.statusLock.Unlock()

	if s.isActive {
		_, err := s.API.Forget(Forget{Forget: s.SubsciptionID})

		if err != nil {
			return err
		}

		s.isActive = false

		s.API.closeRequestChannel(s.reqID)

		return nil
	}

	return nil
}

// Start starts a new subscription by sending a subscription request to the API using the DerivAPI.Send method.
// If an error occurs while sending the subscription request, it returns the error. If the subscription request is
// successful, the method receives the initial response from the API on a channel, and then calls the parseSubscription
// function to deserialize the response into a SubscriptionResponse struct. If an error occurs during deserialization,
// the method returns the error. If deserialization is successful, the SubsciptionID field of the Subscription instance
// is set to the subscription ID returned by the API, and the IsActive field is set to true. The method then sends the
// initial subscription update to the Stream channel, and starts a new goroutine to handle subsequent updates received
// on the channel. If the response object does not implement the ApiResponse interface, the method returns an error
func (s *Subsciption[initResp, Resp]) Start(reqID int, request any) (initResp, error) {
	var resp initResp

	s.statusLock.Lock()
	defer s.statusLock.Unlock()

	if s.isActive {
		return resp, nil
	}

	inChan, err := s.API.Send(reqID, request)

	if err != nil {
		return resp, err
	}

	select {
	case <-time.After(s.API.TimeOut):
		return resp, fmt.Errorf("timeout")
	case initResponse, ok := <-inChan:
		if !ok {
			return resp, fmt.Errorf("connection closed")
		}
		subResp, err := parseSubsciption(initResponse)
		if err != nil {
			s.API.closeRequestChannel(reqID)
			return resp, err
		}
		s.SubsciptionID = subResp.Subscription.ID

		apiResp, ok := any(&resp).(ApiResponse)
		if !ok {
			panic("Response object must implement ApiResponse interface")
		}

		err = apiResp.UnmarshalJSON([]byte(initResponse))
		if err != nil {
			s.API.closeRequestChannel(reqID)
			return resp, err
		}

		s.isActive = true
		s.reqID = reqID

		go s.messageHandler(inChan)

		return resp, nil
	}
}

// messageHandler is a goroutine that handles subscription updates received on the channel passed to it.
func (s *Subsciption[initResp, Resp]) messageHandler(inChan chan string) {
	defer func() {
		s.statusLock.Lock()
		if s.isActive {
			s.isActive = false
			close(s.Stream)
		}
		s.statusLock.Unlock()
	}()
	for rawResponse := range inChan {
		err := parseError(rawResponse)
		if err != nil {
			log.Printf("Error in subsciption message: %v", err)
			continue
		}

		var response Resp
		apiResp, ok := any(&response).(ApiResponse)
		if !ok {
			panic("Response object must implement ApiResponse interface")
		}

		err = apiResp.UnmarshalJSON([]byte(rawResponse))
		if err != nil {
			log.Printf("Error in subsciption message: %v", err)
			continue
		}
		s.statusLock.Lock()
		s.Stream <- response
		s.statusLock.Unlock()
	}
}

// GetStream returns the Stream channel of the Subscription instance.
func (s *Subsciption[initResp, Resp]) GetStream() chan Resp {
	return s.Stream
}

// IsActive returns the IsActive field of the Subscription instance.
func (s *Subsciption[initResp, Resp]) IsActive() bool {
	s.statusLock.Lock()
	defer s.statusLock.Unlock()
	return s.isActive
}
