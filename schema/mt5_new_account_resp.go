// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "fmt"
import "reflect"
import "encoding/json"

// Create MT5 account Receive
type Mt5NewAccountResp struct {
	// Echo of the request made.
	EchoReq Mt5NewAccountRespEchoReq `json:"echo_req"`

	// Action name of the request made.
	MsgType Mt5NewAccountRespMsgType `json:"msg_type"`

	// New MT5 account details
	Mt5NewAccount *Mt5NewAccountRespMt5NewAccount `json:"mt5_new_account,omitempty"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`
}

// Echo of the request made.
type Mt5NewAccountRespEchoReq map[string]interface{}

type Mt5NewAccountRespMsgType string

const Mt5NewAccountRespMsgTypeMt5NewAccount Mt5NewAccountRespMsgType = "mt5_new_account"

// New MT5 account details
type Mt5NewAccountRespMt5NewAccount struct {
	// Account type.
	AccountType *Mt5NewAccountRespMt5NewAccountAccountType `json:"account_type,omitempty"`

	// Agent Details.
	Agent interface{} `json:"agent,omitempty"`

	// Account balance.
	Balance *float64 `json:"balance,omitempty"`

	// MT5 account currency (`USD` or `EUR`) that depends on the MT5 company
	// (`vanuatu`, `svg`, `malta`).
	Currency *string `json:"currency,omitempty"`

	// Account balance, formatted to appropriate decimal places.
	DisplayBalance *string `json:"display_balance,omitempty"`

	// Login ID of the user's new MT5 account. Login could have 2 types of prefixes:
	// MTD, MTR. MTD - for demo accounts and MTR for real money accounts.
	Login *string `json:"login,omitempty"`

	// With default value of conventional, unavailable for `financial_stp` sub account
	// type.
	Mt5AccountCategory *Mt5NewAccountRespMt5NewAccountMt5AccountCategory `json:"mt5_account_category,omitempty"`

	// Sub account type for classic MT5 account.
	Mt5AccountType *Mt5NewAccountRespMt5NewAccountMt5AccountType `json:"mt5_account_type,omitempty"`
}

type Mt5NewAccountRespMt5NewAccountAccountType string

const Mt5NewAccountRespMt5NewAccountAccountTypeAll Mt5NewAccountRespMt5NewAccountAccountType = "all"
const Mt5NewAccountRespMt5NewAccountAccountTypeDemo Mt5NewAccountRespMt5NewAccountAccountType = "demo"
const Mt5NewAccountRespMt5NewAccountAccountTypeFinancial Mt5NewAccountRespMt5NewAccountAccountType = "financial"
const Mt5NewAccountRespMt5NewAccountAccountTypeGaming Mt5NewAccountRespMt5NewAccountAccountType = "gaming"

type Mt5NewAccountRespMt5NewAccountMt5AccountCategory string

const Mt5NewAccountRespMt5NewAccountMt5AccountCategoryConventional Mt5NewAccountRespMt5NewAccountMt5AccountCategory = "conventional"
const Mt5NewAccountRespMt5NewAccountMt5AccountCategorySwapFree Mt5NewAccountRespMt5NewAccountMt5AccountCategory = "swap_free"

type Mt5NewAccountRespMt5NewAccountMt5AccountType string

const Mt5NewAccountRespMt5NewAccountMt5AccountTypeFinancial Mt5NewAccountRespMt5NewAccountMt5AccountType = "financial"
const Mt5NewAccountRespMt5NewAccountMt5AccountTypeFinancialStp Mt5NewAccountRespMt5NewAccountMt5AccountType = "financial_stp"
const Mt5NewAccountRespMt5NewAccountMt5AccountTypeStandard Mt5NewAccountRespMt5NewAccountMt5AccountType = "standard"

var enumValues_Mt5NewAccountRespMsgType = []interface{}{
	"mt5_new_account",
}
var enumValues_Mt5NewAccountRespMt5NewAccountAccountType = []interface{}{
	"demo",
	"gaming",
	"financial",
	"all",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Mt5NewAccountRespMt5NewAccountMt5AccountType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_Mt5NewAccountRespMt5NewAccountMt5AccountType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_Mt5NewAccountRespMt5NewAccountMt5AccountType, v)
	}
	*j = Mt5NewAccountRespMt5NewAccountMt5AccountType(v)
	return nil
}

var enumValues_Mt5NewAccountRespMt5NewAccountMt5AccountType = []interface{}{
	"financial",
	"financial_stp",
	"standard",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Mt5NewAccountRespMt5NewAccountMt5AccountCategory) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_Mt5NewAccountRespMt5NewAccountMt5AccountCategory {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_Mt5NewAccountRespMt5NewAccountMt5AccountCategory, v)
	}
	*j = Mt5NewAccountRespMt5NewAccountMt5AccountCategory(v)
	return nil
}

var enumValues_Mt5NewAccountRespMt5NewAccountMt5AccountCategory = []interface{}{
	"conventional",
	"swap_free",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Mt5NewAccountRespMt5NewAccountAccountType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_Mt5NewAccountRespMt5NewAccountAccountType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_Mt5NewAccountRespMt5NewAccountAccountType, v)
	}
	*j = Mt5NewAccountRespMt5NewAccountAccountType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Mt5NewAccountRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_Mt5NewAccountRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_Mt5NewAccountRespMsgType, v)
	}
	*j = Mt5NewAccountRespMsgType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Mt5NewAccountResp) UnmarshalJSON(b []byte) error {
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
	type Plain Mt5NewAccountResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Mt5NewAccountResp(plain)
	return nil
}