// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// The result of the unsubscribe email request.
type UnsubscribeEmailResp struct {
	// Customer User ID.
	BinaryUserId *float64 `json:"binary_user_id,omitempty"`

	// Echo of the request made.
	EchoReq UnsubscribeEmailRespEchoReq `json:"echo_req"`

	// `1`: email notification unsubscribed sucssesfully, `0`: failed to unsubscribed
	// email notification
	EmailUnsubscribeStatus *UnsubscribeEmailRespEmailUnsubscribeStatus `json:"email_unsubscribe_status,omitempty"`

	// Action name of the request made.
	MsgType UnsubscribeEmailRespMsgType `json:"msg_type"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`
}

// Echo of the request made.
type UnsubscribeEmailRespEchoReq map[string]interface{}

type UnsubscribeEmailRespEmailUnsubscribeStatus int

var enumValues_UnsubscribeEmailRespEmailUnsubscribeStatus = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UnsubscribeEmailRespEmailUnsubscribeStatus) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_UnsubscribeEmailRespEmailUnsubscribeStatus {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_UnsubscribeEmailRespEmailUnsubscribeStatus, v)
	}
	*j = UnsubscribeEmailRespEmailUnsubscribeStatus(v)
	return nil
}

type UnsubscribeEmailRespMsgType string

const UnsubscribeEmailRespMsgTypeUnsubscribeEmail UnsubscribeEmailRespMsgType = "unsubscribe_email"

var enumValues_UnsubscribeEmailRespMsgType = []interface{}{
	"unsubscribe_email",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UnsubscribeEmailRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_UnsubscribeEmailRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_UnsubscribeEmailRespMsgType, v)
	}
	*j = UnsubscribeEmailRespMsgType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UnsubscribeEmailResp) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["echo_req"]; raw != nil && !ok {
		return fmt.Errorf("field echo_req in UnsubscribeEmailResp: required")
	}
	if _, ok := raw["msg_type"]; raw != nil && !ok {
		return fmt.Errorf("field msg_type in UnsubscribeEmailResp: required")
	}
	type Plain UnsubscribeEmailResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = UnsubscribeEmailResp(plain)
	return nil
}
