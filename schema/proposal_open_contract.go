// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "fmt"
import "reflect"
import "encoding/json"

// [Optional] Used to pass data through the websocket, which may be retrieved via
// the `echo_req` output field. Maximum size is 3500 bytes.
type ProposalOpenContractPassthrough map[string]interface{}

type ProposalOpenContractProposalOpenContract int

var enumValues_ProposalOpenContractProposalOpenContract = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ProposalOpenContractProposalOpenContract) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ProposalOpenContractProposalOpenContract {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ProposalOpenContractProposalOpenContract, v)
	}
	*j = ProposalOpenContractProposalOpenContract(v)
	return nil
}

type ProposalOpenContractSubscribe int

var enumValues_ProposalOpenContractSubscribe = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ProposalOpenContractSubscribe) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ProposalOpenContractSubscribe {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ProposalOpenContractSubscribe, v)
	}
	*j = ProposalOpenContractSubscribe(v)
	return nil
}

// Get latest price (and other information) for a contract in the user's portfolio
type ProposalOpenContract struct {
	// [Optional] Contract ID received from a `portfolio` request. If not set, you
	// will receive stream of all open contracts.
	ContractId *int `json:"contract_id,omitempty"`

	// [Optional] Used to pass data through the websocket, which may be retrieved via
	// the `echo_req` output field. Maximum size is 3500 bytes.
	Passthrough ProposalOpenContractPassthrough `json:"passthrough,omitempty"`

	// Must be `1`
	ProposalOpenContract ProposalOpenContractProposalOpenContract `json:"proposal_open_contract"`

	// [Optional] Used to map request to response.
	ReqId *int `json:"req_id,omitempty"`

	// [Optional] `1` to stream.
	Subscribe *ProposalOpenContractSubscribe `json:"subscribe,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ProposalOpenContract) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["proposal_open_contract"]; !ok || v == nil {
		return fmt.Errorf("field proposal_open_contract: required")
	}
	type Plain ProposalOpenContract
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ProposalOpenContract(plain)
	return nil
}