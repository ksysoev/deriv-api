// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "fmt"
import "reflect"
import "encoding/json"

// Echo of the request made.
type TimeRespEchoReq map[string]interface{}

type TimeRespMsgType string

var enumValues_TimeRespMsgType = []interface{}{
	"time",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TimeRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TimeRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TimeRespMsgType, v)
	}
	*j = TimeRespMsgType(v)
	return nil
}

// The result of server time request.
type TimeResp struct {
	// Echo of the request made.
	EchoReq TimeRespEchoReq `json:"echo_req"`

	// Action name of the request made.
	MsgType TimeRespMsgType `json:"msg_type"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`

	// Epoch of server time.
	Time *int `json:"time,omitempty"`
}

const TimeRespMsgTypeTime TimeRespMsgType = "time"

// UnmarshalJSON implements json.Unmarshaler.
func (j *TimeResp) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["echo_req"]; !ok || v == nil {
		return fmt.Errorf("field echo_req: required")
	}
	if v, ok := raw["msg_type"]; !ok || v == nil {
		return fmt.Errorf("field msg_type: required")
	}
	type Plain TimeResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = TimeResp(plain)
	return nil
}