// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "fmt"
import "reflect"
import "encoding/json"

// Echo of the request made.
type NewAccountRealRespEchoReq map[string]interface{}

type NewAccountRealRespMsgType string

var enumValues_NewAccountRealRespMsgType = []interface{}{
	"new_account_real",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NewAccountRealRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_NewAccountRealRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_NewAccountRealRespMsgType, v)
	}
	*j = NewAccountRealRespMsgType(v)
	return nil
}

const NewAccountRealRespMsgTypeNewAccountReal NewAccountRealRespMsgType = "new_account_real"

// New real money account details
type NewAccountRealRespNewAccountReal struct {
	// Client ID of new real money account
	ClientId string `json:"client_id"`

	// Currency of an account
	Currency *string `json:"currency,omitempty"`

	// Landing company full name
	LandingCompany string `json:"landing_company"`

	// Landing company shortcode
	LandingCompanyShort *string `json:"landing_company_short,omitempty"`

	// Landing company shortcode
	LandingCompanyShortcode *string `json:"landing_company_shortcode,omitempty"`

	// OAuth token for client's login session
	OauthToken string `json:"oauth_token"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NewAccountRealRespNewAccountReal) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["client_id"]; !ok || v == nil {
		return fmt.Errorf("field client_id: required")
	}
	if v, ok := raw["landing_company"]; !ok || v == nil {
		return fmt.Errorf("field landing_company: required")
	}
	if v, ok := raw["oauth_token"]; !ok || v == nil {
		return fmt.Errorf("field oauth_token: required")
	}
	type Plain NewAccountRealRespNewAccountReal
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = NewAccountRealRespNewAccountReal(plain)
	return nil
}

// Create real account Receive
type NewAccountRealResp struct {
	// Echo of the request made.
	EchoReq NewAccountRealRespEchoReq `json:"echo_req"`

	// Action name of the request made.
	MsgType NewAccountRealRespMsgType `json:"msg_type"`

	// New real money account details
	NewAccountReal *NewAccountRealRespNewAccountReal `json:"new_account_real,omitempty"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NewAccountRealResp) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["echo_req"]; !ok || v == nil {
		return fmt.Errorf("field echo_req: required")
	}
	if v, ok := raw["msg_type"]; !ok || v == nil {
		return fmt.Errorf("field msg_type: required")
	}
	type Plain NewAccountRealResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = NewAccountRealResp(plain)
	return nil
}