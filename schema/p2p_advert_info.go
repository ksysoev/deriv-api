// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "fmt"
import "reflect"
import "encoding/json"

type P2PAdvertInfoP2PAdvertInfo int

var enumValues_P2PAdvertInfoP2PAdvertInfo = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertInfoP2PAdvertInfo) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertInfoP2PAdvertInfo {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertInfoP2PAdvertInfo, v)
	}
	*j = P2PAdvertInfoP2PAdvertInfo(v)
	return nil
}

// [Optional] Used to pass data through the websocket, which may be retrieved via
// the `echo_req` output field. Maximum size is 3500 bytes.
type P2PAdvertInfoPassthrough map[string]interface{}

type P2PAdvertInfoSubscribe int

var enumValues_P2PAdvertInfoSubscribe = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertInfoSubscribe) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertInfoSubscribe {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertInfoSubscribe, v)
	}
	*j = P2PAdvertInfoSubscribe(v)
	return nil
}

type P2PAdvertInfoUseClientLimits int

var enumValues_P2PAdvertInfoUseClientLimits = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertInfoUseClientLimits) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertInfoUseClientLimits {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertInfoUseClientLimits, v)
	}
	*j = P2PAdvertInfoUseClientLimits(v)
	return nil
}

// Retrieve information about a P2P advert.
type P2PAdvertInfo struct {
	// [Optional] The unique identifier for this advert. Optional when subscribe is 1.
	// If not provided, all advertiser adverts will be subscribed.
	Id *string `json:"id,omitempty"`

	// Must be 1
	P2PAdvertInfo P2PAdvertInfoP2PAdvertInfo `json:"p2p_advert_info"`

	// [Optional] Used to pass data through the websocket, which may be retrieved via
	// the `echo_req` output field. Maximum size is 3500 bytes.
	Passthrough P2PAdvertInfoPassthrough `json:"passthrough,omitempty"`

	// [Optional] Used to map request to response.
	ReqId *int `json:"req_id,omitempty"`

	// [Optional] If set to 1, will send updates when changes occur. Optional when id
	// is provided.
	Subscribe *P2PAdvertInfoSubscribe `json:"subscribe,omitempty"`

	// [Optional] If set to 1, the maximum order amount will be adjusted to the
	// current balance and turnover limits of the account.
	UseClientLimits P2PAdvertInfoUseClientLimits `json:"use_client_limits,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertInfo) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["p2p_advert_info"]; !ok || v == nil {
		return fmt.Errorf("field p2p_advert_info: required")
	}
	type Plain P2PAdvertInfo
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["use_client_limits"]; !ok || v == nil {
		plain.UseClientLimits = 0
	}
	*j = P2PAdvertInfo(plain)
	return nil
}