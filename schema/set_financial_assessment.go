// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// This call sets the financial assessment details based on the client's answers to
// analyze whether they possess the experience and knowledge to understand the
// risks involved with binary options trading.
type SetFinancialAssessment struct {
	// [Optional] The financial information of a client
	FinancialInformation SetFinancialAssessmentFinancialInformation `json:"financial_information,omitempty"`

	// [Optional] The version of the financial information
	FinancialInformationVersion *string `json:"financial_information_version,omitempty"`

	// [Optional] The login id of the user. Mandatory when multiple tokens were
	// provided during authorize.
	Loginid *string `json:"loginid,omitempty"`

	// [Optional] Used to pass data through the websocket, which may be retrieved via
	// the `echo_req` output field.
	Passthrough SetFinancialAssessmentPassthrough `json:"passthrough,omitempty"`

	// [Optional] Used to map request to response.
	ReqId *int `json:"req_id,omitempty"`

	// Must be `1`
	SetFinancialAssessment SetFinancialAssessmentSetFinancialAssessment `json:"set_financial_assessment"`

	// [Optional] The trading experience of a client
	TradingExperience SetFinancialAssessmentTradingExperience `json:"trading_experience,omitempty"`

	// [Optional] The trading experience of a `maltainvest` client
	TradingExperienceRegulated SetFinancialAssessmentTradingExperienceRegulated `json:"trading_experience_regulated,omitempty"`
}

// [Optional] The financial information of a client
type SetFinancialAssessmentFinancialInformation map[string]interface{}

// [Optional] Used to pass data through the websocket, which may be retrieved via
// the `echo_req` output field.
type SetFinancialAssessmentPassthrough map[string]interface{}

type SetFinancialAssessmentSetFinancialAssessment int

var enumValues_SetFinancialAssessmentSetFinancialAssessment = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetFinancialAssessmentSetFinancialAssessment) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SetFinancialAssessmentSetFinancialAssessment {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SetFinancialAssessmentSetFinancialAssessment, v)
	}
	*j = SetFinancialAssessmentSetFinancialAssessment(v)
	return nil
}

// [Optional] The trading experience of a client
type SetFinancialAssessmentTradingExperience map[string]interface{}

// [Optional] The trading experience of a `maltainvest` client
type SetFinancialAssessmentTradingExperienceRegulated map[string]interface{}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetFinancialAssessment) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["set_financial_assessment"]; raw != nil && !ok {
		return fmt.Errorf("field set_financial_assessment in SetFinancialAssessment: required")
	}
	type Plain SetFinancialAssessment
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SetFinancialAssessment(plain)
	return nil
}
