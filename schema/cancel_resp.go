// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "fmt"
import "reflect"
import "encoding/json"

// Receipt for the transaction
type CancelRespCancel struct {
	// New account balance after completion of the sale
	BalanceAfter *float64 `json:"balance_after,omitempty"`

	// Internal contract identifier for the sold contract
	ContractId *int `json:"contract_id,omitempty"`

	// Internal transaction identifier for the corresponding buy transaction
	ReferenceId *int `json:"reference_id,omitempty"`

	// Actual effected sale price
	SoldFor *float64 `json:"sold_for,omitempty"`

	// Internal transaction identifier for the sale transaction
	TransactionId *int `json:"transaction_id,omitempty"`
}

// Echo of the request made.
type CancelRespEchoReq map[string]interface{}

type CancelRespMsgType string

var enumValues_CancelRespMsgType = []interface{}{
	"cancel",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CancelRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CancelRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CancelRespMsgType, v)
	}
	*j = CancelRespMsgType(v)
	return nil
}

// A message with transaction results is received
type CancelResp struct {
	// Receipt for the transaction
	Cancel *CancelRespCancel `json:"cancel,omitempty"`

	// Echo of the request made.
	EchoReq CancelRespEchoReq `json:"echo_req"`

	// Action name of the request made.
	MsgType CancelRespMsgType `json:"msg_type"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`
}

const CancelRespMsgTypeCancel CancelRespMsgType = "cancel"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CancelResp) UnmarshalJSON(b []byte) error {
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
	type Plain CancelResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CancelResp(plain)
	return nil
}