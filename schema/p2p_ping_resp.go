// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// The response of P2P ping request.
type P2PPingResp struct {
	// Echo of the request made.
	EchoReq P2PPingRespEchoReq `json:"echo_req"`

	// Action name of the request made.
	MsgType P2PPingRespMsgType `json:"msg_type"`

	// Will return 'pong'
	P2PPing *P2PPingRespP2PPing `json:"p2p_ping,omitempty"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`
}

// Echo of the request made.
type P2PPingRespEchoReq map[string]interface{}

type P2PPingRespMsgType string

const P2PPingRespMsgTypeP2PPing P2PPingRespMsgType = "p2p_ping"

var enumValues_P2PPingRespMsgType = []interface{}{
	"p2p_ping",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PPingRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PPingRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PPingRespMsgType, v)
	}
	*j = P2PPingRespMsgType(v)
	return nil
}

type P2PPingRespP2PPing string

const P2PPingRespP2PPingPong P2PPingRespP2PPing = "pong"

var enumValues_P2PPingRespP2PPing = []interface{}{
	"pong",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PPingRespP2PPing) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PPingRespP2PPing {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PPingRespP2PPing, v)
	}
	*j = P2PPingRespP2PPing(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PPingResp) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["echo_req"]; raw != nil && !ok {
		return fmt.Errorf("field echo_req in P2PPingResp: required")
	}
	if _, ok := raw["msg_type"]; raw != nil && !ok {
		return fmt.Errorf("field msg_type in P2PPingResp: required")
	}
	type Plain P2PPingResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = P2PPingResp(plain)
	return nil
}
