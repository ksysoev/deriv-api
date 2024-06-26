// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// Return details of user account balance
type BalanceResp struct {
	// Current balance of one or more accounts.
	Balance *BalanceRespBalance `json:"balance,omitempty"`

	// Echo of the request made.
	EchoReq BalanceRespEchoReq `json:"echo_req"`

	// Action name of the request made.
	MsgType BalanceRespMsgType `json:"msg_type"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`

	// For subscription requests only.
	Subscription *BalanceRespSubscription `json:"subscription,omitempty"`
}

// Current balance of one or more accounts.
type BalanceRespBalance struct {
	// List of active accounts.
	Accounts BalanceRespBalanceAccounts `json:"accounts,omitempty"`

	// Balance of current account.
	Balance float64 `json:"balance"`

	// Currency of current account.
	Currency string `json:"currency"`

	// A per-connection unique identifier. Can be passed to the `forget` API call to
	// unsubscribe.
	Id *string `json:"id,omitempty"`

	// Client loginid.
	Loginid string `json:"loginid"`

	// Summary totals of accounts by type.
	Total *BalanceRespBalanceTotal `json:"total,omitempty"`
}

// List of active accounts.
type BalanceRespBalanceAccounts map[string]interface{}

// Summary totals of accounts by type.
type BalanceRespBalanceTotal struct {
	// Total balance of all real money Deriv accounts.
	Deriv *BalanceRespBalanceTotalDeriv `json:"deriv,omitempty"`

	// Total balance of all demo Deriv accounts.
	DerivDemo *BalanceRespBalanceTotalDerivDemo `json:"deriv_demo,omitempty"`

	// Total balance of all MT5 real money accounts.
	Mt5 *BalanceRespBalanceTotalMt5 `json:"mt5,omitempty"`

	// Total balance of all MT5 demo accounts.
	Mt5Demo *BalanceRespBalanceTotalMt5Demo `json:"mt5_demo,omitempty"`
}

// Total balance of all real money Deriv accounts.
type BalanceRespBalanceTotalDeriv struct {
	// Total of balances.
	Amount float64 `json:"amount"`

	// Currency of total.
	Currency string `json:"currency"`
}

// Total balance of all demo Deriv accounts.
type BalanceRespBalanceTotalDerivDemo struct {
	// Total of balances.
	Amount float64 `json:"amount"`

	// Currency of total.
	Currency string `json:"currency"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *BalanceRespBalanceTotalDerivDemo) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["amount"]; raw != nil && !ok {
		return fmt.Errorf("field amount in BalanceRespBalanceTotalDerivDemo: required")
	}
	if _, ok := raw["currency"]; raw != nil && !ok {
		return fmt.Errorf("field currency in BalanceRespBalanceTotalDerivDemo: required")
	}
	type Plain BalanceRespBalanceTotalDerivDemo
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = BalanceRespBalanceTotalDerivDemo(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *BalanceRespBalanceTotalDeriv) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["amount"]; raw != nil && !ok {
		return fmt.Errorf("field amount in BalanceRespBalanceTotalDeriv: required")
	}
	if _, ok := raw["currency"]; raw != nil && !ok {
		return fmt.Errorf("field currency in BalanceRespBalanceTotalDeriv: required")
	}
	type Plain BalanceRespBalanceTotalDeriv
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = BalanceRespBalanceTotalDeriv(plain)
	return nil
}

// Total balance of all MT5 real money accounts.
type BalanceRespBalanceTotalMt5 struct {
	// Total balance of all MT5 accounts
	Amount float64 `json:"amount"`

	// Currency of total.
	Currency string `json:"currency"`
}

// Total balance of all MT5 demo accounts.
type BalanceRespBalanceTotalMt5Demo struct {
	// Total of balances.
	Amount float64 `json:"amount"`

	// Currency of total.
	Currency string `json:"currency"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *BalanceRespBalanceTotalMt5Demo) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["amount"]; raw != nil && !ok {
		return fmt.Errorf("field amount in BalanceRespBalanceTotalMt5Demo: required")
	}
	if _, ok := raw["currency"]; raw != nil && !ok {
		return fmt.Errorf("field currency in BalanceRespBalanceTotalMt5Demo: required")
	}
	type Plain BalanceRespBalanceTotalMt5Demo
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = BalanceRespBalanceTotalMt5Demo(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *BalanceRespBalanceTotalMt5) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["amount"]; raw != nil && !ok {
		return fmt.Errorf("field amount in BalanceRespBalanceTotalMt5: required")
	}
	if _, ok := raw["currency"]; raw != nil && !ok {
		return fmt.Errorf("field currency in BalanceRespBalanceTotalMt5: required")
	}
	type Plain BalanceRespBalanceTotalMt5
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = BalanceRespBalanceTotalMt5(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *BalanceRespBalance) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["balance"]; raw != nil && !ok {
		return fmt.Errorf("field balance in BalanceRespBalance: required")
	}
	if _, ok := raw["currency"]; raw != nil && !ok {
		return fmt.Errorf("field currency in BalanceRespBalance: required")
	}
	if _, ok := raw["loginid"]; raw != nil && !ok {
		return fmt.Errorf("field loginid in BalanceRespBalance: required")
	}
	type Plain BalanceRespBalance
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = BalanceRespBalance(plain)
	return nil
}

// Echo of the request made.
type BalanceRespEchoReq map[string]interface{}

type BalanceRespMsgType string

const BalanceRespMsgTypeBalance BalanceRespMsgType = "balance"

var enumValues_BalanceRespMsgType = []interface{}{
	"balance",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *BalanceRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_BalanceRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_BalanceRespMsgType, v)
	}
	*j = BalanceRespMsgType(v)
	return nil
}

// For subscription requests only.
type BalanceRespSubscription struct {
	// A per-connection unique identifier. Can be passed to the `forget` API call to
	// unsubscribe.
	Id string `json:"id"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *BalanceRespSubscription) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["id"]; raw != nil && !ok {
		return fmt.Errorf("field id in BalanceRespSubscription: required")
	}
	type Plain BalanceRespSubscription
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = BalanceRespSubscription(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *BalanceResp) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["echo_req"]; raw != nil && !ok {
		return fmt.Errorf("field echo_req in BalanceResp: required")
	}
	if _, ok := raw["msg_type"]; raw != nil && !ok {
		return fmt.Errorf("field msg_type in BalanceResp: required")
	}
	type Plain BalanceResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = BalanceResp(plain)
	return nil
}
