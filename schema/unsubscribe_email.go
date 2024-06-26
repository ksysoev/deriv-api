// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// It unsubscribe user from the email subscription.
type UnsubscribeEmail struct {
	// Customer User ID.
	BinaryUserId float64 `json:"binary_user_id"`

	// The generated checksum for the customer.
	Checksum string `json:"checksum"`

	// [Optional] Used to pass data through the websocket, which may be retrieved via
	// the `echo_req` output field.
	Passthrough UnsubscribeEmailPassthrough `json:"passthrough,omitempty"`

	// [Optional] Used to map request to response.
	ReqId *int `json:"req_id,omitempty"`

	// Must be `1`
	UnsubscribeEmail UnsubscribeEmailUnsubscribeEmail `json:"unsubscribe_email"`
}

// [Optional] Used to pass data through the websocket, which may be retrieved via
// the `echo_req` output field.
type UnsubscribeEmailPassthrough map[string]interface{}

type UnsubscribeEmailUnsubscribeEmail int

var enumValues_UnsubscribeEmailUnsubscribeEmail = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UnsubscribeEmailUnsubscribeEmail) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_UnsubscribeEmailUnsubscribeEmail {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_UnsubscribeEmailUnsubscribeEmail, v)
	}
	*j = UnsubscribeEmailUnsubscribeEmail(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UnsubscribeEmail) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["binary_user_id"]; raw != nil && !ok {
		return fmt.Errorf("field binary_user_id in UnsubscribeEmail: required")
	}
	if _, ok := raw["checksum"]; raw != nil && !ok {
		return fmt.Errorf("field checksum in UnsubscribeEmail: required")
	}
	if _, ok := raw["unsubscribe_email"]; raw != nil && !ok {
		return fmt.Errorf("field unsubscribe_email in UnsubscribeEmail: required")
	}
	type Plain UnsubscribeEmail
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = UnsubscribeEmail(plain)
	return nil
}
