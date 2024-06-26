// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// Keeps the connection alive and updates the P2P advertiser's online status. The
// advertiser will be considered offline 60 seconds after a call is made.
type P2PPing struct {
	// [Optional] The login id of the user. Mandatory when multiple tokens were
	// provided during authorize.
	Loginid *string `json:"loginid,omitempty"`

	// Must be `1`
	P2PPing P2PPingP2PPing `json:"p2p_ping"`

	// [Optional] Used to pass data through the websocket, which may be retrieved via
	// the `echo_req` output field.
	Passthrough P2PPingPassthrough `json:"passthrough,omitempty"`

	// [Optional] Used to map request to response.
	ReqId *int `json:"req_id,omitempty"`
}

type P2PPingP2PPing int

var enumValues_P2PPingP2PPing = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PPingP2PPing) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PPingP2PPing {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PPingP2PPing, v)
	}
	*j = P2PPingP2PPing(v)
	return nil
}

// [Optional] Used to pass data through the websocket, which may be retrieved via
// the `echo_req` output field.
type P2PPingPassthrough map[string]interface{}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PPing) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["p2p_ping"]; raw != nil && !ok {
		return fmt.Errorf("field p2p_ping in P2PPing: required")
	}
	type Plain P2PPing
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = P2PPing(plain)
	return nil
}
