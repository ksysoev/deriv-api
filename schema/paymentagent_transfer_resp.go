// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "fmt"
import "reflect"
import "encoding/json"

// Echo of the request made.
type PaymentagentTransferRespEchoReq map[string]interface{}

type PaymentagentTransferRespMsgType string

var enumValues_PaymentagentTransferRespMsgType = []interface{}{
	"paymentagent_transfer",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PaymentagentTransferRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_PaymentagentTransferRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_PaymentagentTransferRespMsgType, v)
	}
	*j = PaymentagentTransferRespMsgType(v)
	return nil
}

const PaymentagentTransferRespMsgTypePaymentagentTransfer PaymentagentTransferRespMsgType = "paymentagent_transfer"

type PaymentagentTransferRespPaymentagentTransfer int

var enumValues_PaymentagentTransferRespPaymentagentTransfer = []interface{}{
	1,
	2,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PaymentagentTransferRespPaymentagentTransfer) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_PaymentagentTransferRespPaymentagentTransfer {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_PaymentagentTransferRespPaymentagentTransfer, v)
	}
	*j = PaymentagentTransferRespPaymentagentTransfer(v)
	return nil
}

// The result of transfer request made.
type PaymentagentTransferResp struct {
	// The `transfer_to` client full name
	ClientToFullName *string `json:"client_to_full_name,omitempty"`

	// The `transfer_to` client loginid
	ClientToLoginid *string `json:"client_to_loginid,omitempty"`

	// Echo of the request made.
	EchoReq PaymentagentTransferRespEchoReq `json:"echo_req"`

	// Action name of the request made.
	MsgType PaymentagentTransferRespMsgType `json:"msg_type"`

	// If set to `1`, transfer success. If set to `2`, dry-run success.
	PaymentagentTransfer *PaymentagentTransferRespPaymentagentTransfer `json:"paymentagent_transfer,omitempty"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`

	// Reference ID of transfer performed
	TransactionId *int `json:"transaction_id,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PaymentagentTransferResp) UnmarshalJSON(b []byte) error {
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
	type Plain PaymentagentTransferResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = PaymentagentTransferResp(plain)
	return nil
}