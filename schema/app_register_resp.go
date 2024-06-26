// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// A message with created application details
type AppRegisterResp struct {
	// The information of the created application.
	AppRegister *AppRegisterRespAppRegister `json:"app_register,omitempty"`

	// Echo of the request made.
	EchoReq AppRegisterRespEchoReq `json:"echo_req"`

	// Action name of the request made.
	MsgType AppRegisterRespMsgType `json:"msg_type"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`
}

// The information of the created application.
type AppRegisterRespAppRegister struct {
	// Active.
	Active *int `json:"active,omitempty"`

	// Application ID.
	AppId int `json:"app_id"`

	// Markup added to contract prices (as a percentage of contract payout).
	AppMarkupPercentage float64 `json:"app_markup_percentage"`

	// Application's App Store URL.
	Appstore string `json:"appstore"`

	// Application's GitHub page (for open-source projects).
	Github string `json:"github"`

	// Application's Google Play URL.
	Googleplay string `json:"googleplay"`

	// Application's homepage URL.
	Homepage string `json:"homepage"`

	// Application name.
	Name string `json:"name"`

	// The URL to redirect to after a successful login.
	RedirectUri string `json:"redirect_uri"`

	// Scope Details.
	Scopes []string `json:"scopes,omitempty"`

	// Used when `verify_email` called. If available, a URL containing the
	// verification token will send to the client's email, otherwise only the token
	// will be sent.
	VerificationUri string `json:"verification_uri"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AppRegisterRespAppRegister) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["app_id"]; raw != nil && !ok {
		return fmt.Errorf("field app_id in AppRegisterRespAppRegister: required")
	}
	if _, ok := raw["app_markup_percentage"]; raw != nil && !ok {
		return fmt.Errorf("field app_markup_percentage in AppRegisterRespAppRegister: required")
	}
	if _, ok := raw["appstore"]; raw != nil && !ok {
		return fmt.Errorf("field appstore in AppRegisterRespAppRegister: required")
	}
	if _, ok := raw["github"]; raw != nil && !ok {
		return fmt.Errorf("field github in AppRegisterRespAppRegister: required")
	}
	if _, ok := raw["googleplay"]; raw != nil && !ok {
		return fmt.Errorf("field googleplay in AppRegisterRespAppRegister: required")
	}
	if _, ok := raw["homepage"]; raw != nil && !ok {
		return fmt.Errorf("field homepage in AppRegisterRespAppRegister: required")
	}
	if _, ok := raw["name"]; raw != nil && !ok {
		return fmt.Errorf("field name in AppRegisterRespAppRegister: required")
	}
	if _, ok := raw["redirect_uri"]; raw != nil && !ok {
		return fmt.Errorf("field redirect_uri in AppRegisterRespAppRegister: required")
	}
	if _, ok := raw["verification_uri"]; raw != nil && !ok {
		return fmt.Errorf("field verification_uri in AppRegisterRespAppRegister: required")
	}
	type Plain AppRegisterRespAppRegister
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = AppRegisterRespAppRegister(plain)
	return nil
}

// Echo of the request made.
type AppRegisterRespEchoReq map[string]interface{}

type AppRegisterRespMsgType string

const AppRegisterRespMsgTypeAppRegister AppRegisterRespMsgType = "app_register"

var enumValues_AppRegisterRespMsgType = []interface{}{
	"app_register",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AppRegisterRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AppRegisterRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AppRegisterRespMsgType, v)
	}
	*j = AppRegisterRespMsgType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AppRegisterResp) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["echo_req"]; raw != nil && !ok {
		return fmt.Errorf("field echo_req in AppRegisterResp: required")
	}
	if _, ok := raw["msg_type"]; raw != nil && !ok {
		return fmt.Errorf("field msg_type in AppRegisterResp: required")
	}
	type Plain AppRegisterResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = AppRegisterResp(plain)
	return nil
}
