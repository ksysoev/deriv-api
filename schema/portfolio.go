// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "fmt"
import "reflect"
import "encoding/json"

// Receive information about my current portfolio of outstanding options
type Portfolio struct {
	// Return only contracts of the specified types
	ContractType []PortfolioContractTypeElem `json:"contract_type,omitempty"`

	// [Optional] Used to pass data through the websocket, which may be retrieved via
	// the `echo_req` output field. Maximum size is 3500 bytes.
	Passthrough PortfolioPassthrough `json:"passthrough,omitempty"`

	// Must be `1`
	Portfolio PortfolioPortfolio `json:"portfolio"`

	// [Optional] Used to map request to response.
	ReqId *int `json:"req_id,omitempty"`
}

type PortfolioContractTypeElem string

const PortfolioContractTypeElemASIAND PortfolioContractTypeElem = "ASIAND"
const PortfolioContractTypeElemASIANU PortfolioContractTypeElem = "ASIANU"
const PortfolioContractTypeElemCALL PortfolioContractTypeElem = "CALL"
const PortfolioContractTypeElemCALLE PortfolioContractTypeElem = "CALLE"
const PortfolioContractTypeElemCALLSPREAD PortfolioContractTypeElem = "CALLSPREAD"
const PortfolioContractTypeElemDIGITDIFF PortfolioContractTypeElem = "DIGITDIFF"
const PortfolioContractTypeElemDIGITEVEN PortfolioContractTypeElem = "DIGITEVEN"
const PortfolioContractTypeElemDIGITMATCH PortfolioContractTypeElem = "DIGITMATCH"
const PortfolioContractTypeElemDIGITODD PortfolioContractTypeElem = "DIGITODD"
const PortfolioContractTypeElemDIGITOVER PortfolioContractTypeElem = "DIGITOVER"
const PortfolioContractTypeElemDIGITUNDER PortfolioContractTypeElem = "DIGITUNDER"
const PortfolioContractTypeElemEXPIRYMISS PortfolioContractTypeElem = "EXPIRYMISS"
const PortfolioContractTypeElemEXPIRYMISSE PortfolioContractTypeElem = "EXPIRYMISSE"
const PortfolioContractTypeElemEXPIRYRANGE PortfolioContractTypeElem = "EXPIRYRANGE"
const PortfolioContractTypeElemEXPIRYRANGEE PortfolioContractTypeElem = "EXPIRYRANGEE"
const PortfolioContractTypeElemLBFLOATCALL PortfolioContractTypeElem = "LBFLOATCALL"
const PortfolioContractTypeElemLBFLOATPUT PortfolioContractTypeElem = "LBFLOATPUT"
const PortfolioContractTypeElemLBHIGHLOW PortfolioContractTypeElem = "LBHIGHLOW"
const PortfolioContractTypeElemMULTDOWN PortfolioContractTypeElem = "MULTDOWN"
const PortfolioContractTypeElemMULTUP PortfolioContractTypeElem = "MULTUP"
const PortfolioContractTypeElemNOTOUCH PortfolioContractTypeElem = "NOTOUCH"
const PortfolioContractTypeElemONETOUCH PortfolioContractTypeElem = "ONETOUCH"
const PortfolioContractTypeElemPUT PortfolioContractTypeElem = "PUT"
const PortfolioContractTypeElemPUTE PortfolioContractTypeElem = "PUTE"
const PortfolioContractTypeElemPUTSPREAD PortfolioContractTypeElem = "PUTSPREAD"
const PortfolioContractTypeElemRANGE PortfolioContractTypeElem = "RANGE"
const PortfolioContractTypeElemRESETCALL PortfolioContractTypeElem = "RESETCALL"
const PortfolioContractTypeElemRESETPUT PortfolioContractTypeElem = "RESETPUT"
const PortfolioContractTypeElemRUNHIGH PortfolioContractTypeElem = "RUNHIGH"
const PortfolioContractTypeElemRUNLOW PortfolioContractTypeElem = "RUNLOW"
const PortfolioContractTypeElemTICKHIGH PortfolioContractTypeElem = "TICKHIGH"
const PortfolioContractTypeElemTICKLOW PortfolioContractTypeElem = "TICKLOW"

// UnmarshalJSON implements json.Unmarshaler.
func (j *PortfolioContractTypeElem) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_PortfolioContractTypeElem {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_PortfolioContractTypeElem, v)
	}
	*j = PortfolioContractTypeElem(v)
	return nil
}

const PortfolioContractTypeElemTURBOSLONG PortfolioContractTypeElem = "TURBOSLONG"
const PortfolioContractTypeElemTURBOSSHORT PortfolioContractTypeElem = "TURBOSSHORT"
const PortfolioContractTypeElemUPORDOWN PortfolioContractTypeElem = "UPORDOWN"
const PortfolioContractTypeElemVANILLALONGCALL PortfolioContractTypeElem = "VANILLALONGCALL"
const PortfolioContractTypeElemVANILLALONGPUT PortfolioContractTypeElem = "VANILLALONGPUT"

// [Optional] Used to pass data through the websocket, which may be retrieved via
// the `echo_req` output field. Maximum size is 3500 bytes.
type PortfolioPassthrough map[string]interface{}

type PortfolioPortfolio int

var enumValues_PortfolioPortfolio = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PortfolioPortfolio) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_PortfolioPortfolio {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_PortfolioPortfolio, v)
	}
	*j = PortfolioPortfolio(v)
	return nil
}

var enumValues_PortfolioContractTypeElem = []interface{}{
	"ASIAND",
	"ASIANU",
	"CALL",
	"CALLE",
	"CALLSPREAD",
	"DIGITDIFF",
	"DIGITEVEN",
	"DIGITMATCH",
	"DIGITODD",
	"DIGITOVER",
	"DIGITUNDER",
	"EXPIRYMISSE",
	"EXPIRYMISS",
	"EXPIRYRANGE",
	"EXPIRYRANGEE",
	"LBFLOATCALL",
	"LBFLOATPUT",
	"LBHIGHLOW",
	"MULTDOWN",
	"MULTUP",
	"NOTOUCH",
	"ONETOUCH",
	"PUT",
	"PUTE",
	"PUTSPREAD",
	"RANGE",
	"RESETCALL",
	"RESETPUT",
	"RUNHIGH",
	"RUNLOW",
	"TICKHIGH",
	"TICKLOW",
	"UPORDOWN",
	"VANILLALONGCALL",
	"VANILLALONGPUT",
	"TURBOSLONG",
	"TURBOSSHORT",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Portfolio) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["portfolio"]; !ok || v == nil {
		return fmt.Errorf("field portfolio: required")
	}
	type Plain Portfolio
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Portfolio(plain)
	return nil
}