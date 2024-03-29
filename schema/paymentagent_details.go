// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// [Optional] Used to pass data through the websocket, which may be retrieved via
// the `echo_req` output field.
type PaymentagentDetailsPassthrough map[string]interface{}

type PaymentagentDetailsPaymentagentDetails int

var enumValues_PaymentagentDetailsPaymentagentDetails = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PaymentagentDetailsPaymentagentDetails) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_PaymentagentDetailsPaymentagentDetails {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_PaymentagentDetailsPaymentagentDetails, v)
	}
	*j = PaymentagentDetailsPaymentagentDetails(v)
	return nil
}

// Gets client's payment agent details.
type PaymentagentDetails struct {
	// [Optional] The login id of the user. If left unspecified, it defaults to the
	// initial authorized token's login id.
	Loginid *string `json:"loginid,omitempty"`

	// [Optional] Used to pass data through the websocket, which may be retrieved via
	// the `echo_req` output field.
	Passthrough PaymentagentDetailsPassthrough `json:"passthrough,omitempty"`

	// Must be 1
	PaymentagentDetails PaymentagentDetailsPaymentagentDetails `json:"paymentagent_details"`

	// [Optional] Used to map request to response.
	ReqId *int `json:"req_id,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PaymentagentDetails) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["paymentagent_details"]; !ok || v == nil {
		return fmt.Errorf("field paymentagent_details in PaymentagentDetails: required")
	}
	type Plain PaymentagentDetails
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = PaymentagentDetails(plain)
	return nil
}
