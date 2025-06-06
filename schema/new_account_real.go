// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// This call opens a new real-money account. This call can be made from a
// virtual-money or a real-money account. If it is the latter, client information
// fields in this call will be ignored and data from your existing real-money
// account will be used.
type NewAccountReal struct {
	// [Optional] Purpose and reason for requesting the account opening.
	AccountOpeningReason *NewAccountRealAccountOpeningReason `json:"account_opening_reason,omitempty"`

	// [Optional] The anticipated account turnover.
	AccountTurnover *NewAccountRealAccountTurnover `json:"account_turnover,omitempty"`

	// [Optional] Within 100 characters.
	AddressCity *string `json:"address_city,omitempty"`

	// Within 70 characters, with no leading whitespaces and may contain
	// letters/numbers and/or any of following characters '.,:;()@#/-
	AddressLine1 *string `json:"address_line_1,omitempty"`

	// [Optional] Within 70 characters.
	AddressLine2 *string `json:"address_line_2,omitempty"`

	// [Optional] Within 20 characters and may not contain '+'.
	AddressPostcode *string `json:"address_postcode,omitempty"`

	// [Optional] Possible value receive from `states_list` call.
	AddressState *string `json:"address_state,omitempty"`

	// [Optional] Affiliate token, within 32 characters.
	AffiliateToken *string `json:"affiliate_token,omitempty"`

	// [Optional] The phone's calling country code. Don't include the `+` sign. Up to
	// 4 digits.
	CallingCountryCode *string `json:"calling_country_code,omitempty"`

	// [Optional] Country of legal citizenship, 2-letter country code.
	Citizen *string `json:"citizen,omitempty"`

	// [Optional] Indicates whether this is for a client requesting an account with
	// professional status.
	ClientType NewAccountRealClientType `json:"client_type,omitempty"`

	// [Optional] To set currency of the account. List of supported currencies can be
	// acquired with `payout_currencies` call.
	Currency *string `json:"currency,omitempty"`

	// Date of birth format: `yyyy-mm-dd`.
	DateOfBirth *string `json:"date_of_birth,omitempty"`

	// Employment Status.
	EmploymentStatus *string `json:"employment_status,omitempty"`

	// [Optional] Indicates client's self-declaration of FATCA.
	FatcaDeclaration *NewAccountRealFatcaDeclaration `json:"fatca_declaration,omitempty"`

	// [Optional] The version of the financial information form.
	FinancialInformationVersion *string `json:"financial_information_version,omitempty"`

	// Within 1-50 characters, use only letters, spaces, hyphens, full-stops or
	// apostrophes.
	FirstName *string `json:"first_name,omitempty"`

	// Within 1-50 characters, use only letters, spaces, hyphens, full-stops or
	// apostrophes.
	LastName *string `json:"last_name,omitempty"`

	// [Optional] The login id of the user. Mandatory when multiple tokens were
	// provided during authorize.
	Loginid *string `json:"loginid,omitempty"`

	// Must be `1`
	NewAccountReal NewAccountRealNewAccountReal `json:"new_account_real"`

	// [Optional] Indicates client's self-declaration of not being a PEP/RCA
	// (Politically Exposed Person/Relatives and Close Associates).
	NonPepDeclaration *int `json:"non_pep_declaration,omitempty"`

	// [Optional] Used to pass data through the websocket, which may be retrieved via
	// the `echo_req` output field.
	Passthrough NewAccountRealPassthrough `json:"passthrough,omitempty"`

	// [Optional] The phone's national format, don't include the `+` sign nor the
	// calling country code. Up to 15 digits are allowed.
	Phone *string `json:"phone,omitempty"`

	// [Optional] Place of birth, 2-letter country code.
	PlaceOfBirth *string `json:"place_of_birth,omitempty"`

	// [Optional] Used to map request to response.
	ReqId *int `json:"req_id,omitempty"`

	// 2-letter country code, possible value receive from `residence_list` call.
	Residence *string `json:"residence,omitempty"`

	// [Optional] Accept any value in enum list.
	Salutation *NewAccountRealSalutation `json:"salutation,omitempty"`

	// [Optional] Answer to secret question, within 4-50 characters. Required for new
	// account and existing client details will be used if client open another
	// account.
	SecretAnswer *string `json:"secret_answer,omitempty"`

	// [Optional] Accept any value in enum list. Required for new account and existing
	// client details will be used if client open another account.
	SecretQuestion *NewAccountRealSecretQuestion `json:"secret_question,omitempty"`

	// [Optional] Tax identification number. Only applicable for real money account.
	// Required for `maltainvest` landing company.
	TaxIdentificationNumber *string `json:"tax_identification_number,omitempty"`

	// [Optional] Residence for tax purpose. Comma separated iso country code if
	// multiple jurisdictions. Only applicable for real money account. Required for
	// `maltainvest` landing company.
	TaxResidence *string `json:"tax_residence,omitempty"`

	// [Optional] Whether the client has skipped the TIN form. Only applicable for
	// real money account.
	TinSkipped *NewAccountRealTinSkipped `json:"tin_skipped,omitempty"`

	// The tnc acceptance status of the user.
	TncAcceptance *NewAccountRealTncAcceptance `json:"tnc_acceptance,omitempty"`
}

type NewAccountRealAccountOpeningReason string

const NewAccountRealAccountOpeningReasonAdditionalRevenue NewAccountRealAccountOpeningReason = "Additional revenue"
const NewAccountRealAccountOpeningReasonHedging NewAccountRealAccountOpeningReason = "Hedging"
const NewAccountRealAccountOpeningReasonIncomeEarning NewAccountRealAccountOpeningReason = "Income Earning"
const NewAccountRealAccountOpeningReasonPeerToPeerExchange NewAccountRealAccountOpeningReason = "Peer-to-peer exchange"
const NewAccountRealAccountOpeningReasonSavings NewAccountRealAccountOpeningReason = "Savings"
const NewAccountRealAccountOpeningReasonSpeculative NewAccountRealAccountOpeningReason = "Speculative"

var enumValues_NewAccountRealAccountOpeningReason = []interface{}{
	"Speculative",
	"Income Earning",
	"Hedging",
	"Peer-to-peer exchange",
	"Additional revenue",
	"Savings",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NewAccountRealAccountOpeningReason) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_NewAccountRealAccountOpeningReason {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_NewAccountRealAccountOpeningReason, v)
	}
	*j = NewAccountRealAccountOpeningReason(v)
	return nil
}

type NewAccountRealAccountTurnover string

const NewAccountRealAccountTurnoverA100001500000 NewAccountRealAccountTurnover = "$100,001 - $500,000"
const NewAccountRealAccountTurnoverA2500050000 NewAccountRealAccountTurnover = "$25,000 - $50,000"
const NewAccountRealAccountTurnoverA50001100000 NewAccountRealAccountTurnover = "$50,001 - $100,000"
const NewAccountRealAccountTurnoverLessThan25000 NewAccountRealAccountTurnover = "Less than $25,000"
const NewAccountRealAccountTurnoverOver500000 NewAccountRealAccountTurnover = "Over $500,000"

var enumValues_NewAccountRealAccountTurnover = []interface{}{
	"Less than $25,000",
	"$25,000 - $50,000",
	"$50,001 - $100,000",
	"$100,001 - $500,000",
	"Over $500,000",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NewAccountRealAccountTurnover) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_NewAccountRealAccountTurnover {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_NewAccountRealAccountTurnover, v)
	}
	*j = NewAccountRealAccountTurnover(v)
	return nil
}

type NewAccountRealClientType string

const NewAccountRealClientTypeProfessional NewAccountRealClientType = "professional"
const NewAccountRealClientTypeRetail NewAccountRealClientType = "retail"

var enumValues_NewAccountRealClientType = []interface{}{
	"professional",
	"retail",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NewAccountRealClientType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_NewAccountRealClientType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_NewAccountRealClientType, v)
	}
	*j = NewAccountRealClientType(v)
	return nil
}

type NewAccountRealFatcaDeclaration int

var enumValues_NewAccountRealFatcaDeclaration = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NewAccountRealFatcaDeclaration) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_NewAccountRealFatcaDeclaration {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_NewAccountRealFatcaDeclaration, v)
	}
	*j = NewAccountRealFatcaDeclaration(v)
	return nil
}

type NewAccountRealNewAccountReal int

var enumValues_NewAccountRealNewAccountReal = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NewAccountRealNewAccountReal) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_NewAccountRealNewAccountReal {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_NewAccountRealNewAccountReal, v)
	}
	*j = NewAccountRealNewAccountReal(v)
	return nil
}

// [Optional] Used to pass data through the websocket, which may be retrieved via
// the `echo_req` output field.
type NewAccountRealPassthrough map[string]interface{}

type NewAccountRealSalutation string

const NewAccountRealSalutationMiss NewAccountRealSalutation = "Miss"
const NewAccountRealSalutationMr NewAccountRealSalutation = "Mr"
const NewAccountRealSalutationMrs NewAccountRealSalutation = "Mrs"
const NewAccountRealSalutationMs NewAccountRealSalutation = "Ms"

var enumValues_NewAccountRealSalutation = []interface{}{
	"Mr",
	"Ms",
	"Miss",
	"Mrs",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NewAccountRealSalutation) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_NewAccountRealSalutation {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_NewAccountRealSalutation, v)
	}
	*j = NewAccountRealSalutation(v)
	return nil
}

type NewAccountRealSecretQuestion string

const NewAccountRealSecretQuestionBrandOfFirstCar NewAccountRealSecretQuestion = "Brand of first car"
const NewAccountRealSecretQuestionFavouriteArtist NewAccountRealSecretQuestion = "Favourite artist"
const NewAccountRealSecretQuestionFavouriteDish NewAccountRealSecretQuestion = "Favourite dish"
const NewAccountRealSecretQuestionMemorableDate NewAccountRealSecretQuestion = "Memorable date"
const NewAccountRealSecretQuestionMemorableTownCity NewAccountRealSecretQuestion = "Memorable town/city"
const NewAccountRealSecretQuestionMotherSMaidenName NewAccountRealSecretQuestion = "Mother's maiden name"
const NewAccountRealSecretQuestionNameOfFirstLove NewAccountRealSecretQuestion = "Name of first love"
const NewAccountRealSecretQuestionNameOfYourPet NewAccountRealSecretQuestion = "Name of your pet"

var enumValues_NewAccountRealSecretQuestion = []interface{}{
	"Mother's maiden name",
	"Name of your pet",
	"Name of first love",
	"Memorable town/city",
	"Memorable date",
	"Favourite dish",
	"Brand of first car",
	"Favourite artist",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NewAccountRealSecretQuestion) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_NewAccountRealSecretQuestion {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_NewAccountRealSecretQuestion, v)
	}
	*j = NewAccountRealSecretQuestion(v)
	return nil
}

type NewAccountRealTinSkipped int

var enumValues_NewAccountRealTinSkipped = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NewAccountRealTinSkipped) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_NewAccountRealTinSkipped {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_NewAccountRealTinSkipped, v)
	}
	*j = NewAccountRealTinSkipped(v)
	return nil
}

type NewAccountRealTncAcceptance int

var enumValues_NewAccountRealTncAcceptance = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NewAccountRealTncAcceptance) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_NewAccountRealTncAcceptance {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_NewAccountRealTncAcceptance, v)
	}
	*j = NewAccountRealTncAcceptance(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NewAccountReal) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["new_account_real"]; raw != nil && !ok {
		return fmt.Errorf("field new_account_real in NewAccountReal: required")
	}
	type Plain NewAccountReal
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["client_type"]; !ok || v == nil {
		plain.ClientType = "retail"
	}
	if plain.SecretAnswer != nil && len(*plain.SecretAnswer) < 4 {
		return fmt.Errorf("field %s length: must be >= %d", "secret_answer", 4)
	}
	if plain.SecretAnswer != nil && len(*plain.SecretAnswer) > 50 {
		return fmt.Errorf("field %s length: must be <= %d", "secret_answer", 50)
	}
	*j = NewAccountReal(plain)
	return nil
}
