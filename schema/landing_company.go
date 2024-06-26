// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"

// The company has a number of licensed subsidiaries in various jurisdictions,
// which are called Landing Companies. This call will return the appropriate
// Landing Company for clients of a given country. The landing company may differ
// for derived contracts (Synthetic Indices) and Financial contracts (Forex, Stock
// Indices, Commodities).
type LandingCompany struct {
	// Client's 2-letter country code (obtained from `residence_list` call).
	LandingCompany string `json:"landing_company"`

	// [Optional] Used to pass data through the websocket, which may be retrieved via
	// the `echo_req` output field.
	Passthrough LandingCompanyPassthrough `json:"passthrough,omitempty"`

	// [Optional] Used to map request to response.
	ReqId *int `json:"req_id,omitempty"`
}

// [Optional] Used to pass data through the websocket, which may be retrieved via
// the `echo_req` output field.
type LandingCompanyPassthrough map[string]interface{}

// UnmarshalJSON implements json.Unmarshaler.
func (j *LandingCompany) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["landing_company"]; raw != nil && !ok {
		return fmt.Errorf("field landing_company in LandingCompany: required")
	}
	type Plain LandingCompany
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = LandingCompany(plain)
	return nil
}
