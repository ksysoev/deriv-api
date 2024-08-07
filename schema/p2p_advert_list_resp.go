// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// Available adverts matching the requested criteria.
type P2PAdvertListResp struct {
	// Echo of the request made.
	EchoReq P2PAdvertListRespEchoReq `json:"echo_req"`

	// Action name of the request made.
	MsgType P2PAdvertListRespMsgType `json:"msg_type"`

	// P2P adverts list.
	P2PAdvertList *P2PAdvertListRespP2PAdvertList `json:"p2p_advert_list,omitempty"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`
}

// Echo of the request made.
type P2PAdvertListRespEchoReq map[string]interface{}

type P2PAdvertListRespMsgType string

const P2PAdvertListRespMsgTypeP2PAdvertList P2PAdvertListRespMsgType = "p2p_advert_list"

var enumValues_P2PAdvertListRespMsgType = []interface{}{
	"p2p_advert_list",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertListRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertListRespMsgType, v)
	}
	*j = P2PAdvertListRespMsgType(v)
	return nil
}

// P2P adverts list.
type P2PAdvertListRespP2PAdvertList struct {
	// List of adverts.
	List []P2PAdvertListRespP2PAdvertListListElem `json:"list"`
}

type P2PAdvertListRespP2PAdvertListListElem struct {
	// Currency for this advert. This is the system currency to be transferred between
	// advertiser and client.
	AccountCurrency string `json:"account_currency"`

	// The number of active orders against this advert.
	ActiveOrders *int `json:"active_orders,omitempty"`

	// Details of the advertiser for this advert.
	AdvertiserDetails P2PAdvertListRespP2PAdvertListListElemAdvertiserDetails `json:"advertiser_details"`

	// The total amount specified in advert, in `account_currency`. It is only visible
	// to the advert owner.
	Amount *float64 `json:"amount,omitempty"`

	// The total amount specified in advert, in `account_currency`, formatted to
	// appropriate decimal places. It is only visible to the advert owner.
	AmountDisplay *string `json:"amount_display,omitempty"`

	// Indicates if this is block trade advert or not.
	BlockTrade P2PAdvertListRespP2PAdvertListListElemBlockTrade `json:"block_trade"`

	// Advertiser contact information. Only applicable for 'sell adverts'.
	ContactInfo *string `json:"contact_info,omitempty"`

	// Type of transaction from the opposite party's perspective.
	CounterpartyType P2PAdvertListRespP2PAdvertListListElemCounterpartyType `json:"counterparty_type"`

	// The target country code of the advert.
	Country string `json:"country"`

	// The advert creation time in epoch.
	CreatedTime int `json:"created_time"`

	// Days until automatic inactivation of this ad, if no activity occurs.
	DaysUntilArchive *int `json:"days_until_archive,omitempty"`

	// General information about the advert.
	Description string `json:"description"`

	// Conversion rate from account currency to local currency, using current market
	// rate if applicable.
	EffectiveRate *float64 `json:"effective_rate"`

	// Conversion rate from account currency to local currency, using current market
	// rate if applicable, formatted to appropriate decimal places.
	EffectiveRateDisplay *string `json:"effective_rate_display"`

	// Reasons why the counterparty terms do not allow the current user to place
	// orders against this advert. Possible values:
	// - `completion_rate`: current user's 30 day completion rate is less than
	// `min_completion_rate`.
	// - `country`: current user's residence is not in `eligible_countries`.
	// - `join_date`: current user registered on P2P less than `min_join_days` in the
	// past.
	// - `rating`: current user's average review rating is less than `min_rating`.
	EligibilityStatus []P2PAdvertListRespP2PAdvertListListElemEligibilityStatusElem `json:"eligibility_status,omitempty"`

	// 2 letter country codes. Counterparties who do not live in these countries are
	// not allowed to place orders against this advert.
	EligibleCountries []string `json:"eligible_countries,omitempty"`

	// The unique identifier for this advert.
	Id string `json:"id"`

	// The activation status of the advert.
	IsActive P2PAdvertListRespP2PAdvertListListElemIsActive `json:"is_active"`

	// Inidcates whether the current user's schedule has availability between now and
	// now + order_expiry_period.
	IsClientScheduleAvailable *P2PAdvertListRespP2PAdvertListListElemIsClientScheduleAvailable `json:"is_client_schedule_available,omitempty"`

	// Indicates that the current user meets the counterparty terms for placing an
	// order.
	IsEligible P2PAdvertListRespP2PAdvertListListElemIsEligible `json:"is_eligible,omitempty"`

	// Indicates that this advert will appear on the main advert list.
	IsVisible P2PAdvertListRespP2PAdvertListListElemIsVisible `json:"is_visible"`

	// Local currency for this advert. This is the form of payment to be arranged
	// directly between advertiser and client.
	LocalCurrency string `json:"local_currency"`

	// Maximum order amount specified in advert, in `account_currency`. It is only
	// visible for advertisers.
	MaxOrderAmount *float64 `json:"max_order_amount,omitempty"`

	// Maximum order amount specified in advert, in `account_currency`, formatted to
	// appropriate decimal places. It is only visible to the advert owner.
	MaxOrderAmountDisplay *string `json:"max_order_amount_display,omitempty"`

	// Maximum order amount at this time, in `account_currency`.
	MaxOrderAmountLimit float64 `json:"max_order_amount_limit"`

	// Maximum order amount at this time, in `account_currency`, formatted to
	// appropriate decimal places.
	MaxOrderAmountLimitDisplay string `json:"max_order_amount_limit_display"`

	// Counterparties who have a 30 day completion rate less than this value are not
	// allowed to place orders against this advert.
	MinCompletionRate *float64 `json:"min_completion_rate,omitempty"`

	// Counterparties who joined less than this number of days ago are not allowed to
	// place orders against this advert.
	MinJoinDays *int `json:"min_join_days,omitempty"`

	// Minimum order amount specified in advert, in `account_currency`. It is only
	// visible for advertisers.
	MinOrderAmount *float64 `json:"min_order_amount,omitempty"`

	// Minimum order amount specified in advert, in `account_currency`, formatted to
	// appropriate decimal places. It is only visible to the advert owner.
	MinOrderAmountDisplay *string `json:"min_order_amount_display,omitempty"`

	// Minimum order amount at this time, in `account_currency`.
	MinOrderAmountLimit float64 `json:"min_order_amount_limit"`

	// Minimum order amount at this time, in `account_currency`, formatted to
	// appropriate decimal places.
	MinOrderAmountLimitDisplay string `json:"min_order_amount_limit_display"`

	// Counterparties who have an average rating less than this value are not allowed
	// to place orders against this advert.
	MinRating *float64 `json:"min_rating,omitempty"`

	// Expiry period (seconds) for order created against this ad.
	OrderExpiryPeriod int `json:"order_expiry_period"`

	// Payment instructions. Only applicable for 'sell adverts'.
	PaymentInfo *string `json:"payment_info,omitempty"`

	// Payment method name (deprecated).
	PaymentMethod *string `json:"payment_method"`

	// Names of supported payment methods.
	PaymentMethodNames []string `json:"payment_method_names,omitempty"`

	// Cost of the advert in local currency.
	Price *float64 `json:"price"`

	// Cost of the advert in local currency, formatted to appropriate decimal places.
	PriceDisplay *string `json:"price_display"`

	// Conversion rate from advertiser's account currency to `local_currency`. An
	// absolute rate value (fixed), or percentage offset from current market rate
	// (floating).
	Rate float64 `json:"rate"`

	// Conversion rate formatted to appropriate decimal places.
	RateDisplay string `json:"rate_display"`

	// Type of rate, fixed or floating.
	RateType P2PAdvertListRespP2PAdvertListListElemRateType `json:"rate_type"`

	// Amount currently available for orders, in `account_currency`. It is only
	// visible to the advert owner.
	RemainingAmount *float64 `json:"remaining_amount,omitempty"`

	// Amount currently available for orders, in `account_currency`, formatted to
	// appropriate decimal places. It is only visible to the advert owner.
	RemainingAmountDisplay *string `json:"remaining_amount_display,omitempty"`

	// Whether this is a buy or a sell.
	Type P2PAdvertListRespP2PAdvertListListElemType `json:"type"`

	// Reasons why an advert is not visible. Possible values:
	// - `advert_fixed_rate_disabled`: fixed rate adverts are no longer available in
	// the advert's country.
	// - `advert_float_rate_disabled`: floating rate adverts are no longer available
	// in the advert's country.
	// - `advert_inactive`: the advert is set inactive.
	// - `advert_max_limit`: the minimum order amount exceeds the system maximum
	// order.
	// - `advert_min_limit`: the maximum order amount is too small to be shown on the
	// advert list.
	// - `advert_remaining`: the remaining amount of the advert is below the minimum
	// order.
	// - `advert_no_payment_methods`: the advert has no valid payment methods.
	// - `advertiser_ads_paused`: the advertiser has paused all adverts.
	// - `advertiser_approval`: the advertiser's proof of identity is not verified.
	// - `advertiser_balance`: the advertiser's P2P balance is less than the minimum
	// order.
	// - `advertiser_schedule`: the advertiser's schedule does not have availability
	// between now and now + order_expiry_period.
	// - `advertiser_block_trade_ineligible`: the advertiser is not currently eligible
	// for block trading.
	// - `advertiser_daily_limit`: the advertiser's remaining daily limit is less than
	// the minimum order.
	// - `advertiser_temp_ban`: the advertiser is temporarily banned from P2P.
	VisibilityStatus []P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElem `json:"visibility_status,omitempty"`
}

// Details of the advertiser for this advert.
type P2PAdvertListRespP2PAdvertListListElemAdvertiserDetails struct {
	// The total number of orders completed in the past 30 days.
	CompletedOrdersCount int `json:"completed_orders_count"`

	// The advertiser's first name.
	FirstName *string `json:"first_name,omitempty"`

	// The advertiser's unique identifier.
	Id string `json:"id"`

	// Indicates that the advertiser is blocked by the current user.
	IsBlocked *P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsBlocked `json:"is_blocked,omitempty"`

	// Indicates that the advertiser is a favourite.
	IsFavourite *P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsFavourite `json:"is_favourite,omitempty"`

	// Indicates if the advertiser is currently online.
	IsOnline P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsOnline `json:"is_online"`

	// Indicates that the advertiser was recommended in the most recent review by the
	// current user.
	IsRecommended *P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsRecommended `json:"is_recommended,omitempty"`

	// Inidcates whether the advertiser's schedule has availability between now and
	// now + order_expiry_period.
	IsScheduleAvailable P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsScheduleAvailable `json:"is_schedule_available"`

	// The advertiser's last name.
	LastName *string `json:"last_name,omitempty"`

	// Epoch of the latest time the advertiser was online, up to 6 months.
	LastOnlineTime *int `json:"last_online_time"`

	// The advertiser's displayed name.
	Name string `json:"name"`

	// Average rating of the advertiser, range is 1-5.
	RatingAverage *float64 `json:"rating_average"`

	// Number of ratings given to the advertiser.
	RatingCount int `json:"rating_count"`

	// Percentage of users who have recommended the advertiser.
	RecommendedAverage *float64 `json:"recommended_average"`

	// Number of times the advertiser has been recommended.
	RecommendedCount *float64 `json:"recommended_count"`

	// The percentage of successfully completed orders made by or placed against the
	// advertiser within the past 30 days.
	TotalCompletionRate *float64 `json:"total_completion_rate"`
}

type P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsBlocked int

var enumValues_P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsBlocked = []interface{}{
	0,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsBlocked) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsBlocked {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsBlocked, v)
	}
	*j = P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsBlocked(v)
	return nil
}

type P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsFavourite int

var enumValues_P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsFavourite = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsFavourite) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsFavourite {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsFavourite, v)
	}
	*j = P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsFavourite(v)
	return nil
}

type P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsOnline int

var enumValues_P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsOnline = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsOnline) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsOnline {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsOnline, v)
	}
	*j = P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsOnline(v)
	return nil
}

type P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsRecommended struct {
	Value interface{}
}

// MarshalJSON implements json.Marshaler.
func (j *P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsRecommended) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Value)
}

var enumValues_P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsRecommended = []interface{}{
	nil,
	0.0,
	1.0,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsRecommended) UnmarshalJSON(b []byte) error {
	var v struct {
		Value interface{}
	}
	if err := json.Unmarshal(b, &v.Value); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsRecommended {
		if reflect.DeepEqual(v.Value, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsRecommended, v.Value)
	}
	*j = P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsRecommended(v)
	return nil
}

type P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsScheduleAvailable int

var enumValues_P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsScheduleAvailable = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsScheduleAvailable) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsScheduleAvailable {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsScheduleAvailable, v)
	}
	*j = P2PAdvertListRespP2PAdvertListListElemAdvertiserDetailsIsScheduleAvailable(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListRespP2PAdvertListListElemAdvertiserDetails) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["completed_orders_count"]; raw != nil && !ok {
		return fmt.Errorf("field completed_orders_count in P2PAdvertListRespP2PAdvertListListElemAdvertiserDetails: required")
	}
	if _, ok := raw["id"]; raw != nil && !ok {
		return fmt.Errorf("field id in P2PAdvertListRespP2PAdvertListListElemAdvertiserDetails: required")
	}
	if _, ok := raw["is_online"]; raw != nil && !ok {
		return fmt.Errorf("field is_online in P2PAdvertListRespP2PAdvertListListElemAdvertiserDetails: required")
	}
	if _, ok := raw["is_schedule_available"]; raw != nil && !ok {
		return fmt.Errorf("field is_schedule_available in P2PAdvertListRespP2PAdvertListListElemAdvertiserDetails: required")
	}
	if _, ok := raw["last_online_time"]; raw != nil && !ok {
		return fmt.Errorf("field last_online_time in P2PAdvertListRespP2PAdvertListListElemAdvertiserDetails: required")
	}
	if _, ok := raw["name"]; raw != nil && !ok {
		return fmt.Errorf("field name in P2PAdvertListRespP2PAdvertListListElemAdvertiserDetails: required")
	}
	if _, ok := raw["rating_average"]; raw != nil && !ok {
		return fmt.Errorf("field rating_average in P2PAdvertListRespP2PAdvertListListElemAdvertiserDetails: required")
	}
	if _, ok := raw["rating_count"]; raw != nil && !ok {
		return fmt.Errorf("field rating_count in P2PAdvertListRespP2PAdvertListListElemAdvertiserDetails: required")
	}
	if _, ok := raw["recommended_average"]; raw != nil && !ok {
		return fmt.Errorf("field recommended_average in P2PAdvertListRespP2PAdvertListListElemAdvertiserDetails: required")
	}
	if _, ok := raw["recommended_count"]; raw != nil && !ok {
		return fmt.Errorf("field recommended_count in P2PAdvertListRespP2PAdvertListListElemAdvertiserDetails: required")
	}
	if _, ok := raw["total_completion_rate"]; raw != nil && !ok {
		return fmt.Errorf("field total_completion_rate in P2PAdvertListRespP2PAdvertListListElemAdvertiserDetails: required")
	}
	type Plain P2PAdvertListRespP2PAdvertListListElemAdvertiserDetails
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = P2PAdvertListRespP2PAdvertListListElemAdvertiserDetails(plain)
	return nil
}

type P2PAdvertListRespP2PAdvertListListElemBlockTrade int

var enumValues_P2PAdvertListRespP2PAdvertListListElemBlockTrade = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListRespP2PAdvertListListElemBlockTrade) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertListRespP2PAdvertListListElemBlockTrade {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertListRespP2PAdvertListListElemBlockTrade, v)
	}
	*j = P2PAdvertListRespP2PAdvertListListElemBlockTrade(v)
	return nil
}

type P2PAdvertListRespP2PAdvertListListElemCounterpartyType string

const P2PAdvertListRespP2PAdvertListListElemCounterpartyTypeBuy P2PAdvertListRespP2PAdvertListListElemCounterpartyType = "buy"
const P2PAdvertListRespP2PAdvertListListElemCounterpartyTypeSell P2PAdvertListRespP2PAdvertListListElemCounterpartyType = "sell"

var enumValues_P2PAdvertListRespP2PAdvertListListElemCounterpartyType = []interface{}{
	"buy",
	"sell",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListRespP2PAdvertListListElemCounterpartyType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertListRespP2PAdvertListListElemCounterpartyType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertListRespP2PAdvertListListElemCounterpartyType, v)
	}
	*j = P2PAdvertListRespP2PAdvertListListElemCounterpartyType(v)
	return nil
}

type P2PAdvertListRespP2PAdvertListListElemEligibilityStatusElem string

const P2PAdvertListRespP2PAdvertListListElemEligibilityStatusElemCompletionRate P2PAdvertListRespP2PAdvertListListElemEligibilityStatusElem = "completion_rate"
const P2PAdvertListRespP2PAdvertListListElemEligibilityStatusElemCountry P2PAdvertListRespP2PAdvertListListElemEligibilityStatusElem = "country"
const P2PAdvertListRespP2PAdvertListListElemEligibilityStatusElemJoinDate P2PAdvertListRespP2PAdvertListListElemEligibilityStatusElem = "join_date"
const P2PAdvertListRespP2PAdvertListListElemEligibilityStatusElemRatingAverage P2PAdvertListRespP2PAdvertListListElemEligibilityStatusElem = "rating_average"

var enumValues_P2PAdvertListRespP2PAdvertListListElemEligibilityStatusElem = []interface{}{
	"completion_rate",
	"country",
	"join_date",
	"rating_average",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListRespP2PAdvertListListElemEligibilityStatusElem) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertListRespP2PAdvertListListElemEligibilityStatusElem {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertListRespP2PAdvertListListElemEligibilityStatusElem, v)
	}
	*j = P2PAdvertListRespP2PAdvertListListElemEligibilityStatusElem(v)
	return nil
}

type P2PAdvertListRespP2PAdvertListListElemIsActive int

var enumValues_P2PAdvertListRespP2PAdvertListListElemIsActive = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListRespP2PAdvertListListElemIsActive) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertListRespP2PAdvertListListElemIsActive {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertListRespP2PAdvertListListElemIsActive, v)
	}
	*j = P2PAdvertListRespP2PAdvertListListElemIsActive(v)
	return nil
}

type P2PAdvertListRespP2PAdvertListListElemIsClientScheduleAvailable int

var enumValues_P2PAdvertListRespP2PAdvertListListElemIsClientScheduleAvailable = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListRespP2PAdvertListListElemIsClientScheduleAvailable) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertListRespP2PAdvertListListElemIsClientScheduleAvailable {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertListRespP2PAdvertListListElemIsClientScheduleAvailable, v)
	}
	*j = P2PAdvertListRespP2PAdvertListListElemIsClientScheduleAvailable(v)
	return nil
}

type P2PAdvertListRespP2PAdvertListListElemIsEligible int

var enumValues_P2PAdvertListRespP2PAdvertListListElemIsEligible = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListRespP2PAdvertListListElemIsEligible) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertListRespP2PAdvertListListElemIsEligible {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertListRespP2PAdvertListListElemIsEligible, v)
	}
	*j = P2PAdvertListRespP2PAdvertListListElemIsEligible(v)
	return nil
}

type P2PAdvertListRespP2PAdvertListListElemIsVisible int

var enumValues_P2PAdvertListRespP2PAdvertListListElemIsVisible = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListRespP2PAdvertListListElemIsVisible) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertListRespP2PAdvertListListElemIsVisible {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertListRespP2PAdvertListListElemIsVisible, v)
	}
	*j = P2PAdvertListRespP2PAdvertListListElemIsVisible(v)
	return nil
}

type P2PAdvertListRespP2PAdvertListListElemRateType string

const P2PAdvertListRespP2PAdvertListListElemRateTypeFixed P2PAdvertListRespP2PAdvertListListElemRateType = "fixed"
const P2PAdvertListRespP2PAdvertListListElemRateTypeFloat P2PAdvertListRespP2PAdvertListListElemRateType = "float"

var enumValues_P2PAdvertListRespP2PAdvertListListElemRateType = []interface{}{
	"fixed",
	"float",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListRespP2PAdvertListListElemRateType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertListRespP2PAdvertListListElemRateType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertListRespP2PAdvertListListElemRateType, v)
	}
	*j = P2PAdvertListRespP2PAdvertListListElemRateType(v)
	return nil
}

type P2PAdvertListRespP2PAdvertListListElemType string

const P2PAdvertListRespP2PAdvertListListElemTypeBuy P2PAdvertListRespP2PAdvertListListElemType = "buy"
const P2PAdvertListRespP2PAdvertListListElemTypeSell P2PAdvertListRespP2PAdvertListListElemType = "sell"

var enumValues_P2PAdvertListRespP2PAdvertListListElemType = []interface{}{
	"buy",
	"sell",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListRespP2PAdvertListListElemType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertListRespP2PAdvertListListElemType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertListRespP2PAdvertListListElemType, v)
	}
	*j = P2PAdvertListRespP2PAdvertListListElemType(v)
	return nil
}

type P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElem string

const P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElemAdvertFixedRateDisabled P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElem = "advert_fixed_rate_disabled"
const P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElemAdvertFloatRateDisabled P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElem = "advert_float_rate_disabled"
const P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElemAdvertInactive P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElem = "advert_inactive"
const P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElemAdvertMaxLimit P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElem = "advert_max_limit"
const P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElemAdvertMinLimit P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElem = "advert_min_limit"
const P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElemAdvertNoPaymentMethods P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElem = "advert_no_payment_methods"
const P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElemAdvertRemaining P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElem = "advert_remaining"
const P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElemAdvertiserAdsPaused P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElem = "advertiser_ads_paused"
const P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElemAdvertiserApproval P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElem = "advertiser_approval"
const P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElemAdvertiserBalance P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElem = "advertiser_balance"
const P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElemAdvertiserBlockTradeIneligible P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElem = "advertiser_block_trade_ineligible"
const P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElemAdvertiserDailyLimit P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElem = "advertiser_daily_limit"
const P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElemAdvertiserSchedule P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElem = "advertiser_schedule"
const P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElemAdvertiserTempBan P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElem = "advertiser_temp_ban"

var enumValues_P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElem = []interface{}{
	"advert_fixed_rate_disabled",
	"advert_float_rate_disabled",
	"advert_inactive",
	"advert_max_limit",
	"advert_min_limit",
	"advert_remaining",
	"advert_no_payment_methods",
	"advertiser_ads_paused",
	"advertiser_approval",
	"advertiser_balance",
	"advertiser_block_trade_ineligible",
	"advertiser_daily_limit",
	"advertiser_schedule",
	"advertiser_temp_ban",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElem) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElem {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElem, v)
	}
	*j = P2PAdvertListRespP2PAdvertListListElemVisibilityStatusElem(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListRespP2PAdvertListListElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["account_currency"]; raw != nil && !ok {
		return fmt.Errorf("field account_currency in P2PAdvertListRespP2PAdvertListListElem: required")
	}
	if _, ok := raw["advertiser_details"]; raw != nil && !ok {
		return fmt.Errorf("field advertiser_details in P2PAdvertListRespP2PAdvertListListElem: required")
	}
	if _, ok := raw["block_trade"]; raw != nil && !ok {
		return fmt.Errorf("field block_trade in P2PAdvertListRespP2PAdvertListListElem: required")
	}
	if _, ok := raw["counterparty_type"]; raw != nil && !ok {
		return fmt.Errorf("field counterparty_type in P2PAdvertListRespP2PAdvertListListElem: required")
	}
	if _, ok := raw["country"]; raw != nil && !ok {
		return fmt.Errorf("field country in P2PAdvertListRespP2PAdvertListListElem: required")
	}
	if _, ok := raw["created_time"]; raw != nil && !ok {
		return fmt.Errorf("field created_time in P2PAdvertListRespP2PAdvertListListElem: required")
	}
	if _, ok := raw["description"]; raw != nil && !ok {
		return fmt.Errorf("field description in P2PAdvertListRespP2PAdvertListListElem: required")
	}
	if _, ok := raw["effective_rate"]; raw != nil && !ok {
		return fmt.Errorf("field effective_rate in P2PAdvertListRespP2PAdvertListListElem: required")
	}
	if _, ok := raw["effective_rate_display"]; raw != nil && !ok {
		return fmt.Errorf("field effective_rate_display in P2PAdvertListRespP2PAdvertListListElem: required")
	}
	if _, ok := raw["id"]; raw != nil && !ok {
		return fmt.Errorf("field id in P2PAdvertListRespP2PAdvertListListElem: required")
	}
	if _, ok := raw["is_active"]; raw != nil && !ok {
		return fmt.Errorf("field is_active in P2PAdvertListRespP2PAdvertListListElem: required")
	}
	if _, ok := raw["local_currency"]; raw != nil && !ok {
		return fmt.Errorf("field local_currency in P2PAdvertListRespP2PAdvertListListElem: required")
	}
	if _, ok := raw["max_order_amount_limit"]; raw != nil && !ok {
		return fmt.Errorf("field max_order_amount_limit in P2PAdvertListRespP2PAdvertListListElem: required")
	}
	if _, ok := raw["max_order_amount_limit_display"]; raw != nil && !ok {
		return fmt.Errorf("field max_order_amount_limit_display in P2PAdvertListRespP2PAdvertListListElem: required")
	}
	if _, ok := raw["min_order_amount_limit"]; raw != nil && !ok {
		return fmt.Errorf("field min_order_amount_limit in P2PAdvertListRespP2PAdvertListListElem: required")
	}
	if _, ok := raw["min_order_amount_limit_display"]; raw != nil && !ok {
		return fmt.Errorf("field min_order_amount_limit_display in P2PAdvertListRespP2PAdvertListListElem: required")
	}
	if _, ok := raw["order_expiry_period"]; raw != nil && !ok {
		return fmt.Errorf("field order_expiry_period in P2PAdvertListRespP2PAdvertListListElem: required")
	}
	if _, ok := raw["payment_method"]; raw != nil && !ok {
		return fmt.Errorf("field payment_method in P2PAdvertListRespP2PAdvertListListElem: required")
	}
	if _, ok := raw["price"]; raw != nil && !ok {
		return fmt.Errorf("field price in P2PAdvertListRespP2PAdvertListListElem: required")
	}
	if _, ok := raw["price_display"]; raw != nil && !ok {
		return fmt.Errorf("field price_display in P2PAdvertListRespP2PAdvertListListElem: required")
	}
	if _, ok := raw["rate"]; raw != nil && !ok {
		return fmt.Errorf("field rate in P2PAdvertListRespP2PAdvertListListElem: required")
	}
	if _, ok := raw["rate_display"]; raw != nil && !ok {
		return fmt.Errorf("field rate_display in P2PAdvertListRespP2PAdvertListListElem: required")
	}
	if _, ok := raw["rate_type"]; raw != nil && !ok {
		return fmt.Errorf("field rate_type in P2PAdvertListRespP2PAdvertListListElem: required")
	}
	if _, ok := raw["type"]; raw != nil && !ok {
		return fmt.Errorf("field type in P2PAdvertListRespP2PAdvertListListElem: required")
	}
	type Plain P2PAdvertListRespP2PAdvertListListElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["is_eligible"]; !ok || v == nil {
		plain.IsEligible = 0.0
	}
	if v, ok := raw["is_visible"]; !ok || v == nil {
		plain.IsVisible = 0.0
	}
	*j = P2PAdvertListRespP2PAdvertListListElem(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListRespP2PAdvertList) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["list"]; raw != nil && !ok {
		return fmt.Errorf("field list in P2PAdvertListRespP2PAdvertList: required")
	}
	type Plain P2PAdvertListRespP2PAdvertList
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = P2PAdvertListRespP2PAdvertList(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListResp) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["echo_req"]; raw != nil && !ok {
		return fmt.Errorf("field echo_req in P2PAdvertListResp: required")
	}
	if _, ok := raw["msg_type"]; raw != nil && !ok {
		return fmt.Errorf("field msg_type in P2PAdvertListResp: required")
	}
	type Plain P2PAdvertListResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = P2PAdvertListResp(plain)
	return nil
}
