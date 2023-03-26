package deriv

// This package provides functionality for working with subscriptions using the Deriv API.

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// Subscription represents a subscription instance.
type Subsciption[Resp any] struct {
	API           *DerivAPI
	Stream        chan Resp
	reqID         int
	IsActive      bool
	SubsciptionID string
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
func NewSubcription[Resp any](api *DerivAPI) *Subsciption[Resp] {
	return &Subsciption[Resp]{
		API:      api,
		Stream:   make(chan Resp, 1),
		IsActive: false,
	}
}

// Forget cancels an active subscription by sending a Forget request to the API using the DerivAPI.Forget method.
// If the subscription is not currently active, this method does nothing. If an error occurs while sending the Forget
// request, it returns the error.
func (s *Subsciption[Resp]) Forget() error {
	if s.IsActive {
		_, err := s.API.Forget(Forget{Forget: s.SubsciptionID})

		if err != nil {
			return err
		}

		s.IsActive = false

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
func (s *Subsciption[Resp]) Start(reqID int, request any) error {
	if s.IsActive {
		return nil
	}

	inChan, err := s.API.Send(reqID, request)

	if err != nil {
		return err
	}

	select {
	case <-time.After(s.API.TimeOut):
		return fmt.Errorf("timeout")
	case initResponse, ok := <-inChan:
		if !ok {
			return fmt.Errorf("connection closed")
		}
		subResp, err := parseSubsciption(initResponse)
		if err != nil {
			s.API.closeRequestChannel(reqID)
			return err
		}
		s.SubsciptionID = subResp.Subscription.ID

		var response Resp
		apiResp, ok := any(&response).(ApiResponse)
		if !ok {
			panic("Response object must implement ApiResponse interface")
		}

		err = apiResp.UnmarshalJSON([]byte(initResponse))
		if err != nil {
			s.API.closeRequestChannel(reqID)
			return err
		}

		s.IsActive = true
		s.reqID = reqID
		s.Stream <- response

		go s.messageHandler(inChan)

		return nil
	}
}

// messageHandler is a goroutine that handles subscription updates received on the channel passed to it.
func (s *Subsciption[Resp]) messageHandler(inChan chan string) {
	defer func() {
		close(s.Stream)
		s.IsActive = false
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
		s.Stream <- response
	}
}
