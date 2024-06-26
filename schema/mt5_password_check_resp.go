// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// MT5 user password check receive
type Mt5PasswordCheckResp struct {
	// Echo of the request made.
	EchoReq Mt5PasswordCheckRespEchoReq `json:"echo_req"`

	// Action name of the request made.
	MsgType Mt5PasswordCheckRespMsgType `json:"msg_type"`

	// `1` on success
	Mt5PasswordCheck *int `json:"mt5_password_check,omitempty"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`
}

// Echo of the request made.
type Mt5PasswordCheckRespEchoReq map[string]interface{}

type Mt5PasswordCheckRespMsgType string

const Mt5PasswordCheckRespMsgTypeMt5PasswordCheck Mt5PasswordCheckRespMsgType = "mt5_password_check"

var enumValues_Mt5PasswordCheckRespMsgType = []interface{}{
	"mt5_password_check",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Mt5PasswordCheckRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_Mt5PasswordCheckRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_Mt5PasswordCheckRespMsgType, v)
	}
	*j = Mt5PasswordCheckRespMsgType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Mt5PasswordCheckResp) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["echo_req"]; raw != nil && !ok {
		return fmt.Errorf("field echo_req in Mt5PasswordCheckResp: required")
	}
	if _, ok := raw["msg_type"]; raw != nil && !ok {
		return fmt.Errorf("field msg_type in Mt5PasswordCheckResp: required")
	}
	type Plain Mt5PasswordCheckResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Mt5PasswordCheckResp(plain)
	return nil
}
