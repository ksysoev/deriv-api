// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// Request P2P Settings information.
type P2PSettings struct {
	// [Optional] The login id of the user. Mandatory when multiple tokens were
	// provided during authorize.
	Loginid *string `json:"loginid,omitempty"`

	// Must be `1`
	P2PSettings P2PSettingsP2PSettings `json:"p2p_settings"`

	// [Optional] Used to pass data through the websocket, which may be retrieved via
	// the `echo_req` output field.
	Passthrough P2PSettingsPassthrough `json:"passthrough,omitempty"`

	// [Optional] Used to map request to response.
	ReqId *int `json:"req_id,omitempty"`

	// [Optional] If set to `1`, will send updates whenever there is an update to P2P
	// settings.
	Subscribe *P2PSettingsSubscribe `json:"subscribe,omitempty"`
}

type P2PSettingsP2PSettings int

var enumValues_P2PSettingsP2PSettings = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PSettingsP2PSettings) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PSettingsP2PSettings {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PSettingsP2PSettings, v)
	}
	*j = P2PSettingsP2PSettings(v)
	return nil
}

// [Optional] Used to pass data through the websocket, which may be retrieved via
// the `echo_req` output field.
type P2PSettingsPassthrough map[string]interface{}

type P2PSettingsSubscribe int

var enumValues_P2PSettingsSubscribe = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PSettingsSubscribe) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PSettingsSubscribe {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PSettingsSubscribe, v)
	}
	*j = P2PSettingsSubscribe(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PSettings) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["p2p_settings"]; raw != nil && !ok {
		return fmt.Errorf("field p2p_settings in P2PSettings: required")
	}
	type Plain P2PSettings
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = P2PSettings(plain)
	return nil
}
