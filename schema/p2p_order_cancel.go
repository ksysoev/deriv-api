// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// Cancel a P2P order.
type P2POrderCancel struct {
	// The unique identifier for this order.
	Id string `json:"id"`

	// [Optional] The login id of the user. Mandatory when multiple tokens were
	// provided during authorize.
	Loginid *string `json:"loginid,omitempty"`

	// Must be 1
	P2POrderCancel P2POrderCancelP2POrderCancel `json:"p2p_order_cancel"`

	// [Optional] Used to pass data through the websocket, which may be retrieved via
	// the `echo_req` output field.
	Passthrough P2POrderCancelPassthrough `json:"passthrough,omitempty"`

	// [Optional] Used to map request to response.
	ReqId *int `json:"req_id,omitempty"`
}

type P2POrderCancelP2POrderCancel int

var enumValues_P2POrderCancelP2POrderCancel = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2POrderCancelP2POrderCancel) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2POrderCancelP2POrderCancel {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2POrderCancelP2POrderCancel, v)
	}
	*j = P2POrderCancelP2POrderCancel(v)
	return nil
}

// [Optional] Used to pass data through the websocket, which may be retrieved via
// the `echo_req` output field.
type P2POrderCancelPassthrough map[string]interface{}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2POrderCancel) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["id"]; raw != nil && !ok {
		return fmt.Errorf("field id in P2POrderCancel: required")
	}
	if _, ok := raw["p2p_order_cancel"]; raw != nil && !ok {
		return fmt.Errorf("field p2p_order_cancel in P2POrderCancel: required")
	}
	type Plain P2POrderCancel
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = P2POrderCancel(plain)
	return nil
}
