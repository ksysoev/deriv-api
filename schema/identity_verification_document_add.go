// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// Adds document information such as issuing country, id and type for identity
// verification processes.
type IdentityVerificationDocumentAdd struct {
	// [Optional] Additional info required by some document types.
	DocumentAdditional *string `json:"document_additional,omitempty"`

	// The identification number of the document.
	DocumentNumber string `json:"document_number"`

	// The type of the document based on provided `issuing_country` (can obtained from
	// `residence_list` call).
	DocumentType string `json:"document_type"`

	// Must be `1`
	IdentityVerificationDocumentAdd IdentityVerificationDocumentAddIdentityVerificationDocumentAdd `json:"identity_verification_document_add"`

	// 2-letter country code (can obtained from `residence_list` call).
	IssuingCountry string `json:"issuing_country"`

	// [Optional] The login id of the user. Mandatory when multiple tokens were
	// provided during authorize.
	Loginid *string `json:"loginid,omitempty"`

	// [Optional] Used to pass data through the websocket, which may be retrieved via
	// the `echo_req` output field.
	Passthrough IdentityVerificationDocumentAddPassthrough `json:"passthrough,omitempty"`

	// [Optional] Used to map request to response.
	ReqId *int `json:"req_id,omitempty"`
}

type IdentityVerificationDocumentAddIdentityVerificationDocumentAdd int

var enumValues_IdentityVerificationDocumentAddIdentityVerificationDocumentAdd = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdentityVerificationDocumentAddIdentityVerificationDocumentAdd) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_IdentityVerificationDocumentAddIdentityVerificationDocumentAdd {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_IdentityVerificationDocumentAddIdentityVerificationDocumentAdd, v)
	}
	*j = IdentityVerificationDocumentAddIdentityVerificationDocumentAdd(v)
	return nil
}

// [Optional] Used to pass data through the websocket, which may be retrieved via
// the `echo_req` output field.
type IdentityVerificationDocumentAddPassthrough map[string]interface{}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdentityVerificationDocumentAdd) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["document_number"]; raw != nil && !ok {
		return fmt.Errorf("field document_number in IdentityVerificationDocumentAdd: required")
	}
	if _, ok := raw["document_type"]; raw != nil && !ok {
		return fmt.Errorf("field document_type in IdentityVerificationDocumentAdd: required")
	}
	if _, ok := raw["identity_verification_document_add"]; raw != nil && !ok {
		return fmt.Errorf("field identity_verification_document_add in IdentityVerificationDocumentAdd: required")
	}
	if _, ok := raw["issuing_country"]; raw != nil && !ok {
		return fmt.Errorf("field issuing_country in IdentityVerificationDocumentAdd: required")
	}
	type Plain IdentityVerificationDocumentAdd
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = IdentityVerificationDocumentAdd(plain)
	return nil
}
