// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// Status of set account currency call
type SetAccountCurrencyResp struct {
	// Echo of the request made.
	EchoReq SetAccountCurrencyRespEchoReq `json:"echo_req"`

	// Action name of the request made.
	MsgType SetAccountCurrencyRespMsgType `json:"msg_type"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`

	// `1`: success, `0`: no change
	SetAccountCurrency *SetAccountCurrencyRespSetAccountCurrency `json:"set_account_currency,omitempty"`
}

// Echo of the request made.
type SetAccountCurrencyRespEchoReq map[string]interface{}

type SetAccountCurrencyRespMsgType string

const SetAccountCurrencyRespMsgTypeSetAccountCurrency SetAccountCurrencyRespMsgType = "set_account_currency"

var enumValues_SetAccountCurrencyRespMsgType = []interface{}{
	"set_account_currency",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetAccountCurrencyRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SetAccountCurrencyRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SetAccountCurrencyRespMsgType, v)
	}
	*j = SetAccountCurrencyRespMsgType(v)
	return nil
}

type SetAccountCurrencyRespSetAccountCurrency int

var enumValues_SetAccountCurrencyRespSetAccountCurrency = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetAccountCurrencyRespSetAccountCurrency) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SetAccountCurrencyRespSetAccountCurrency {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SetAccountCurrencyRespSetAccountCurrency, v)
	}
	*j = SetAccountCurrencyRespSetAccountCurrency(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetAccountCurrencyResp) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["echo_req"]; raw != nil && !ok {
		return fmt.Errorf("field echo_req in SetAccountCurrencyResp: required")
	}
	if _, ok := raw["msg_type"]; raw != nil && !ok {
		return fmt.Errorf("field msg_type in SetAccountCurrencyResp: required")
	}
	type Plain SetAccountCurrencyResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SetAccountCurrencyResp(plain)
	return nil
}
