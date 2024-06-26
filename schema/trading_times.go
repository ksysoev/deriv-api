// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"

// Receive a list of market opening times for a given date.
type TradingTimes struct {
	// [Optional] Used to pass data through the websocket, which may be retrieved via
	// the `echo_req` output field.
	Passthrough TradingTimesPassthrough `json:"passthrough,omitempty"`

	// [Optional] Used to map request to response.
	ReqId *int `json:"req_id,omitempty"`

	// Date to receive market opening times for. (`yyyy-mm-dd` format. `today` can
	// also be specified).
	TradingTimes string `json:"trading_times"`
}

// [Optional] Used to pass data through the websocket, which may be retrieved via
// the `echo_req` output field.
type TradingTimesPassthrough map[string]interface{}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TradingTimes) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["trading_times"]; raw != nil && !ok {
		return fmt.Errorf("field trading_times in TradingTimes: required")
	}
	type Plain TradingTimes
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = TradingTimes(plain)
	return nil
}
