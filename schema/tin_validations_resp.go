// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// A message with validations for Tax Identification Numbers (TIN)
type TinValidationsResp struct {
	// Echo of the request made.
	EchoReq TinValidationsRespEchoReq `json:"echo_req"`

	// Action name of the request made.
	MsgType TinValidationsRespMsgType `json:"msg_type"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`

	// Validations for Tax Identification Numbers (TIN)
	TinValidations *TinValidationsRespTinValidations `json:"tin_validations,omitempty"`
}

// Echo of the request made.
type TinValidationsRespEchoReq map[string]interface{}

type TinValidationsRespMsgType string

const TinValidationsRespMsgTypeTinValidations TinValidationsRespMsgType = "tin_validations"

var enumValues_TinValidationsRespMsgType = []interface{}{
	"tin_validations",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TinValidationsRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TinValidationsRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TinValidationsRespMsgType, v)
	}
	*j = TinValidationsRespMsgType(v)
	return nil
}

// Validations for Tax Identification Numbers (TIN)
type TinValidationsRespTinValidations struct {
	// Invalid regex patterns for tin validation
	InvalidPatterns []string `json:"invalid_patterns,omitempty"`

	// Whether the TIN is mandatory for the selected country
	IsTinMandatory *TinValidationsRespTinValidationsIsTinMandatory `json:"is_tin_mandatory,omitempty"`

	// List of employment statuses that bypass TIN requirements for the selected
	// country
	TinEmploymentStatusBypass []string `json:"tin_employment_status_bypass,omitempty"`

	// Country tax identifier formats.
	TinFormat []string `json:"tin_format,omitempty"`

	// Human-readable description of the TIN format for the selected country
	TinFormatDescription *string `json:"tin_format_description,omitempty"`
}

type TinValidationsRespTinValidationsIsTinMandatory int

var enumValues_TinValidationsRespTinValidationsIsTinMandatory = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TinValidationsRespTinValidationsIsTinMandatory) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TinValidationsRespTinValidationsIsTinMandatory {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TinValidationsRespTinValidationsIsTinMandatory, v)
	}
	*j = TinValidationsRespTinValidationsIsTinMandatory(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TinValidationsResp) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["echo_req"]; raw != nil && !ok {
		return fmt.Errorf("field echo_req in TinValidationsResp: required")
	}
	if _, ok := raw["msg_type"]; raw != nil && !ok {
		return fmt.Errorf("field msg_type in TinValidationsResp: required")
	}
	type Plain TinValidationsResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = TinValidationsResp(plain)
	return nil
}
