// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "fmt"
import "reflect"
import "encoding/json"

type P2POrderReviewP2POrderReview int

var enumValues_P2POrderReviewP2POrderReview = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2POrderReviewP2POrderReview) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2POrderReviewP2POrderReview {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2POrderReviewP2POrderReview, v)
	}
	*j = P2POrderReviewP2POrderReview(v)
	return nil
}

// [Optional] Used to pass data through the websocket, which may be retrieved via
// the `echo_req` output field. Maximum size is 3500 bytes.
type P2POrderReviewPassthrough map[string]interface{}

type P2POrderReviewRecommended struct {
	Value interface{}
}

var enumValues_P2POrderReviewRecommended = []interface{}{
	nil,
	0,
	1,
}

// MarshalJSON implements json.Marshaler.
func (j *P2POrderReviewRecommended) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2POrderReviewRecommended) UnmarshalJSON(b []byte) error {
	var v struct {
		Value interface{}
	}
	if err := json.Unmarshal(b, &v.Value); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2POrderReviewRecommended {
		if reflect.DeepEqual(v.Value, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2POrderReviewRecommended, v.Value)
	}
	*j = P2POrderReviewRecommended(v)
	return nil
}

// Creates a review for the specified order.
type P2POrderReview struct {
	// The order identification number.
	OrderId string `json:"order_id"`

	// Must be 1
	P2POrderReview P2POrderReviewP2POrderReview `json:"p2p_order_review"`

	// [Optional] Used to pass data through the websocket, which may be retrieved via
	// the `echo_req` output field. Maximum size is 3500 bytes.
	Passthrough P2POrderReviewPassthrough `json:"passthrough,omitempty"`

	// Rating for the transaction, 1 to 5.
	Rating int `json:"rating"`

	// [Optional] `1` if the counterparty is recommendable to others, otherwise `0`.
	Recommended *P2POrderReviewRecommended `json:"recommended,omitempty"`

	// [Optional] Used to map request to response.
	ReqId *int `json:"req_id,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2POrderReview) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["order_id"]; !ok || v == nil {
		return fmt.Errorf("field order_id: required")
	}
	if v, ok := raw["p2p_order_review"]; !ok || v == nil {
		return fmt.Errorf("field p2p_order_review: required")
	}
	if v, ok := raw["rating"]; !ok || v == nil {
		return fmt.Errorf("field rating: required")
	}
	type Plain P2POrderReview
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = P2POrderReview(plain)
	return nil
}