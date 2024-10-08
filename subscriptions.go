package deriv

// This package provides functionality for working with subscriptions using the Deriv API.

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/ksysoev/deriv-api/schema"
)

// Subscription represents a subscription instance.
type Subsciption[initResp any, Resp any] struct {
	ctx           context.Context
	API           *Client
	Stream        chan Resp
	SubsciptionID string
	reqID         int
	statusLock    sync.Mutex
	isActive      bool
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
func parseSubsciption(rawResponse []byte) (SubscriptionResponse, error) {
	var sub SubscriptionResponse

	err := json.Unmarshal(rawResponse, &sub)
	if err != nil {
		return sub, err
	}

	if sub.Error.Code != "" {
		return sub, &sub.Error
	}

	if sub.Subscription.ID == "" {
		return sub, ErrEmptySubscriptionID
	}

	return sub, nil
}

// NewSubscription creates and returns a new Subscription instance with the given DerivAPI client.
// The Subscription instance has a Stream channel that will receive subscription updates, and an
// IsActive boolean that is set to false initially.
func NewSubcription[initResp any, Resp any](ctx context.Context, api *Client) *Subsciption[initResp, Resp] {
	return &Subsciption[initResp, Resp]{
		API:      api,
		Stream:   make(chan Resp, 1),
		isActive: false,
		ctx:      ctx,
	}
}

// Forget cancels an active subscription by sending a Forget request to the API using the DerivAPI.Forget method.
// If the subscription is not currently active, this method does nothing. If an error occurs while sending the Forget
// request, it returns the error.
func (s *Subsciption[initResp, Resp]) Forget() error {
	s.statusLock.Lock()
	defer s.statusLock.Unlock()

	if s.isActive {
		_, err := s.API.Forget(s.ctx, schema.Forget{Forget: s.SubsciptionID})

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
// on the channel.
func (s *Subsciption[initResp, Resp]) Start(reqID int, request any) (initResp, error) {
	var resp initResp

	s.statusLock.Lock()
	defer s.statusLock.Unlock()

	if s.isActive {
		return resp, nil
	}

	inChan, err := s.API.Send(s.ctx, reqID, request)

	if err != nil {
		return resp, err
	}

	select {
	case <-s.ctx.Done():
		s.API.logDebugf("Timeout waiting for response for request %d", reqID)
		s.API.closeRequestChannel(reqID)

		return resp, s.ctx.Err()
	case initResponse, ok := <-inChan:
		if !ok {
			s.API.logDebugf("Connection closed while waiting for response for request %d", reqID)

			return resp, ErrConnectionClosed
		}

		subResp, err := parseSubsciption(initResponse)
		if err != nil {
			s.API.closeRequestChannel(reqID)

			return resp, err
		}

		s.SubsciptionID = subResp.Subscription.ID

		err = json.Unmarshal(initResponse, &resp)
		if err != nil {
			s.API.logDebugf("Failed to parse response for request %d: %s", reqID, err.Error())
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
func (s *Subsciption[initResp, Resp]) messageHandler(inChan chan []byte) {
	defer func() {
		s.API.closeRequestChannel(s.reqID)

		s.statusLock.Lock()
		if s.isActive {
			s.isActive = false
			close(s.Stream)
		}
		s.statusLock.Unlock()
	}()

	for {
		select {
		case <-s.ctx.Done():
			s.API.logDebugf("Subscription %s closed", s.SubsciptionID)
			return
		case rawResponse, ok := <-inChan:
			if !ok {
				s.API.logDebugf("Connection closed while waiting for response for request %d", s.reqID)
				return
			}

			if err := parseError(rawResponse); err != nil {
				s.API.logDebugf("Error in subsciption message: %v", err)
				continue
			}

			var response Resp

			if err := json.Unmarshal(rawResponse, &response); err != nil {
				s.API.logDebugf("Failed to parse response in subscription: %s", err.Error())
				continue
			}

			s.statusLock.Lock()
			select {
			case <-s.ctx.Done():
			case s.Stream <- response:
			}
			s.statusLock.Unlock()
		}
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
