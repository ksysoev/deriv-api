// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "fmt"
import "reflect"
import "encoding/json"

type GetSelfExclusionGetSelfExclusion int

var enumValues_GetSelfExclusionGetSelfExclusion = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetSelfExclusionGetSelfExclusion) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GetSelfExclusionGetSelfExclusion {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GetSelfExclusionGetSelfExclusion, v)
	}
	*j = GetSelfExclusionGetSelfExclusion(v)
	return nil
}

// Allows users to exclude themselves from the website for certain periods of time,
// or to set limits on their trading activities. This facility is a regulatory
// requirement for certain Landing Companies.
type GetSelfExclusion struct {
	// Must be `1`
	GetSelfExclusion GetSelfExclusionGetSelfExclusion `json:"get_self_exclusion"`

	// [Optional] Used to pass data through the websocket, which may be retrieved via
	// the `echo_req` output field. Maximum size is 3500 bytes.
	Passthrough GetSelfExclusionPassthrough `json:"passthrough,omitempty"`

	// [Optional] Used to map request to response.
	ReqId *int `json:"req_id,omitempty"`
}

// [Optional] Used to pass data through the websocket, which may be retrieved via
// the `echo_req` output field. Maximum size is 3500 bytes.
type GetSelfExclusionPassthrough map[string]interface{}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetSelfExclusion) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["get_self_exclusion"]; !ok || v == nil {
		return fmt.Errorf("field get_self_exclusion: required")
	}
	type Plain GetSelfExclusion
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetSelfExclusion(plain)
	return nil
}