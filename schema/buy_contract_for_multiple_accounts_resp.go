// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// A message with transaction results is received
type BuyContractForMultipleAccountsResp struct {
	// Receipt confirmation for the purchase
	BuyContractForMultipleAccounts *BuyContractForMultipleAccountsRespBuyContractForMultipleAccounts `json:"buy_contract_for_multiple_accounts,omitempty"`

	// Echo of the request made.
	EchoReq BuyContractForMultipleAccountsRespEchoReq `json:"echo_req"`

	// Action name of the request made.
	MsgType BuyContractForMultipleAccountsRespMsgType `json:"msg_type"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`
}

// Receipt confirmation for the purchase
type BuyContractForMultipleAccountsRespBuyContractForMultipleAccounts struct {
	// List of results containing transactions and/or errors for the bought contracts.
	Result []interface{} `json:"result"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *BuyContractForMultipleAccountsRespBuyContractForMultipleAccounts) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["result"]; raw != nil && !ok {
		return fmt.Errorf("field result in BuyContractForMultipleAccountsRespBuyContractForMultipleAccounts: required")
	}
	type Plain BuyContractForMultipleAccountsRespBuyContractForMultipleAccounts
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = BuyContractForMultipleAccountsRespBuyContractForMultipleAccounts(plain)
	return nil
}

// Echo of the request made.
type BuyContractForMultipleAccountsRespEchoReq map[string]interface{}

type BuyContractForMultipleAccountsRespMsgType string

const BuyContractForMultipleAccountsRespMsgTypeBuyContractForMultipleAccounts BuyContractForMultipleAccountsRespMsgType = "buy_contract_for_multiple_accounts"

var enumValues_BuyContractForMultipleAccountsRespMsgType = []interface{}{
	"buy_contract_for_multiple_accounts",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *BuyContractForMultipleAccountsRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_BuyContractForMultipleAccountsRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_BuyContractForMultipleAccountsRespMsgType, v)
	}
	*j = BuyContractForMultipleAccountsRespMsgType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *BuyContractForMultipleAccountsResp) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["echo_req"]; raw != nil && !ok {
		return fmt.Errorf("field echo_req in BuyContractForMultipleAccountsResp: required")
	}
	if _, ok := raw["msg_type"]; raw != nil && !ok {
		return fmt.Errorf("field msg_type in BuyContractForMultipleAccountsResp: required")
	}
	type Plain BuyContractForMultipleAccountsResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = BuyContractForMultipleAccountsResp(plain)
	return nil
}
