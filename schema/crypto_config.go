// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// The request for cryptocurrencies configuration.
type CryptoConfig struct {
	// Must be `1`
	CryptoConfig CryptoConfigCryptoConfig `json:"crypto_config"`

	// [Optional] Cryptocurrency code. Sending request with currency_code provides
	// crypto config for the sent cryptocurrency code only.
	CurrencyCode *string `json:"currency_code,omitempty"`

	// [Optional] The login id of the user. Mandatory when multiple tokens were
	// provided during authorize.
	Loginid *string `json:"loginid,omitempty"`

	// [Optional] Used to pass data through the websocket, which may be retrieved via
	// the `echo_req` output field.
	Passthrough CryptoConfigPassthrough `json:"passthrough,omitempty"`

	// [Optional] Used to map request to response.
	ReqId *int `json:"req_id,omitempty"`
}

type CryptoConfigCryptoConfig int

var enumValues_CryptoConfigCryptoConfig = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CryptoConfigCryptoConfig) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CryptoConfigCryptoConfig {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CryptoConfigCryptoConfig, v)
	}
	*j = CryptoConfigCryptoConfig(v)
	return nil
}

// [Optional] Used to pass data through the websocket, which may be retrieved via
// the `echo_req` output field.
type CryptoConfigPassthrough map[string]interface{}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CryptoConfig) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["crypto_config"]; raw != nil && !ok {
		return fmt.Errorf("field crypto_config in CryptoConfig: required")
	}
	type Plain CryptoConfig
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CryptoConfig(plain)
	return nil
}
