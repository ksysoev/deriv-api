// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "fmt"
import "encoding/json"

// Sell a Contract as identified from a previous `portfolio` call.
type Sell struct {
	// [Optional] Used to pass data through the websocket, which may be retrieved via
	// the `echo_req` output field. Maximum size is 3500 bytes.
	Passthrough SellPassthrough `json:"passthrough,omitempty"`

	// Minimum price at which to sell the contract, or `0` for 'sell at market'.
	Price float64 `json:"price"`

	// [Optional] Used to map request to response.
	ReqId *int `json:"req_id,omitempty"`

	// Pass contract_id received from the `portfolio` call.
	Sell int `json:"sell"`
}

// [Optional] Used to pass data through the websocket, which may be retrieved via
// the `echo_req` output field. Maximum size is 3500 bytes.
type SellPassthrough map[string]interface{}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Sell) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["price"]; !ok || v == nil {
		return fmt.Errorf("field price: required")
	}
	if v, ok := raw["sell"]; !ok || v == nil {
		return fmt.Errorf("field sell: required")
	}
	type Plain Sell
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Sell(plain)
	return nil
}