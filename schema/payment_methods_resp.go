// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// List of available payment methods for a given country.
type PaymentMethodsResp struct {
	// Echo of the request made.
	EchoReq PaymentMethodsRespEchoReq `json:"echo_req"`

	// Action name of the request made.
	MsgType PaymentMethodsRespMsgType `json:"msg_type"`

	// Available payment methods for a given country. Note: if a user is logged in,
	// the residence country will be considered.
	PaymentMethods []PaymentMethodsRespPaymentMethodsElem `json:"payment_methods,omitempty"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`
}

// Echo of the request made.
type PaymentMethodsRespEchoReq map[string]interface{}

type PaymentMethodsRespMsgType string

const PaymentMethodsRespMsgTypePaymentMethods PaymentMethodsRespMsgType = "payment_methods"

var enumValues_PaymentMethodsRespMsgType = []interface{}{
	"payment_methods",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PaymentMethodsRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_PaymentMethodsRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_PaymentMethodsRespMsgType, v)
	}
	*j = PaymentMethodsRespMsgType(v)
	return nil
}

// A payment method suported for the given country
type PaymentMethodsRespPaymentMethodsElem struct {
	// The min and max values for deposits.
	DepositLimits PaymentMethodsRespPaymentMethodsElemDepositLimits `json:"deposit_limits"`

	// How much time it takes for a deposit to be processed.
	DepositTime string `json:"deposit_time"`

	// Short description explaining the payment method.
	Description string `json:"description"`

	// Common name for the payment method.
	DisplayName string `json:"display_name"`

	// Unique identifier for the payment method.
	Id string `json:"id"`

	// Payment processor for this payment method.
	PaymentProcessor string `json:"payment_processor"`

	// A list of predefined amounts for withdraw or deposit.
	PredefinedAmounts []int `json:"predefined_amounts"`

	// Sign up link for this payment method.
	SignupLink string `json:"signup_link"`

	// Currencies supported for this payment method.
	SupportedCurrencies []string `json:"supported_currencies"`

	// Type of Payment Method.
	Type string `json:"type"`

	// A printable description for type of payment method.
	TypeDisplayName string `json:"type_display_name"`

	// Withdrawal limits per currency.
	WithdrawLimits PaymentMethodsRespPaymentMethodsElemWithdrawLimits `json:"withdraw_limits"`

	// How much time takes a withdrawal to be processed.
	WithdrawalTime string `json:"withdrawal_time"`
}

// The min and max values for deposits.
type PaymentMethodsRespPaymentMethodsElemDepositLimits map[string]interface{}

// Withdrawal limits per currency.
type PaymentMethodsRespPaymentMethodsElemWithdrawLimits map[string]interface{}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PaymentMethodsRespPaymentMethodsElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["deposit_limits"]; raw != nil && !ok {
		return fmt.Errorf("field deposit_limits in PaymentMethodsRespPaymentMethodsElem: required")
	}
	if _, ok := raw["deposit_time"]; raw != nil && !ok {
		return fmt.Errorf("field deposit_time in PaymentMethodsRespPaymentMethodsElem: required")
	}
	if _, ok := raw["description"]; raw != nil && !ok {
		return fmt.Errorf("field description in PaymentMethodsRespPaymentMethodsElem: required")
	}
	if _, ok := raw["display_name"]; raw != nil && !ok {
		return fmt.Errorf("field display_name in PaymentMethodsRespPaymentMethodsElem: required")
	}
	if _, ok := raw["id"]; raw != nil && !ok {
		return fmt.Errorf("field id in PaymentMethodsRespPaymentMethodsElem: required")
	}
	if _, ok := raw["payment_processor"]; raw != nil && !ok {
		return fmt.Errorf("field payment_processor in PaymentMethodsRespPaymentMethodsElem: required")
	}
	if _, ok := raw["predefined_amounts"]; raw != nil && !ok {
		return fmt.Errorf("field predefined_amounts in PaymentMethodsRespPaymentMethodsElem: required")
	}
	if _, ok := raw["signup_link"]; raw != nil && !ok {
		return fmt.Errorf("field signup_link in PaymentMethodsRespPaymentMethodsElem: required")
	}
	if _, ok := raw["supported_currencies"]; raw != nil && !ok {
		return fmt.Errorf("field supported_currencies in PaymentMethodsRespPaymentMethodsElem: required")
	}
	if _, ok := raw["type"]; raw != nil && !ok {
		return fmt.Errorf("field type in PaymentMethodsRespPaymentMethodsElem: required")
	}
	if _, ok := raw["type_display_name"]; raw != nil && !ok {
		return fmt.Errorf("field type_display_name in PaymentMethodsRespPaymentMethodsElem: required")
	}
	if _, ok := raw["withdraw_limits"]; raw != nil && !ok {
		return fmt.Errorf("field withdraw_limits in PaymentMethodsRespPaymentMethodsElem: required")
	}
	if _, ok := raw["withdrawal_time"]; raw != nil && !ok {
		return fmt.Errorf("field withdrawal_time in PaymentMethodsRespPaymentMethodsElem: required")
	}
	type Plain PaymentMethodsRespPaymentMethodsElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = PaymentMethodsRespPaymentMethodsElem(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PaymentMethodsResp) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["echo_req"]; raw != nil && !ok {
		return fmt.Errorf("field echo_req in PaymentMethodsResp: required")
	}
	if _, ok := raw["msg_type"]; raw != nil && !ok {
		return fmt.Errorf("field msg_type in PaymentMethodsResp: required")
	}
	type Plain PaymentMethodsResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = PaymentMethodsResp(plain)
	return nil
}
