// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "fmt"
import "reflect"
import "encoding/json"

// Echo of the request made.
type P2PAdvertUpdateRespEchoReq map[string]interface{}

type P2PAdvertUpdateRespMsgType string

var enumValues_P2PAdvertUpdateRespMsgType = []interface{}{
	"p2p_advert_update",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertUpdateRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertUpdateRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertUpdateRespMsgType, v)
	}
	*j = P2PAdvertUpdateRespMsgType(v)
	return nil
}

const P2PAdvertUpdateRespMsgTypeP2PAdvertUpdate P2PAdvertUpdateRespMsgType = "p2p_advert_update"

type P2PAdvertUpdateRespP2PAdvertUpdateAdvertiserDetailsIsOnline int

var enumValues_P2PAdvertUpdateRespP2PAdvertUpdateAdvertiserDetailsIsOnline = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertUpdateRespP2PAdvertUpdateAdvertiserDetailsIsOnline) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertUpdateRespP2PAdvertUpdateAdvertiserDetailsIsOnline {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertUpdateRespP2PAdvertUpdateAdvertiserDetailsIsOnline, v)
	}
	*j = P2PAdvertUpdateRespP2PAdvertUpdateAdvertiserDetailsIsOnline(v)
	return nil
}

// Details of the advertiser for this advert.
type P2PAdvertUpdateRespP2PAdvertUpdateAdvertiserDetails struct {
	// The total number of orders completed in the past 30 days.
	CompletedOrdersCount int `json:"completed_orders_count"`

	// The advertiser's first name.
	FirstName *string `json:"first_name,omitempty"`

	// The advertiser's unique identifier.
	Id string `json:"id"`

	// Indicates if the advertiser is currently online.
	IsOnline P2PAdvertUpdateRespP2PAdvertUpdateAdvertiserDetailsIsOnline `json:"is_online"`

	// The advertiser's last name.
	LastName *string `json:"last_name,omitempty"`

	// Epoch of the latest time the advertiser was online, up to 6 months.
	LastOnlineTime interface{} `json:"last_online_time"`

	// The advertiser's displayed name.
	Name string `json:"name"`

	// Average rating of the advertiser, range is 1-5.
	RatingAverage interface{} `json:"rating_average"`

	// Number of ratings given to the advertiser.
	RatingCount int `json:"rating_count"`

	// Percentage of users who have recommended the advertiser.
	RecommendedAverage interface{} `json:"recommended_average"`

	// Number of times the advertiser has been recommended.
	RecommendedCount interface{} `json:"recommended_count"`

	// The percentage of successfully completed orders made by or placed against the
	// advertiser within the past 30 days.
	TotalCompletionRate interface{} `json:"total_completion_rate"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertUpdateRespP2PAdvertUpdateAdvertiserDetails) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["completed_orders_count"]; !ok || v == nil {
		return fmt.Errorf("field completed_orders_count: required")
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	if v, ok := raw["is_online"]; !ok || v == nil {
		return fmt.Errorf("field is_online: required")
	}
	if v, ok := raw["last_online_time"]; !ok || v == nil {
		return fmt.Errorf("field last_online_time: required")
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	if v, ok := raw["rating_average"]; !ok || v == nil {
		return fmt.Errorf("field rating_average: required")
	}
	if v, ok := raw["rating_count"]; !ok || v == nil {
		return fmt.Errorf("field rating_count: required")
	}
	if v, ok := raw["recommended_average"]; !ok || v == nil {
		return fmt.Errorf("field recommended_average: required")
	}
	if v, ok := raw["recommended_count"]; !ok || v == nil {
		return fmt.Errorf("field recommended_count: required")
	}
	if v, ok := raw["total_completion_rate"]; !ok || v == nil {
		return fmt.Errorf("field total_completion_rate: required")
	}
	type Plain P2PAdvertUpdateRespP2PAdvertUpdateAdvertiserDetails
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = P2PAdvertUpdateRespP2PAdvertUpdateAdvertiserDetails(plain)
	return nil
}

type P2PAdvertUpdateRespP2PAdvertUpdateBlockTrade int

var enumValues_P2PAdvertUpdateRespP2PAdvertUpdateBlockTrade = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertUpdateRespP2PAdvertUpdateBlockTrade) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertUpdateRespP2PAdvertUpdateBlockTrade {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertUpdateRespP2PAdvertUpdateBlockTrade, v)
	}
	*j = P2PAdvertUpdateRespP2PAdvertUpdateBlockTrade(v)
	return nil
}

type P2PAdvertUpdateRespP2PAdvertUpdateCounterpartyType string

var enumValues_P2PAdvertUpdateRespP2PAdvertUpdateCounterpartyType = []interface{}{
	"buy",
	"sell",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertUpdateRespP2PAdvertUpdateCounterpartyType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertUpdateRespP2PAdvertUpdateCounterpartyType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertUpdateRespP2PAdvertUpdateCounterpartyType, v)
	}
	*j = P2PAdvertUpdateRespP2PAdvertUpdateCounterpartyType(v)
	return nil
}

const P2PAdvertUpdateRespP2PAdvertUpdateCounterpartyTypeBuy P2PAdvertUpdateRespP2PAdvertUpdateCounterpartyType = "buy"
const P2PAdvertUpdateRespP2PAdvertUpdateCounterpartyTypeSell P2PAdvertUpdateRespP2PAdvertUpdateCounterpartyType = "sell"

type P2PAdvertUpdateRespP2PAdvertUpdateDeleted int

var enumValues_P2PAdvertUpdateRespP2PAdvertUpdateDeleted = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertUpdateRespP2PAdvertUpdateDeleted) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertUpdateRespP2PAdvertUpdateDeleted {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertUpdateRespP2PAdvertUpdateDeleted, v)
	}
	*j = P2PAdvertUpdateRespP2PAdvertUpdateDeleted(v)
	return nil
}

type P2PAdvertUpdateRespP2PAdvertUpdateIsActive int

var enumValues_P2PAdvertUpdateRespP2PAdvertUpdateIsActive = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertUpdateRespP2PAdvertUpdateIsActive) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertUpdateRespP2PAdvertUpdateIsActive {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertUpdateRespP2PAdvertUpdateIsActive, v)
	}
	*j = P2PAdvertUpdateRespP2PAdvertUpdateIsActive(v)
	return nil
}

type P2PAdvertUpdateRespP2PAdvertUpdateIsVisible int

var enumValues_P2PAdvertUpdateRespP2PAdvertUpdateIsVisible = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertUpdateRespP2PAdvertUpdateIsVisible) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertUpdateRespP2PAdvertUpdateIsVisible {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertUpdateRespP2PAdvertUpdateIsVisible, v)
	}
	*j = P2PAdvertUpdateRespP2PAdvertUpdateIsVisible(v)
	return nil
}

// Details of available payment methods (sell adverts only).
type P2PAdvertUpdateRespP2PAdvertUpdatePaymentMethodDetails map[string]interface{}

type P2PAdvertUpdateRespP2PAdvertUpdateRateType string

var enumValues_P2PAdvertUpdateRespP2PAdvertUpdateRateType = []interface{}{
	"fixed",
	"float",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertUpdateRespP2PAdvertUpdateRateType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertUpdateRespP2PAdvertUpdateRateType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertUpdateRespP2PAdvertUpdateRateType, v)
	}
	*j = P2PAdvertUpdateRespP2PAdvertUpdateRateType(v)
	return nil
}

const P2PAdvertUpdateRespP2PAdvertUpdateRateTypeFixed P2PAdvertUpdateRespP2PAdvertUpdateRateType = "fixed"
const P2PAdvertUpdateRespP2PAdvertUpdateRateTypeFloat P2PAdvertUpdateRespP2PAdvertUpdateRateType = "float"

type P2PAdvertUpdateRespP2PAdvertUpdateType string

var enumValues_P2PAdvertUpdateRespP2PAdvertUpdateType = []interface{}{
	"buy",
	"sell",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertUpdateRespP2PAdvertUpdateType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertUpdateRespP2PAdvertUpdateType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertUpdateRespP2PAdvertUpdateType, v)
	}
	*j = P2PAdvertUpdateRespP2PAdvertUpdateType(v)
	return nil
}

const P2PAdvertUpdateRespP2PAdvertUpdateTypeBuy P2PAdvertUpdateRespP2PAdvertUpdateType = "buy"
const P2PAdvertUpdateRespP2PAdvertUpdateTypeSell P2PAdvertUpdateRespP2PAdvertUpdateType = "sell"

type P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElem string

var enumValues_P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElem = []interface{}{
	"advert_inactive",
	"advert_max_limit",
	"advert_min_limit",
	"advert_remaining",
	"advertiser_ads_paused",
	"advertiser_approval",
	"advertiser_balance",
	"advertiser_block_trade_ineligible",
	"advertiser_daily_limit",
	"advertiser_temp_ban",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElem) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElem {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElem, v)
	}
	*j = P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElem(v)
	return nil
}

// P2P updated advert information.
type P2PAdvertUpdateRespP2PAdvertUpdate struct {
	// Currency for this advert. This is the system currency to be transferred between
	// advertiser and client.
	AccountCurrency *string `json:"account_currency,omitempty"`

	// The number of active orders against this advert.
	ActiveOrders *int `json:"active_orders,omitempty"`

	// Details of the advertiser for this advert.
	AdvertiserDetails *P2PAdvertUpdateRespP2PAdvertUpdateAdvertiserDetails `json:"advertiser_details,omitempty"`

	// The total amount specified in advert, in `account_currency`.
	Amount *float64 `json:"amount,omitempty"`

	// The total amount specified in advert, in `account_currency`, formatted to
	// appropriate decimal places.
	AmountDisplay *string `json:"amount_display,omitempty"`

	// Indicates if this is block trade advert or not.
	BlockTrade *P2PAdvertUpdateRespP2PAdvertUpdateBlockTrade `json:"block_trade,omitempty"`

	// Advertiser contact information. Only applicable for 'sell adverts'.
	ContactInfo *string `json:"contact_info,omitempty"`

	// Type of transaction from the opposite party's perspective.
	CounterpartyType *P2PAdvertUpdateRespP2PAdvertUpdateCounterpartyType `json:"counterparty_type,omitempty"`

	// The target country code of the advert.
	Country *string `json:"country,omitempty"`

	// The advert creation time in epoch.
	CreatedTime *int `json:"created_time,omitempty"`

	// Days until automatic inactivation of this ad, if no activity occurs.
	DaysUntilArchive *int `json:"days_until_archive,omitempty"`

	// Indicates that the advert has been deleted.
	Deleted *P2PAdvertUpdateRespP2PAdvertUpdateDeleted `json:"deleted,omitempty"`

	// General information about the advert.
	Description *string `json:"description,omitempty"`

	// Conversion rate from account currency to local currency, using current market
	// rate if applicable.
	EffectiveRate interface{} `json:"effective_rate,omitempty"`

	// Conversion rate from account currency to local currency, using current market
	// rate if applicable, formatted to appropriate decimal places.
	EffectiveRateDisplay interface{} `json:"effective_rate_display,omitempty"`

	// The unique identifier for this advert.
	Id string `json:"id"`

	// The activation status of the advert.
	IsActive *P2PAdvertUpdateRespP2PAdvertUpdateIsActive `json:"is_active,omitempty"`

	// Indicates that this advert will appear on the main advert list.
	IsVisible P2PAdvertUpdateRespP2PAdvertUpdateIsVisible `json:"is_visible,omitempty"`

	// Local currency for this advert. This is the form of payment to be arranged
	// directly between advertiser and client.
	LocalCurrency *string `json:"local_currency,omitempty"`

	// Maximum order amount specified in advert, in `account_currency`.
	MaxOrderAmount *float64 `json:"max_order_amount,omitempty"`

	// Maximum order amount specified in advert, in `account_currency`, formatted to
	// appropriate decimal places.
	MaxOrderAmountDisplay *string `json:"max_order_amount_display,omitempty"`

	// Maximum order amount at this time, in `account_currency`.
	MaxOrderAmountLimit *float64 `json:"max_order_amount_limit,omitempty"`

	// Maximum order amount at this time, in `account_currency`, formatted to
	// appropriate decimal places.
	MaxOrderAmountLimitDisplay *string `json:"max_order_amount_limit_display,omitempty"`

	// Minimum order amount specified in advert, in `account_currency`. It is only
	// visible to the advert owner.
	MinOrderAmount *float64 `json:"min_order_amount,omitempty"`

	// Minimum order amount specified in advert, in `account_currency`, formatted to
	// appropriate decimal places.
	MinOrderAmountDisplay *string `json:"min_order_amount_display,omitempty"`

	// Minimum order amount at this time, in `account_currency`.
	MinOrderAmountLimit *float64 `json:"min_order_amount_limit,omitempty"`

	// Minimum order amount at this time, in `account_currency`, formatted to
	// appropriate decimal places.
	MinOrderAmountLimitDisplay *string `json:"min_order_amount_limit_display,omitempty"`

	// Payment instructions. Only applicable for 'sell adverts'.
	PaymentInfo *string `json:"payment_info,omitempty"`

	// Payment method name (deprecated).
	PaymentMethod interface{} `json:"payment_method,omitempty"`

	// Details of available payment methods (sell adverts only).
	PaymentMethodDetails P2PAdvertUpdateRespP2PAdvertUpdatePaymentMethodDetails `json:"payment_method_details,omitempty"`

	// Names of supported payment methods.
	PaymentMethodNames []string `json:"payment_method_names,omitempty"`

	// Cost of the advert in local currency.
	Price interface{} `json:"price,omitempty"`

	// Cost of the advert in local currency, formatted to appropriate decimal places.
	PriceDisplay interface{} `json:"price_display,omitempty"`

	// Conversion rate from advertiser's account currency to `local_currency`. An
	// absolute rate value (fixed), or percentage offset from current market rate
	// (floating).
	Rate *float64 `json:"rate,omitempty"`

	// Conversion rate formatted to appropriate decimal places.
	RateDisplay *string `json:"rate_display,omitempty"`

	// Type of rate, fixed or floating.
	RateType *P2PAdvertUpdateRespP2PAdvertUpdateRateType `json:"rate_type,omitempty"`

	// Amount currently available for orders, in `account_currency`.
	RemainingAmount *float64 `json:"remaining_amount,omitempty"`

	// Amount currently available for orders, in `account_currency`, formatted to
	// appropriate decimal places.
	RemainingAmountDisplay *string `json:"remaining_amount_display,omitempty"`

	// Whether this is a buy or a sell.
	Type *P2PAdvertUpdateRespP2PAdvertUpdateType `json:"type,omitempty"`

	// Reasons why an advert is not visible. Possible values:
	// - `advert_inactive`: the advert is set inactive.
	// - `advert_max_limit`: the minimum order amount exceeds the system maximum
	// order.
	// - `advert_min_limit`: the maximum order amount is too small to be shown on the
	// advert list.
	// - `advert_remaining`: the remaining amount of the advert is below the minimum
	// order.
	// - `advertiser_ads_paused`: the advertiser has paused all adverts.
	// - `advertiser_approval`: the advertiser's proof of identity is not verified.
	// - `advertiser_balance`: the advertiser's P2P balance is less than the minimum
	// order.
	// - `advertiser_block_trade_ineligible`: the advertiser is not currently eligible
	// for block trading.
	// - `advertiser_daily_limit`: the advertiser's remaining daily limit is less than
	// the minimum order.
	// - `advertiser_temp_ban`: the advertiser is temporarily banned from P2P.
	VisibilityStatus []P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElem `json:"visibility_status,omitempty"`
}

const P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElemAdvertInactive P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElem = "advert_inactive"
const P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElemAdvertMaxLimit P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElem = "advert_max_limit"
const P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElemAdvertMinLimit P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElem = "advert_min_limit"
const P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElemAdvertRemaining P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElem = "advert_remaining"
const P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElemAdvertiserAdsPaused P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElem = "advertiser_ads_paused"
const P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElemAdvertiserApproval P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElem = "advertiser_approval"
const P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElemAdvertiserBalance P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElem = "advertiser_balance"
const P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElemAdvertiserBlockTradeIneligible P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElem = "advertiser_block_trade_ineligible"
const P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElemAdvertiserDailyLimit P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElem = "advertiser_daily_limit"
const P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElemAdvertiserTempBan P2PAdvertUpdateRespP2PAdvertUpdateVisibilityStatusElem = "advertiser_temp_ban"

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertUpdateRespP2PAdvertUpdate) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain P2PAdvertUpdateRespP2PAdvertUpdate
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["is_visible"]; !ok || v == nil {
		plain.IsVisible = 0
	}
	*j = P2PAdvertUpdateRespP2PAdvertUpdate(plain)
	return nil
}

// Returns information about the updated advert.
type P2PAdvertUpdateResp struct {
	// Echo of the request made.
	EchoReq P2PAdvertUpdateRespEchoReq `json:"echo_req"`

	// Action name of the request made.
	MsgType P2PAdvertUpdateRespMsgType `json:"msg_type"`

	// P2P updated advert information.
	P2PAdvertUpdate *P2PAdvertUpdateRespP2PAdvertUpdate `json:"p2p_advert_update,omitempty"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertUpdateResp) UnmarshalJSON(b []byte) error {
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
	type Plain P2PAdvertUpdateResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = P2PAdvertUpdateResp(plain)
	return nil
}