// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "fmt"
import "reflect"
import "encoding/json"

type ExchangeRatesExchangeRates int

var enumValues_ExchangeRatesExchangeRates = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ExchangeRatesExchangeRates) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ExchangeRatesExchangeRates {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ExchangeRatesExchangeRates, v)
	}
	*j = ExchangeRatesExchangeRates(v)
	return nil
}

// [Optional] Used to pass data through the websocket, which may be retrieved via
// the `echo_req` output field. Maximum size is 3500 bytes.
type ExchangeRatesPassthrough map[string]interface{}

type ExchangeRatesSubscribe int

var enumValues_ExchangeRatesSubscribe = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ExchangeRatesSubscribe) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ExchangeRatesSubscribe {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ExchangeRatesSubscribe, v)
	}
	*j = ExchangeRatesSubscribe(v)
	return nil
}

// Retrieves the exchange rates from a base currency to all currencies supported by
// the system.
type ExchangeRates struct {
	// Base currency (can be obtained from `payout_currencies` call)
	BaseCurrency string `json:"base_currency"`

	// Must be `1`
	ExchangeRates ExchangeRatesExchangeRates `json:"exchange_rates"`

	// [Optional] Used to pass data through the websocket, which may be retrieved via
	// the `echo_req` output field. Maximum size is 3500 bytes.
	Passthrough ExchangeRatesPassthrough `json:"passthrough,omitempty"`

	// [Optional] Used to map request to response.
	ReqId *int `json:"req_id,omitempty"`

	// [Optional] 1 - to initiate a realtime stream of exchange rates relative to base
	// currency.
	Subscribe *ExchangeRatesSubscribe `json:"subscribe,omitempty"`

	// [Optional] Local currency
	TargetCurrency *string `json:"target_currency,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ExchangeRates) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["base_currency"]; !ok || v == nil {
		return fmt.Errorf("field base_currency: required")
	}
	if v, ok := raw["exchange_rates"]; !ok || v == nil {
		return fmt.Errorf("field exchange_rates: required")
	}
	type Plain ExchangeRates
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ExchangeRates(plain)
	return nil
}