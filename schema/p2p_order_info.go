// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// Retrieves the information about a P2P order.
type P2POrderInfo struct {
	// The unique identifier for the order.
	Id string `json:"id"`

	// [Optional] The login id of the user. Mandatory when multiple tokens were
	// provided during authorize.
	Loginid *string `json:"loginid,omitempty"`

	// Must be 1
	P2POrderInfo P2POrderInfoP2POrderInfo `json:"p2p_order_info"`

	// [Optional] Used to pass data through the websocket, which may be retrieved via
	// the `echo_req` output field.
	Passthrough P2POrderInfoPassthrough `json:"passthrough,omitempty"`

	// [Optional] Used to map request to response.
	ReqId *int `json:"req_id,omitempty"`

	// [Optional] If set to 1, will send updates whenever there is an update to order
	Subscribe *P2POrderInfoSubscribe `json:"subscribe,omitempty"`
}

type P2POrderInfoP2POrderInfo int

var enumValues_P2POrderInfoP2POrderInfo = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2POrderInfoP2POrderInfo) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2POrderInfoP2POrderInfo {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2POrderInfoP2POrderInfo, v)
	}
	*j = P2POrderInfoP2POrderInfo(v)
	return nil
}

// [Optional] Used to pass data through the websocket, which may be retrieved via
// the `echo_req` output field.
type P2POrderInfoPassthrough map[string]interface{}

type P2POrderInfoSubscribe int

var enumValues_P2POrderInfoSubscribe = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2POrderInfoSubscribe) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2POrderInfoSubscribe {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2POrderInfoSubscribe, v)
	}
	*j = P2POrderInfoSubscribe(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2POrderInfo) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["id"]; raw != nil && !ok {
		return fmt.Errorf("field id in P2POrderInfo: required")
	}
	if _, ok := raw["p2p_order_info"]; raw != nil && !ok {
		return fmt.Errorf("field p2p_order_info in P2POrderInfo: required")
	}
	type Plain P2POrderInfo
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = P2POrderInfo(plain)
	return nil
}
