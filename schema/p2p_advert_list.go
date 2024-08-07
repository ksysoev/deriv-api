// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// Returns available adverts for use with `p2p_order_create` .
type P2PAdvertList struct {
	// [Optional] ID of the advertiser to list adverts for.
	AdvertiserId *string `json:"advertiser_id,omitempty"`

	// [Optional] Search for advertiser by name. Partial matches will be returned.
	AdvertiserName *string `json:"advertiser_name,omitempty"`

	// [Optional] How much to buy or sell, used to calculate prices.
	Amount *float64 `json:"amount,omitempty"`

	// [Optional] Return block trade adverts when 1, non-block trade adverts when 0
	// (default).
	BlockTrade P2PAdvertListBlockTrade `json:"block_trade,omitempty"`

	// [Optional] Filter the adverts by `counterparty_type`.
	CounterpartyType *P2PAdvertListCounterpartyType `json:"counterparty_type,omitempty"`

	// [Optional] Only show adverts from favourite advertisers. Default is 0.
	FavouritesOnly *P2PAdvertListFavouritesOnly `json:"favourites_only,omitempty"`

	// [Optional] If set to 1, adverts for which the current user's shcedule does not
	// have availability from now until the full possible order expiry are not
	// returned.
	HideClientScheduleUnavailable P2PAdvertListHideClientScheduleUnavailable `json:"hide_client_schedule_unavailable,omitempty"`

	// [Optional] If set to 1, adverts for which the current user does not meet
	// counteryparty terms are not returned.
	HideIneligible P2PAdvertListHideIneligible `json:"hide_ineligible,omitempty"`

	// [Optional] Used for paging.
	Limit int `json:"limit,omitempty"`

	// [Optional] Currency to conduct payment transaction in. If not provided, only
	// ads from country of residence will be returned.
	LocalCurrency *string `json:"local_currency,omitempty"`

	// [Optional] The login id of the user. Mandatory when multiple tokens were
	// provided during authorize.
	Loginid *string `json:"loginid,omitempty"`

	// [Optional] Used for paging.
	Offset int `json:"offset,omitempty"`

	// Must be 1
	P2PAdvertList P2PAdvertListP2PAdvertList `json:"p2p_advert_list"`

	// [Optional] Used to pass data through the websocket, which may be retrieved via
	// the `echo_req` output field.
	Passthrough P2PAdvertListPassthrough `json:"passthrough,omitempty"`

	// [Optional] Search by supported payment methods.
	PaymentMethod []string `json:"payment_method,omitempty"`

	// [Optional] Used to map request to response.
	ReqId *int `json:"req_id,omitempty"`

	// [Optional] How the results are sorted.
	SortBy P2PAdvertListSortBy `json:"sort_by,omitempty"`

	// [Optional] If set to 1, ads that exceed this account's balance or turnover
	// limits will not be shown.
	UseClientLimits P2PAdvertListUseClientLimits `json:"use_client_limits,omitempty"`
}

type P2PAdvertListBlockTrade int

var enumValues_P2PAdvertListBlockTrade = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListBlockTrade) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertListBlockTrade {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertListBlockTrade, v)
	}
	*j = P2PAdvertListBlockTrade(v)
	return nil
}

type P2PAdvertListCounterpartyType string

const P2PAdvertListCounterpartyTypeBuy P2PAdvertListCounterpartyType = "buy"
const P2PAdvertListCounterpartyTypeSell P2PAdvertListCounterpartyType = "sell"

var enumValues_P2PAdvertListCounterpartyType = []interface{}{
	"buy",
	"sell",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListCounterpartyType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertListCounterpartyType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertListCounterpartyType, v)
	}
	*j = P2PAdvertListCounterpartyType(v)
	return nil
}

type P2PAdvertListFavouritesOnly int

var enumValues_P2PAdvertListFavouritesOnly = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListFavouritesOnly) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertListFavouritesOnly {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertListFavouritesOnly, v)
	}
	*j = P2PAdvertListFavouritesOnly(v)
	return nil
}

type P2PAdvertListHideClientScheduleUnavailable int

var enumValues_P2PAdvertListHideClientScheduleUnavailable = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListHideClientScheduleUnavailable) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertListHideClientScheduleUnavailable {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertListHideClientScheduleUnavailable, v)
	}
	*j = P2PAdvertListHideClientScheduleUnavailable(v)
	return nil
}

type P2PAdvertListHideIneligible int

var enumValues_P2PAdvertListHideIneligible = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListHideIneligible) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertListHideIneligible {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertListHideIneligible, v)
	}
	*j = P2PAdvertListHideIneligible(v)
	return nil
}

type P2PAdvertListP2PAdvertList int

var enumValues_P2PAdvertListP2PAdvertList = []interface{}{
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListP2PAdvertList) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertListP2PAdvertList {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertListP2PAdvertList, v)
	}
	*j = P2PAdvertListP2PAdvertList(v)
	return nil
}

// [Optional] Used to pass data through the websocket, which may be retrieved via
// the `echo_req` output field.
type P2PAdvertListPassthrough map[string]interface{}

type P2PAdvertListSortBy string

const P2PAdvertListSortByCompletion P2PAdvertListSortBy = "completion"
const P2PAdvertListSortByRate P2PAdvertListSortBy = "rate"
const P2PAdvertListSortByRating P2PAdvertListSortBy = "rating"
const P2PAdvertListSortByRecommended P2PAdvertListSortBy = "recommended"

var enumValues_P2PAdvertListSortBy = []interface{}{
	"completion",
	"rate",
	"rating",
	"recommended",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListSortBy) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertListSortBy {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertListSortBy, v)
	}
	*j = P2PAdvertListSortBy(v)
	return nil
}

type P2PAdvertListUseClientLimits int

var enumValues_P2PAdvertListUseClientLimits = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertListUseClientLimits) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertListUseClientLimits {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertListUseClientLimits, v)
	}
	*j = P2PAdvertListUseClientLimits(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertList) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["p2p_advert_list"]; raw != nil && !ok {
		return fmt.Errorf("field p2p_advert_list in P2PAdvertList: required")
	}
	type Plain P2PAdvertList
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["block_trade"]; !ok || v == nil {
		plain.BlockTrade = 0.0
	}
	if v, ok := raw["hide_client_schedule_unavailable"]; !ok || v == nil {
		plain.HideClientScheduleUnavailable = 0.0
	}
	if v, ok := raw["hide_ineligible"]; !ok || v == nil {
		plain.HideIneligible = 0.0
	}
	if v, ok := raw["limit"]; !ok || v == nil {
		plain.Limit = 50.0
	}
	if v, ok := raw["offset"]; !ok || v == nil {
		plain.Offset = 0.0
	}
	if v, ok := raw["sort_by"]; !ok || v == nil {
		plain.SortBy = "rate"
	}
	if v, ok := raw["use_client_limits"]; !ok || v == nil {
		plain.UseClientLimits = 0.0
	}
	*j = P2PAdvertList(plain)
	return nil
}
