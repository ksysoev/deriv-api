// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "fmt"
import "reflect"
import "encoding/json"

// [Optional] Used to pass data through the websocket, which may be retrieved via
// the `echo_req` output field. Maximum size is 3500 bytes.
type TradingPlatformPasswordResetPassthrough map[string]interface{}

type TradingPlatformPasswordResetPlatform string

var enumValues_TradingPlatformPasswordResetPlatform = []interface{}{
	"dxtrade",
	"mt5",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TradingPlatformPasswordResetPlatform) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TradingPlatformPasswordResetPlatform {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TradingPlatformPasswordResetPlatform, v)
	}
	*j = TradingPlatformPasswordResetPlatform(v)
	return nil
}

const TradingPlatformPasswordResetPlatformDxtrade TradingPlatformPasswordResetPlatform = "dxtrade"
const TradingPlatformPasswordResetPlatformMt5 TradingPlatformPasswordResetPlatform = "mt5"

type TradingPlatformPasswordResetTradingPlatformPasswordReset int

var enumValues_TradingPlatformPasswordResetTradingPlatformPasswordReset = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TradingPlatformPasswordResetTradingPlatformPasswordReset) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TradingPlatformPasswordResetTradingPlatformPasswordReset {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TradingPlatformPasswordResetTradingPlatformPasswordReset, v)
	}
	*j = TradingPlatformPasswordResetTradingPlatformPasswordReset(v)
	return nil
}

// Reset the password of a Trading Platform Account
type TradingPlatformPasswordReset struct {
	// New password of the account. For validation (Accepts any printable ASCII
	// character. Must be within 8-25 characters, and include numbers, lowercase and
	// uppercase letters. Must not be the same as the user's email address).
	NewPassword string `json:"new_password"`

	// [Optional] Used to pass data through the websocket, which may be retrieved via
	// the `echo_req` output field. Maximum size is 3500 bytes.
	Passthrough TradingPlatformPasswordResetPassthrough `json:"passthrough,omitempty"`

	// Name of trading platform.
	Platform TradingPlatformPasswordResetPlatform `json:"platform"`

	// [Optional] Used to map request to response.
	ReqId *int `json:"req_id,omitempty"`

	// Must be `1`
	TradingPlatformPasswordReset TradingPlatformPasswordResetTradingPlatformPasswordReset `json:"trading_platform_password_reset"`

	// Email verification code (received from a `verify_email` call, which must be
	// done first)
	VerificationCode string `json:"verification_code"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TradingPlatformPasswordReset) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["new_password"]; !ok || v == nil {
		return fmt.Errorf("field new_password: required")
	}
	if v, ok := raw["platform"]; !ok || v == nil {
		return fmt.Errorf("field platform: required")
	}
	if v, ok := raw["trading_platform_password_reset"]; !ok || v == nil {
		return fmt.Errorf("field trading_platform_password_reset: required")
	}
	if v, ok := raw["verification_code"]; !ok || v == nil {
		return fmt.Errorf("field verification_code: required")
	}
	type Plain TradingPlatformPasswordReset
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = TradingPlatformPasswordReset(plain)
	return nil
}