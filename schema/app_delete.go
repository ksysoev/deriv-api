// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "fmt"
import "encoding/json"

// The request for deleting an application.
type AppDelete struct {
	// Application app_id
	AppDelete int `json:"app_delete"`

	// [Optional] Used to pass data through the websocket, which may be retrieved via
	// the `echo_req` output field. Maximum size is 3500 bytes.
	Passthrough AppDeletePassthrough `json:"passthrough,omitempty"`

	// [Optional] Used to map request to response.
	ReqId *int `json:"req_id,omitempty"`
}

// [Optional] Used to pass data through the websocket, which may be retrieved via
// the `echo_req` output field. Maximum size is 3500 bytes.
type AppDeletePassthrough map[string]interface{}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AppDelete) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["app_delete"]; !ok || v == nil {
		return fmt.Errorf("field app_delete: required")
	}
	type Plain AppDelete
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = AppDelete(plain)
	return nil
}