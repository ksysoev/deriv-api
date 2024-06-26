// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// Verify Email Cellxpert Receive
type VerifyEmailCellxpertResp struct {
	// Echo of the request made.
	EchoReq VerifyEmailCellxpertRespEchoReq `json:"echo_req"`

	// Action name of the request made.
	MsgType VerifyEmailCellxpertRespMsgType `json:"msg_type"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`

	// 1 for success (secure code has been sent to the email address)
	VerifyEmailCellxpert *VerifyEmailCellxpertRespVerifyEmailCellxpert `json:"verify_email_cellxpert,omitempty"`
}

// Echo of the request made.
type VerifyEmailCellxpertRespEchoReq map[string]interface{}

type VerifyEmailCellxpertRespMsgType string

const VerifyEmailCellxpertRespMsgTypeVerifyEmailCellxpert VerifyEmailCellxpertRespMsgType = "verify_email_cellxpert"

var enumValues_VerifyEmailCellxpertRespMsgType = []interface{}{
	"verify_email_cellxpert",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VerifyEmailCellxpertRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_VerifyEmailCellxpertRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_VerifyEmailCellxpertRespMsgType, v)
	}
	*j = VerifyEmailCellxpertRespMsgType(v)
	return nil
}

type VerifyEmailCellxpertRespVerifyEmailCellxpert int

var enumValues_VerifyEmailCellxpertRespVerifyEmailCellxpert = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VerifyEmailCellxpertRespVerifyEmailCellxpert) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_VerifyEmailCellxpertRespVerifyEmailCellxpert {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_VerifyEmailCellxpertRespVerifyEmailCellxpert, v)
	}
	*j = VerifyEmailCellxpertRespVerifyEmailCellxpert(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VerifyEmailCellxpertResp) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["echo_req"]; raw != nil && !ok {
		return fmt.Errorf("field echo_req in VerifyEmailCellxpertResp: required")
	}
	if _, ok := raw["msg_type"]; raw != nil && !ok {
		return fmt.Errorf("field msg_type in VerifyEmailCellxpertResp: required")
	}
	type Plain VerifyEmailCellxpertResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = VerifyEmailCellxpertResp(plain)
	return nil
}
