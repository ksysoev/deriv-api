// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// UnmarshalJSON implements json.Unmarshaler.
func (j *TransferBetweenAccountsRespAccountsElemDemoAccount) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TransferBetweenAccountsRespAccountsElemDemoAccount {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TransferBetweenAccountsRespAccountsElemDemoAccount, v)
	}
	*j = TransferBetweenAccountsRespAccountsElemDemoAccount(v)
	return nil
}

const TransferBetweenAccountsRespAccountsElemTransfersNone TransferBetweenAccountsRespAccountsElemTransfers = "none"

// UnmarshalJSON implements json.Unmarshaler.
func (j *TransferBetweenAccountsRespAccountsElemAccountCategory) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TransferBetweenAccountsRespAccountsElemAccountCategory {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TransferBetweenAccountsRespAccountsElemAccountCategory, v)
	}
	*j = TransferBetweenAccountsRespAccountsElemAccountCategory(v)
	return nil
}

const TransferBetweenAccountsRespAccountsElemAccountCategoryTrading TransferBetweenAccountsRespAccountsElemAccountCategory = "trading"
const TransferBetweenAccountsRespAccountsElemAccountCategoryWallet TransferBetweenAccountsRespAccountsElemAccountCategory = "wallet"

type TransferBetweenAccountsRespAccountsElemAccountType string

// UnmarshalJSON implements json.Unmarshaler.
func (j *TransferBetweenAccountsResp) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["echo_req"]; !ok || v == nil {
		return fmt.Errorf("field echo_req in TransferBetweenAccountsResp: required")
	}
	if v, ok := raw["msg_type"]; !ok || v == nil {
		return fmt.Errorf("field msg_type in TransferBetweenAccountsResp: required")
	}
	type Plain TransferBetweenAccountsResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = TransferBetweenAccountsResp(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TransferBetweenAccountsRespAccountsElemAccountType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TransferBetweenAccountsRespAccountsElemAccountType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TransferBetweenAccountsRespAccountsElemAccountType, v)
	}
	*j = TransferBetweenAccountsRespAccountsElemAccountType(v)
	return nil
}

const TransferBetweenAccountsRespAccountsElemAccountTypeBinary TransferBetweenAccountsRespAccountsElemAccountType = "binary"
const TransferBetweenAccountsRespAccountsElemAccountTypeCrypto TransferBetweenAccountsRespAccountsElemAccountType = "crypto"
const TransferBetweenAccountsRespAccountsElemAccountTypeCtrader TransferBetweenAccountsRespAccountsElemAccountType = "ctrader"
const TransferBetweenAccountsRespAccountsElemAccountTypeDoughflow TransferBetweenAccountsRespAccountsElemAccountType = "doughflow"
const TransferBetweenAccountsRespAccountsElemAccountTypeDxtrade TransferBetweenAccountsRespAccountsElemAccountType = "dxtrade"
const TransferBetweenAccountsRespAccountsElemAccountTypeDerivez TransferBetweenAccountsRespAccountsElemAccountType = "derivez"
const TransferBetweenAccountsRespAccountsElemAccountTypeMt5 TransferBetweenAccountsRespAccountsElemAccountType = "mt5"
const TransferBetweenAccountsRespAccountsElemAccountTypeP2P TransferBetweenAccountsRespAccountsElemAccountType = "p2p"
const TransferBetweenAccountsRespAccountsElemAccountTypePaymentagent TransferBetweenAccountsRespAccountsElemAccountType = "paymentagent"
const TransferBetweenAccountsRespAccountsElemAccountTypePaymentagentClient TransferBetweenAccountsRespAccountsElemAccountType = "paymentagent_client"
const TransferBetweenAccountsRespAccountsElemAccountTypeStandard TransferBetweenAccountsRespAccountsElemAccountType = "standard"
const TransferBetweenAccountsRespAccountsElemAccountTypeVirtual TransferBetweenAccountsRespAccountsElemAccountType = "virtual"

type TransferBetweenAccountsRespAccountsElemDemoAccount int

type TransferBetweenAccountsRespAccountsElemAccountCategory string

// The result of transfer order.
type TransferBetweenAccountsResp struct {
	// The available accounts to transfer, or the accounts affected by a successful
	// transfer.
	Accounts []TransferBetweenAccountsRespAccountsElem `json:"accounts,omitempty"`

	// The account to client full name
	ClientToFullName *string `json:"client_to_full_name,omitempty"`

	// The account to client loginid
	ClientToLoginid *string `json:"client_to_loginid,omitempty"`

	// Echo of the request made.
	EchoReq TransferBetweenAccountsRespEchoReq `json:"echo_req"`

	// Action name of the request made.
	MsgType TransferBetweenAccountsRespMsgType `json:"msg_type"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`

	// Reference ID of transfer performed
	TransactionId *int `json:"transaction_id,omitempty"`

	// If set to 1, transfer succeeded.
	TransferBetweenAccounts *TransferBetweenAccountsRespTransferBetweenAccounts `json:"transfer_between_accounts,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TransferBetweenAccountsRespTransferBetweenAccounts) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TransferBetweenAccountsRespTransferBetweenAccounts {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TransferBetweenAccountsRespTransferBetweenAccounts, v)
	}
	*j = TransferBetweenAccountsRespTransferBetweenAccounts(v)
	return nil
}

const TransferBetweenAccountsRespAccountsElemMarketTypeFinancial TransferBetweenAccountsRespAccountsElemMarketType = "financial"

// UnmarshalJSON implements json.Unmarshaler.
func (j *TransferBetweenAccountsRespAccountsElemMarketType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TransferBetweenAccountsRespAccountsElemMarketType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TransferBetweenAccountsRespAccountsElemMarketType, v)
	}
	*j = TransferBetweenAccountsRespAccountsElemMarketType(v)
	return nil
}

const TransferBetweenAccountsRespAccountsElemMarketTypeAll TransferBetweenAccountsRespAccountsElemMarketType = "all"

type TransferBetweenAccountsRespAccountsElemMarketType string

const TransferBetweenAccountsRespAccountsElemMarketTypeSynthetic TransferBetweenAccountsRespAccountsElemMarketType = "synthetic"

type TransferBetweenAccountsRespAccountsElemTransfers string

// UnmarshalJSON implements json.Unmarshaler.
func (j *TransferBetweenAccountsRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TransferBetweenAccountsRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TransferBetweenAccountsRespMsgType, v)
	}
	*j = TransferBetweenAccountsRespMsgType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TransferBetweenAccountsRespAccountsElemTransfers) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TransferBetweenAccountsRespAccountsElemTransfers {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TransferBetweenAccountsRespAccountsElemTransfers, v)
	}
	*j = TransferBetweenAccountsRespAccountsElemTransfers(v)
	return nil
}

const TransferBetweenAccountsRespAccountsElemTransfersAll TransferBetweenAccountsRespAccountsElemTransfers = "all"
const TransferBetweenAccountsRespAccountsElemTransfersDeposit TransferBetweenAccountsRespAccountsElemTransfers = "deposit"

type TransferBetweenAccountsRespAccountsElem struct {
	// Category of the account.
	AccountCategory *TransferBetweenAccountsRespAccountsElemAccountCategory `json:"account_category,omitempty"`

	// Type of the account.
	AccountType *TransferBetweenAccountsRespAccountsElemAccountType `json:"account_type,omitempty"`

	// Account balance.
	Balance *string `json:"balance,omitempty"`

	// Default account currency.
	Currency *string `json:"currency,omitempty"`

	// 0 for real accounts; 1 for virtual/demo accounts.
	DemoAccount *TransferBetweenAccountsRespAccountsElemDemoAccount `json:"demo_account,omitempty"`

	// The group of derivez account.
	DerivezGroup *string `json:"derivez_group,omitempty"`

	// Account identifier used for system transfers.
	Loginid *string `json:"loginid,omitempty"`

	// Market type of account.
	MarketType *TransferBetweenAccountsRespAccountsElemMarketType `json:"market_type,omitempty"`

	// The group of mt5 account.
	Mt5Group *string `json:"mt5_group,omitempty"`

	// The status of account.
	Status *string `json:"status,omitempty"`

	// Type of transfers allowed between the account and the currently authorized
	// account.
	Transfers *TransferBetweenAccountsRespAccountsElemTransfers `json:"transfers,omitempty"`
}

const TransferBetweenAccountsRespAccountsElemTransfersWithdrawal TransferBetweenAccountsRespAccountsElemTransfers = "withdrawal"

// Echo of the request made.
type TransferBetweenAccountsRespEchoReq map[string]interface{}

type TransferBetweenAccountsRespMsgType string

const TransferBetweenAccountsRespMsgTypeTransferBetweenAccounts TransferBetweenAccountsRespMsgType = "transfer_between_accounts"

type TransferBetweenAccountsRespTransferBetweenAccounts int

var enumValues_TransferBetweenAccountsRespAccountsElemAccountCategory = []interface{}{
	"trading",
	"wallet",
}
var enumValues_TransferBetweenAccountsRespAccountsElemAccountType = []interface{}{
	"binary",
	"crypto",
	"ctrader",
	"doughflow",
	"dxtrade",
	"derivez",
	"mt5",
	"p2p",
	"paymentagent",
	"paymentagent_client",
	"standard",
	"virtual",
}
var enumValues_TransferBetweenAccountsRespAccountsElemDemoAccount = []interface{}{
	0,
	1,
}
var enumValues_TransferBetweenAccountsRespAccountsElemMarketType = []interface{}{
	"all",
	"financial",
	"synthetic",
}
var enumValues_TransferBetweenAccountsRespAccountsElemTransfers = []interface{}{
	"all",
	"deposit",
	"none",
	"withdrawal",
}
var enumValues_TransferBetweenAccountsRespMsgType = []interface{}{
	"transfer_between_accounts",
}
var enumValues_TransferBetweenAccountsRespTransferBetweenAccounts = []interface{}{
	0,
	1,
}
