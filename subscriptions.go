package deriv

import (
	"encoding/json"
	"fmt"
	"log"
)

type Subsciption[Resp any] struct {
	API           *DerivAPI
	Stream        chan Resp
	IsActive      bool
	SubsciptionID string
}

type SubscriptionResponse struct {
	Error        APIError `json:"error"`
	Subscription struct {
		ID string `json:"id"`
	} `json:"subscription,omitempty"`
}

func ParseSubsciption(rawResponse string) (SubscriptionResponse, error) {
	var sub SubscriptionResponse

	err := json.Unmarshal([]byte(rawResponse), &sub)
	if err != nil {
		return sub, err
	}

	if sub.Error.Code != "" {
		return sub, &sub.Error
	}

	return sub, nil
}

func NewSubcription[Resp any](api *DerivAPI) *Subsciption[Resp] {
	return &Subsciption[Resp]{
		API:      api,
		Stream:   make(chan Resp, 1),
		IsActive: false,
	}
}

func (s *Subsciption[Resp]) Forget() error {
	if s.IsActive {
		_, err := s.API.Forget(Forget{Forget: s.SubsciptionID})

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Subsciption[Resp]) Start(reqID int, request any) error {
	if s.IsActive {
		return nil
	}

	inChan, err := s.API.Send(reqID, request)

	if err != nil {
		return err
	}

	initResponse := <-inChan

	subResp, err := ParseSubsciption(initResponse)
	if err != nil {
		close(inChan)
		delete(s.API.responseMap, reqID)
		return err
	}
	s.SubsciptionID = subResp.Subscription.ID

	var response Resp
	apiResp, ok := any(&response).(ApiObjectResponse)
	if !ok {
		return fmt.Errorf("response object must implement ApiObjectResponse")
	}

	err = apiResp.UnmarshalJSON([]byte(initResponse))
	if err != nil {
		close(inChan)
		delete(s.API.responseMap, reqID)
		return err
	}

	s.IsActive = true

	s.Stream <- response

	go s.messageHandler(inChan)

	return nil
}

func (s *Subsciption[Resp]) messageHandler(inChan chan string) {
	defer func() {
		close(s.Stream)
		s.IsActive = false
	}()

	for rawResponse := range inChan {
		err := ParseError(rawResponse)
		if err != nil {
			log.Printf("Error in subsciption message: %v", err)
			continue
		}

		var response Resp
		apiResp, ok := any(&response).(ApiObjectResponse)
		if !ok {
			log.Fatal("Response object must implement ApiObjectResponse")
			return
		}

		err = apiResp.UnmarshalJSON([]byte(rawResponse))
		if err != nil {
			log.Printf("Error in subsciption message: %v", err)
			continue
		}
		s.Stream <- response
	}
}
