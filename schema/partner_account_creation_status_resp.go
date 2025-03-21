// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// A response message containing the status of partner account creation process
type PartnerAccountCreationStatusResp struct {
	// Echo of the request made.
	EchoReq PartnerAccountCreationStatusRespEchoReq `json:"echo_req"`

	// Action name of the request made.
	MsgType PartnerAccountCreationStatusRespMsgType `json:"msg_type"`

	// Status information for the partner account creation process
	PartnerAccountCreationStatus *PartnerAccountCreationStatusRespPartnerAccountCreationStatus `json:"partner_account_creation_status,omitempty"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`
}

// Echo of the request made.
type PartnerAccountCreationStatusRespEchoReq map[string]interface{}

type PartnerAccountCreationStatusRespMsgType string

const PartnerAccountCreationStatusRespMsgTypePartnerAccountCreationStatus PartnerAccountCreationStatusRespMsgType = "partner_account_creation_status"

var enumValues_PartnerAccountCreationStatusRespMsgType = []interface{}{
	"partner_account_creation_status",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PartnerAccountCreationStatusRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_PartnerAccountCreationStatusRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_PartnerAccountCreationStatusRespMsgType, v)
	}
	*j = PartnerAccountCreationStatusRespMsgType(v)
	return nil
}

// Status information for the partner account creation process
type PartnerAccountCreationStatusRespPartnerAccountCreationStatus struct {
	// Status of CFD account creation
	CreateCFDAccount PartnerAccountCreationStatusRespPartnerAccountCreationStatusCreateCFDAccount `json:"create_CFD_account"`

	// Status of EU partner account creation
	CreateEuPartner PartnerAccountCreationStatusRespPartnerAccountCreationStatusCreateEuPartner `json:"create_eu_partner"`

	// Status of ROW partner account creation
	CreateRowPartner PartnerAccountCreationStatusRespPartnerAccountCreationStatusCreateRowPartner `json:"create_row_partner"`

	// Status of third party user creation
	CreateThirdPartyProviderUser PartnerAccountCreationStatusRespPartnerAccountCreationStatusCreateThirdPartyProviderUser `json:"create_third_party_provider_user"`

	// Status of linking EU partner account
	LinkPartnerEu PartnerAccountCreationStatusRespPartnerAccountCreationStatusLinkPartnerEu `json:"link_partner_eu"`

	// Status of linking ROW partner account
	LinkPartnerRow PartnerAccountCreationStatusRespPartnerAccountCreationStatusLinkPartnerRow `json:"link_partner_row"`
}

// Status of CFD account creation
type PartnerAccountCreationStatusRespPartnerAccountCreationStatusCreateCFDAccount struct {
	// Response data if step completed successfully
	Response PartnerAccountCreationStatusRespPartnerAccountCreationStatusCreateCFDAccountResponse `json:"response,omitempty"`

	// Current status of this step
	Status *string `json:"status,omitempty"`
}

// Response data if step completed successfully
type PartnerAccountCreationStatusRespPartnerAccountCreationStatusCreateCFDAccountResponse map[string]interface{}

// Status of EU partner account creation
type PartnerAccountCreationStatusRespPartnerAccountCreationStatusCreateEuPartner struct {
	// Response data if step completed successfully
	Response PartnerAccountCreationStatusRespPartnerAccountCreationStatusCreateEuPartnerResponse `json:"response,omitempty"`

	// Current status of this step
	Status *string `json:"status,omitempty"`
}

// Response data if step completed successfully
type PartnerAccountCreationStatusRespPartnerAccountCreationStatusCreateEuPartnerResponse map[string]interface{}

// Status of ROW partner account creation
type PartnerAccountCreationStatusRespPartnerAccountCreationStatusCreateRowPartner struct {
	// Response data if step completed successfully
	Response PartnerAccountCreationStatusRespPartnerAccountCreationStatusCreateRowPartnerResponse `json:"response,omitempty"`

	// Current status of this step
	Status *string `json:"status,omitempty"`
}

// Response data if step completed successfully
type PartnerAccountCreationStatusRespPartnerAccountCreationStatusCreateRowPartnerResponse map[string]interface{}

// Status of third party user creation
type PartnerAccountCreationStatusRespPartnerAccountCreationStatusCreateThirdPartyProviderUser struct {
	// Response data if step completed successfully
	Response PartnerAccountCreationStatusRespPartnerAccountCreationStatusCreateThirdPartyProviderUserResponse `json:"response,omitempty"`

	// Current status of this step
	Status *string `json:"status,omitempty"`
}

// Response data if step completed successfully
type PartnerAccountCreationStatusRespPartnerAccountCreationStatusCreateThirdPartyProviderUserResponse map[string]interface{}

// Status of linking EU partner account
type PartnerAccountCreationStatusRespPartnerAccountCreationStatusLinkPartnerEu struct {
	// Response data if step completed successfully
	Response PartnerAccountCreationStatusRespPartnerAccountCreationStatusLinkPartnerEuResponse `json:"response,omitempty"`

	// Current status of this step
	Status *string `json:"status,omitempty"`
}

// Response data if step completed successfully
type PartnerAccountCreationStatusRespPartnerAccountCreationStatusLinkPartnerEuResponse map[string]interface{}

// Status of linking ROW partner account
type PartnerAccountCreationStatusRespPartnerAccountCreationStatusLinkPartnerRow struct {
	// Response data if step completed successfully
	Response PartnerAccountCreationStatusRespPartnerAccountCreationStatusLinkPartnerRowResponse `json:"response,omitempty"`

	// Current status of this step
	Status *string `json:"status,omitempty"`
}

// Response data if step completed successfully
type PartnerAccountCreationStatusRespPartnerAccountCreationStatusLinkPartnerRowResponse map[string]interface{}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PartnerAccountCreationStatusRespPartnerAccountCreationStatus) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["create_CFD_account"]; raw != nil && !ok {
		return fmt.Errorf("field create_CFD_account in PartnerAccountCreationStatusRespPartnerAccountCreationStatus: required")
	}
	if _, ok := raw["create_eu_partner"]; raw != nil && !ok {
		return fmt.Errorf("field create_eu_partner in PartnerAccountCreationStatusRespPartnerAccountCreationStatus: required")
	}
	if _, ok := raw["create_row_partner"]; raw != nil && !ok {
		return fmt.Errorf("field create_row_partner in PartnerAccountCreationStatusRespPartnerAccountCreationStatus: required")
	}
	if _, ok := raw["create_third_party_provider_user"]; raw != nil && !ok {
		return fmt.Errorf("field create_third_party_provider_user in PartnerAccountCreationStatusRespPartnerAccountCreationStatus: required")
	}
	if _, ok := raw["link_partner_eu"]; raw != nil && !ok {
		return fmt.Errorf("field link_partner_eu in PartnerAccountCreationStatusRespPartnerAccountCreationStatus: required")
	}
	if _, ok := raw["link_partner_row"]; raw != nil && !ok {
		return fmt.Errorf("field link_partner_row in PartnerAccountCreationStatusRespPartnerAccountCreationStatus: required")
	}
	type Plain PartnerAccountCreationStatusRespPartnerAccountCreationStatus
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = PartnerAccountCreationStatusRespPartnerAccountCreationStatus(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PartnerAccountCreationStatusResp) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["echo_req"]; raw != nil && !ok {
		return fmt.Errorf("field echo_req in PartnerAccountCreationStatusResp: required")
	}
	if _, ok := raw["msg_type"]; raw != nil && !ok {
		return fmt.Errorf("field msg_type in PartnerAccountCreationStatusResp: required")
	}
	type Plain PartnerAccountCreationStatusResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = PartnerAccountCreationStatusResp(plain)
	return nil
}
