// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "fmt"
import "reflect"
import "encoding/json"

// Dispute a P2P order.
type P2POrderDispute struct {
	// The predefined dispute reason
	DisputeReason P2POrderDisputeDisputeReason `json:"dispute_reason"`

	// The unique identifier for this order.
	Id string `json:"id"`

	// Must be 1
	P2POrderDispute P2POrderDisputeP2POrderDispute `json:"p2p_order_dispute"`

	// [Optional] Used to pass data through the websocket, which may be retrieved via
	// the `echo_req` output field. Maximum size is 3500 bytes.
	Passthrough P2POrderDisputePassthrough `json:"passthrough,omitempty"`

	// [Optional] Used to map request to response.
	ReqId *int `json:"req_id,omitempty"`
}

type P2POrderDisputeDisputeReason string

const P2POrderDisputeDisputeReasonBuyerNotPaid P2POrderDisputeDisputeReason = "buyer_not_paid"
const P2POrderDisputeDisputeReasonBuyerOverpaid P2POrderDisputeDisputeReason = "buyer_overpaid"
const P2POrderDisputeDisputeReasonBuyerThirdPartyPaymentMethod P2POrderDisputeDisputeReason = "buyer_third_party_payment_method"
const P2POrderDisputeDisputeReasonBuyerUnderpaid P2POrderDisputeDisputeReason = "buyer_underpaid"
const P2POrderDisputeDisputeReasonSellerNotReleased P2POrderDisputeDisputeReason = "seller_not_released"

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2POrderDisputeDisputeReason) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2POrderDisputeDisputeReason {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2POrderDisputeDisputeReason, v)
	}
	*j = P2POrderDisputeDisputeReason(v)
	return nil
}

type P2POrderDisputeP2POrderDispute int

var enumValues_P2POrderDisputeP2POrderDispute = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2POrderDisputeP2POrderDispute) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2POrderDisputeP2POrderDispute {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2POrderDisputeP2POrderDispute, v)
	}
	*j = P2POrderDisputeP2POrderDispute(v)
	return nil
}

// [Optional] Used to pass data through the websocket, which may be retrieved via
// the `echo_req` output field. Maximum size is 3500 bytes.
type P2POrderDisputePassthrough map[string]interface{}

var enumValues_P2POrderDisputeDisputeReason = []interface{}{
	"seller_not_released",
	"buyer_underpaid",
	"buyer_overpaid",
	"buyer_not_paid",
	"buyer_third_party_payment_method",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2POrderDispute) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["dispute_reason"]; !ok || v == nil {
		return fmt.Errorf("field dispute_reason: required")
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	if v, ok := raw["p2p_order_dispute"]; !ok || v == nil {
		return fmt.Errorf("field p2p_order_dispute: required")
	}
	type Plain P2POrderDispute
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = P2POrderDispute(plain)
	return nil
}