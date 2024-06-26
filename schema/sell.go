// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"

// Sell a Contract as identified from a previous `portfolio` call.
type Sell struct {
	// [Optional] The login id of the user. Mandatory when multiple tokens were
	// provided during authorize.
	Loginid *string `json:"loginid,omitempty"`

	// [Optional] Used to pass data through the websocket, which may be retrieved via
	// the `echo_req` output field.
	Passthrough SellPassthrough `json:"passthrough,omitempty"`

	// Minimum price at which to sell the contract, or `0` for 'sell at market'.
	Price float64 `json:"price"`

	// [Optional] Used to map request to response.
	ReqId *int `json:"req_id,omitempty"`

	// Pass contract_id received from the `portfolio` call.
	Sell int `json:"sell"`
}

// [Optional] Used to pass data through the websocket, which may be retrieved via
// the `echo_req` output field.
type SellPassthrough map[string]interface{}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Sell) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["price"]; raw != nil && !ok {
		return fmt.Errorf("field price in Sell: required")
	}
	if _, ok := raw["sell"]; raw != nil && !ok {
		return fmt.Errorf("field sell in Sell: required")
	}
	type Plain Sell
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Sell(plain)
	return nil
}
