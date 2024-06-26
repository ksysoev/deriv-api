// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// The result of virtual money top up
type TopupVirtualResp struct {
	// Echo of the request made.
	EchoReq TopupVirtualRespEchoReq `json:"echo_req"`

	// Action name of the request made.
	MsgType TopupVirtualRespMsgType `json:"msg_type"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`

	// The information regarding a successful top up for a virtual money account
	TopupVirtual *TopupVirtualRespTopupVirtual `json:"topup_virtual,omitempty"`
}

// Echo of the request made.
type TopupVirtualRespEchoReq map[string]interface{}

type TopupVirtualRespMsgType string

const TopupVirtualRespMsgTypeTopupVirtual TopupVirtualRespMsgType = "topup_virtual"

var enumValues_TopupVirtualRespMsgType = []interface{}{
	"topup_virtual",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TopupVirtualRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TopupVirtualRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TopupVirtualRespMsgType, v)
	}
	*j = TopupVirtualRespMsgType(v)
	return nil
}

// The information regarding a successful top up for a virtual money account
type TopupVirtualRespTopupVirtual struct {
	// Top up amount
	Amount *float64 `json:"amount,omitempty"`

	// Top up currency string
	Currency *string `json:"currency,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TopupVirtualResp) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["echo_req"]; raw != nil && !ok {
		return fmt.Errorf("field echo_req in TopupVirtualResp: required")
	}
	if _, ok := raw["msg_type"]; raw != nil && !ok {
		return fmt.Errorf("field msg_type in TopupVirtualResp: required")
	}
	type Plain TopupVirtualResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = TopupVirtualResp(plain)
	return nil
}
