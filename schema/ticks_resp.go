// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "fmt"
import "reflect"
import "encoding/json"

// Echo of the request made.
type TicksRespEchoReq map[string]interface{}

type TicksRespMsgType string

var enumValues_TicksRespMsgType = []interface{}{
	"tick",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TicksRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TicksRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TicksRespMsgType, v)
	}
	*j = TicksRespMsgType(v)
	return nil
}

const TicksRespMsgTypeTick TicksRespMsgType = "tick"

// For subscription requests only.
type TicksRespSubscription struct {
	// A per-connection unique identifier. Can be passed to the `forget` API call to
	// unsubscribe.
	Id string `json:"id"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TicksRespSubscription) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain TicksRespSubscription
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = TicksRespSubscription(plain)
	return nil
}

// Tick by tick list of streamed data
type TicksRespTick struct {
	// Market ask at the epoch
	Ask *float64 `json:"ask,omitempty"`

	// Market bid at the epoch
	Bid *float64 `json:"bid,omitempty"`

	// Epoch time of the tick
	Epoch *int `json:"epoch,omitempty"`

	// A per-connection unique identifier. Can be passed to the `forget` API call to
	// unsubscribe.
	Id *string `json:"id,omitempty"`

	// Indicates the number of decimal points that the returned amounts must be
	// displayed with
	PipSize float64 `json:"pip_size"`

	// Market value at the epoch
	Quote *float64 `json:"quote,omitempty"`

	// Symbol
	Symbol *string `json:"symbol,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TicksRespTick) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["pip_size"]; !ok || v == nil {
		return fmt.Errorf("field pip_size: required")
	}
	type Plain TicksRespTick
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = TicksRespTick(plain)
	return nil
}

// Latest spot price for a given symbol. Continuous responses with a frequency of
// up to one second.
type TicksResp struct {
	// Echo of the request made.
	EchoReq TicksRespEchoReq `json:"echo_req"`

	// Type of the response.
	MsgType TicksRespMsgType `json:"msg_type"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`

	// For subscription requests only.
	Subscription *TicksRespSubscription `json:"subscription,omitempty"`

	// Tick by tick list of streamed data
	Tick *TicksRespTick `json:"tick,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TicksResp) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["echo_req"]; !ok || v == nil {
		return fmt.Errorf("field echo_req: required")
	}
	if v, ok := raw["msg_type"]; !ok || v == nil {
		return fmt.Errorf("field msg_type: required")
	}
	type Plain TicksResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = TicksResp(plain)
	return nil
}