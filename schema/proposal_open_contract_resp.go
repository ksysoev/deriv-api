// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// Latest price and other details for an open contract in the user's portfolio
type ProposalOpenContractResp struct {
	// Echo of the request made.
	EchoReq ProposalOpenContractRespEchoReq `json:"echo_req"`

	// Action name of the request made.
	MsgType *ProposalOpenContractRespMsgType `json:"msg_type,omitempty"`

	// Latest price and other details for an open contract
	ProposalOpenContract *ProposalOpenContractRespProposalOpenContract `json:"proposal_open_contract,omitempty"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`

	// For subscription requests only.
	Subscription *ProposalOpenContractRespSubscription `json:"subscription,omitempty"`
}

// Echo of the request made.
type ProposalOpenContractRespEchoReq map[string]interface{}

type ProposalOpenContractRespMsgType string

const ProposalOpenContractRespMsgTypeProposalOpenContract ProposalOpenContractRespMsgType = "proposal_open_contract"

var enumValues_ProposalOpenContractRespMsgType = []interface{}{
	"proposal_open_contract",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ProposalOpenContractRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ProposalOpenContractRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ProposalOpenContractRespMsgType, v)
	}
	*j = ProposalOpenContractRespMsgType(v)
	return nil
}

// Latest price and other details for an open contract
type ProposalOpenContractRespProposalOpenContract struct {
	// Account Id
	AccountId *float64 `json:"account_id,omitempty"`

	// The markup amount charged on a client's stake amount
	AppMarkupAmount *string `json:"app_markup_amount,omitempty"`

	// Tick details around contract start and end time.
	AuditDetails *ProposalOpenContractRespProposalOpenContractAuditDetails `json:"audit_details,omitempty"`

	// Barrier of the contract (if any).
	Barrier *string `json:"barrier,omitempty"`

	// The number of barriers a contract has.
	BarrierCount *float64 `json:"barrier_count,omitempty"`

	// [Only for accumulator] Absolute difference between high/low barrier and spot
	BarrierSpotDistance *string `json:"barrier_spot_distance,omitempty"`

	// Price at which the contract could be sold back to the company.
	BidPrice *float64 `json:"bid_price,omitempty"`

	// Price at which contract was purchased
	BuyPrice *float64 `json:"buy_price,omitempty"`

	// Contains information about contract cancellation option.
	Cancellation *ProposalOpenContractRespProposalOpenContractCancellation `json:"cancellation,omitempty"`

	// The caution price for the Snowball contract. Breaching this price will reset
	// the coupons accrued to 0.
	CautionPrice *string `json:"caution_price,omitempty"`

	// Commission in payout currency amount.
	Commision *float64 `json:"commision,omitempty"`

	// Commission in payout currency amount.
	Commission *float64 `json:"commission,omitempty"`

	// The internal contract identifier
	ContractId *int `json:"contract_id,omitempty"`

	// Contract type.
	ContractType *string `json:"contract_type,omitempty"`

	// The epoch times at which the coupons will be accrued for the Snowball contract.
	CouponCollectionEpochs []int `json:"coupon_collection_epochs,omitempty"`

	// The coupon rate for the Snowball contract at which the stake will grow for each
	// coupon accrued.
	CouponRate *string `json:"coupon_rate,omitempty"`

	// The currency code of the contract.
	Currency *string `json:"currency,omitempty"`

	// Spot value if we have license to stream this symbol.
	CurrentSpot *float64 `json:"current_spot,omitempty"`

	// Spot value with the correct precision if we have license to stream this symbol.
	CurrentSpotDisplayValue *string `json:"current_spot_display_value,omitempty"`

	// [Applicable for accumulator] High barrier based on current spot.
	CurrentSpotHighBarrier *string `json:"current_spot_high_barrier,omitempty"`

	// [Applicable for accumulator] Low barrier based on current spot.
	CurrentSpotLowBarrier *string `json:"current_spot_low_barrier,omitempty"`

	// The corresponding time of the current spot.
	CurrentSpotTime *int `json:"current_spot_time,omitempty"`

	// Expiry date (epoch) of the Contract. Please note that it is not applicable for
	// tick trade contracts.
	DateExpiry *int `json:"date_expiry,omitempty"`

	// Settlement date (epoch) of the contract.
	DateSettlement *int `json:"date_settlement,omitempty"`

	// Start date (epoch) of the contract.
	DateStart *int `json:"date_start,omitempty"`

	// Display name of underlying
	DisplayName *string `json:"display_name,omitempty"`

	// [Only for vanilla or turbos options] The implied number of contracts
	DisplayNumberOfContracts *string `json:"display_number_of_contracts,omitempty"`

	// The `bid_price` with the correct precision
	DisplayValue *string `json:"display_value,omitempty"`

	// Same as `entry_tick`. For backwards compatibility.
	EntrySpot *float64 `json:"entry_spot,omitempty"`

	// Same as `entry_tick_display_value`. For backwards compatibility.
	EntrySpotDisplayValue *string `json:"entry_spot_display_value,omitempty"`

	// This is the entry spot of the contract. For contracts starting immediately it
	// is the next tick after the start time. For forward-starting contracts it is the
	// spot at the start time.
	EntryTick *float64 `json:"entry_tick,omitempty"`

	// This is the entry spot with the correct precision of the contract. For
	// contracts starting immediately it is the next tick after the start time. For
	// forward-starting contracts it is the spot at the start time.
	EntryTickDisplayValue *string `json:"entry_tick_display_value,omitempty"`

	// This is the epoch time of the entry tick.
	EntryTickTime *int `json:"entry_tick_time,omitempty"`

	// Exit tick can refer to the latest tick at the end time, the tick that fulfils
	// the contract's winning or losing condition for path dependent contracts
	// (Touch/No Touch and Stays Between/Goes Outside) or the tick at which the
	// contract is sold before expiry.
	ExitTick *float64 `json:"exit_tick,omitempty"`

	// Exit tick can refer to the latest tick at the end time, the tick that fulfils
	// the contract's winning or losing condition for path dependent contracts
	// (Touch/No Touch and Stays Between/Goes Outside) or the tick at which the
	// contract is sold before expiry.
	ExitTickDisplayValue *string `json:"exit_tick_display_value,omitempty"`

	// This is the epoch time of the exit tick. Note that since certain instruments
	// don't tick every second, the exit tick time may be a few seconds before the end
	// time.
	ExitTickTime *int `json:"exit_tick_time,omitempty"`

	// This is the expiry time.
	ExpiryTime *int `json:"expiry_time,omitempty"`

	// [Only for accumulator] Growth rate of an accumulator contract.
	GrowthRate *float64 `json:"growth_rate,omitempty"`

	// High barrier of the contract (if any).
	HighBarrier *string `json:"high_barrier,omitempty"`

	// A per-connection unique identifier. Can be passed to the `forget` API call to
	// unsubscribe.
	Id *string `json:"id,omitempty"`

	// Whether the contract is expired or not.
	IsExpired *ProposalOpenContractRespProposalOpenContractIsExpired `json:"is_expired,omitempty"`

	// Whether the contract is forward-starting or not.
	IsForwardStarting *ProposalOpenContractRespProposalOpenContractIsForwardStarting `json:"is_forward_starting,omitempty"`

	// Whether the contract is an intraday contract.
	IsIntraday *ProposalOpenContractRespProposalOpenContractIsIntraday `json:"is_intraday,omitempty"`

	// Whether the contract expiry price will depend on the path of the market (e.g.
	// One Touch contract).
	IsPathDependent *ProposalOpenContractRespProposalOpenContractIsPathDependent `json:"is_path_dependent,omitempty"`

	// Whether the contract is settleable or not.
	IsSettleable *ProposalOpenContractRespProposalOpenContractIsSettleable `json:"is_settleable,omitempty"`

	// Whether the contract is sold or not.
	IsSold *ProposalOpenContractRespProposalOpenContractIsSold `json:"is_sold,omitempty"`

	// Whether the contract can be cancelled.
	IsValidToCancel *ProposalOpenContractRespProposalOpenContractIsValidToCancel `json:"is_valid_to_cancel,omitempty"`

	// Whether the contract can be sold back to the company.
	IsValidToSell *ProposalOpenContractRespProposalOpenContractIsValidToSell `json:"is_valid_to_sell,omitempty"`

	// [Optional] Indicator whether take profit, stop loss, and/or stop out is allowed
	// to be updated.
	IsValidToUpdate *ProposalOpenContractRespProposalOpenContractIsValidToUpdate `json:"is_valid_to_update,omitempty"`

	// Orders are applicable to `MULTUP` and `MULTDOWN` contracts only.
	LimitOrder *ProposalOpenContractRespProposalOpenContractLimitOrder `json:"limit_order,omitempty"`

	// Text description of the contract purchased, Example: Win payout if Volatility
	// 100 Index is strictly higher than entry spot at 10 minutes after contract start
	// time.
	Longcode *string `json:"longcode,omitempty"`

	// Low barrier of the contract (if any).
	LowBarrier *string `json:"low_barrier,omitempty"`

	// [Only for lookback trades] Multiplier applies when calculating the final payoff
	// for each type of lookback. e.g. (Exit spot - Lowest historical price) *
	// multiplier = Payout
	Multiplier *float64 `json:"multiplier,omitempty"`

	// The maximum number of coupons available for the Snowball contract.
	NumOfCoupons *int `json:"num_of_coupons,omitempty"`

	// Payout value of the contract.
	Payout *float64 `json:"payout,omitempty"`

	// The latest bid price minus buy price.
	Profit *float64 `json:"profit,omitempty"`

	// Profit in percentage.
	ProfitPercentage *float64 `json:"profit_percentage,omitempty"`

	// The profit price for the Snowball contract. Breaching this price will close the
	// contract immediately.
	ProfitPrice *string `json:"profit_price,omitempty"`

	// Epoch of purchase time, will be same as `date_start` for all contracts except
	// forward starting contracts.
	PurchaseTime *int `json:"purchase_time,omitempty"`

	// [Only for reset trades i.e. RESETCALL and RESETPUT] Reset barrier of the
	// contract.
	ResetBarrier *string `json:"reset_barrier,omitempty"`

	// [Only for reset trades i.e. RESETCALL and RESETPUT] The epoch time of a barrier
	// reset.
	ResetTime *int `json:"reset_time,omitempty"`

	// Spot value at the selected tick for the contract.
	SelectedSpot *float64 `json:"selected_spot,omitempty"`

	// [Only for highlowticks trades i.e. TICKHIGH and TICKLOW] Selected tick for the
	// contract.
	SelectedTick *int `json:"selected_tick,omitempty"`

	// Price at which contract was sold, only available when contract has been sold.
	SellPrice *float64 `json:"sell_price,omitempty"`

	// Latest spot value at the sell time. (only present for contracts already sold).
	// Will no longer be supported in the next API release.
	SellSpot *float64 `json:"sell_spot,omitempty"`

	// Latest spot value with the correct precision at the sell time. (only present
	// for contracts already sold). Will no longer be supported in the next API
	// release.
	SellSpotDisplayValue *string `json:"sell_spot_display_value,omitempty"`

	// Epoch time of the sell spot. Note that since certain underlyings don't tick
	// every second, the sell spot time may be a few seconds before the sell time.
	// (only present for contracts already sold). Will no longer be supported in the
	// next API release.
	SellSpotTime *int `json:"sell_spot_time,omitempty"`

	// Epoch time of when the contract was sold (only present for contracts already
	// sold)
	SellTime *int `json:"sell_time,omitempty"`

	// Coded description of the contract purchased.
	Shortcode *string `json:"shortcode,omitempty"`

	// Contract status. Will be `sold` if the contract was sold back before expiry,
	// `won` if won and `lost` if lost at expiry. Otherwise will be `open`
	Status *ProposalOpenContractRespProposalOpenContractStatus `json:"status,omitempty"`

	// Only for tick trades, number of ticks
	TickCount *int `json:"tick_count,omitempty"`

	// [Only for accumulator] Number of ticks passed since entry_tick
	TickPassed *int `json:"tick_passed,omitempty"`

	// Tick stream from entry to end time.
	TickStream []ProposalOpenContractRespProposalOpenContractTickStreamElem `json:"tick_stream,omitempty"`

	// The trade risk profile for the Snowball contract.
	TradeRiskProfile *ProposalOpenContractRespProposalOpenContractTradeRiskProfile `json:"trade_risk_profile,omitempty"`

	// Every contract has buy and sell transaction ids, i.e. when you purchase a
	// contract we associate it with buy transaction id, and if contract is already
	// sold we associate that with sell transaction id.
	TransactionIds *ProposalOpenContractRespProposalOpenContractTransactionIds `json:"transaction_ids,omitempty"`

	// The underlying symbol code.
	Underlying *string `json:"underlying,omitempty"`

	// Error message if validation fails
	ValidationError *string `json:"validation_error,omitempty"`

	// Error code if validation fails
	ValidationErrorCode *string `json:"validation_error_code,omitempty"`

	// Contains contract validation information.
	ValidationParams *ProposalOpenContractRespProposalOpenContractValidationParams `json:"validation_params,omitempty"`
}

// Tick details around contract start and end time.
type ProposalOpenContractRespProposalOpenContractAuditDetails struct {
	// Ticks for tick expiry contract from start time till expiry.
	AllTicks []ProposalOpenContractRespProposalOpenContractAuditDetailsAllTicksElem `json:"all_ticks,omitempty"`

	// Ticks around contract end time.
	ContractEnd []ProposalOpenContractRespProposalOpenContractAuditDetailsContractEndElem `json:"contract_end,omitempty"`

	// Ticks around contract start time.
	ContractStart []ProposalOpenContractRespProposalOpenContractAuditDetailsContractStartElem `json:"contract_start,omitempty"`
}

type ProposalOpenContractRespProposalOpenContractAuditDetailsAllTicksElem struct {
	// Epoch time of a tick or the contract start or end time.
	Epoch *int `json:"epoch,omitempty"`

	// A flag used to highlight the record in front-end applications.
	Flag *string `json:"flag,omitempty"`

	// A short description of the data. It could be a tick or a time associated with
	// the contract.
	Name *string `json:"name,omitempty"`

	// The spot value at the given epoch.
	Tick *float64 `json:"tick,omitempty"`

	// The spot value with the correct precision at the given epoch.
	TickDisplayValue *string `json:"tick_display_value,omitempty"`
}

type ProposalOpenContractRespProposalOpenContractAuditDetailsContractEndElem struct {
	// Epoch time of a tick or the contract start or end time.
	Epoch *int `json:"epoch,omitempty"`

	// A flag used to highlight the record in front-end applications.
	Flag *string `json:"flag,omitempty"`

	// A short description of the data. It could be a tick or a time associated with
	// the contract.
	Name *string `json:"name,omitempty"`

	// The spot value at the given epoch.
	Tick *float64 `json:"tick,omitempty"`

	// The spot value with the correct precision at the given epoch.
	TickDisplayValue *string `json:"tick_display_value,omitempty"`
}

type ProposalOpenContractRespProposalOpenContractAuditDetailsContractStartElem struct {
	// Epoch time of a tick or the contract start or end time.
	Epoch *int `json:"epoch,omitempty"`

	// A flag used to highlight the record in front-end applications.
	Flag *string `json:"flag,omitempty"`

	// A short description of the data. It could be a tick or a time associated with
	// the contract.
	Name *string `json:"name,omitempty"`

	// The spot value at the given epoch.
	Tick *float64 `json:"tick,omitempty"`

	// The spot value with the correct precision at the given epoch.
	TickDisplayValue *string `json:"tick_display_value,omitempty"`
}

// Contains information about contract cancellation option.
type ProposalOpenContractRespProposalOpenContractCancellation struct {
	// Ask price of contract cancellation option.
	AskPrice *float64 `json:"ask_price,omitempty"`

	// Expiry time in epoch for contract cancellation option.
	DateExpiry *int `json:"date_expiry,omitempty"`
}

type ProposalOpenContractRespProposalOpenContractIsExpired int

var enumValues_ProposalOpenContractRespProposalOpenContractIsExpired = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ProposalOpenContractRespProposalOpenContractIsExpired) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ProposalOpenContractRespProposalOpenContractIsExpired {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ProposalOpenContractRespProposalOpenContractIsExpired, v)
	}
	*j = ProposalOpenContractRespProposalOpenContractIsExpired(v)
	return nil
}

type ProposalOpenContractRespProposalOpenContractIsForwardStarting int

var enumValues_ProposalOpenContractRespProposalOpenContractIsForwardStarting = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ProposalOpenContractRespProposalOpenContractIsForwardStarting) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ProposalOpenContractRespProposalOpenContractIsForwardStarting {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ProposalOpenContractRespProposalOpenContractIsForwardStarting, v)
	}
	*j = ProposalOpenContractRespProposalOpenContractIsForwardStarting(v)
	return nil
}

type ProposalOpenContractRespProposalOpenContractIsIntraday int

var enumValues_ProposalOpenContractRespProposalOpenContractIsIntraday = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ProposalOpenContractRespProposalOpenContractIsIntraday) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ProposalOpenContractRespProposalOpenContractIsIntraday {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ProposalOpenContractRespProposalOpenContractIsIntraday, v)
	}
	*j = ProposalOpenContractRespProposalOpenContractIsIntraday(v)
	return nil
}

type ProposalOpenContractRespProposalOpenContractIsPathDependent int

var enumValues_ProposalOpenContractRespProposalOpenContractIsPathDependent = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ProposalOpenContractRespProposalOpenContractIsPathDependent) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ProposalOpenContractRespProposalOpenContractIsPathDependent {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ProposalOpenContractRespProposalOpenContractIsPathDependent, v)
	}
	*j = ProposalOpenContractRespProposalOpenContractIsPathDependent(v)
	return nil
}

type ProposalOpenContractRespProposalOpenContractIsSettleable int

var enumValues_ProposalOpenContractRespProposalOpenContractIsSettleable = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ProposalOpenContractRespProposalOpenContractIsSettleable) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ProposalOpenContractRespProposalOpenContractIsSettleable {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ProposalOpenContractRespProposalOpenContractIsSettleable, v)
	}
	*j = ProposalOpenContractRespProposalOpenContractIsSettleable(v)
	return nil
}

type ProposalOpenContractRespProposalOpenContractIsSold int

var enumValues_ProposalOpenContractRespProposalOpenContractIsSold = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ProposalOpenContractRespProposalOpenContractIsSold) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ProposalOpenContractRespProposalOpenContractIsSold {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ProposalOpenContractRespProposalOpenContractIsSold, v)
	}
	*j = ProposalOpenContractRespProposalOpenContractIsSold(v)
	return nil
}

type ProposalOpenContractRespProposalOpenContractIsValidToCancel int

var enumValues_ProposalOpenContractRespProposalOpenContractIsValidToCancel = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ProposalOpenContractRespProposalOpenContractIsValidToCancel) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ProposalOpenContractRespProposalOpenContractIsValidToCancel {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ProposalOpenContractRespProposalOpenContractIsValidToCancel, v)
	}
	*j = ProposalOpenContractRespProposalOpenContractIsValidToCancel(v)
	return nil
}

type ProposalOpenContractRespProposalOpenContractIsValidToSell int

var enumValues_ProposalOpenContractRespProposalOpenContractIsValidToSell = []interface{}{
	0,
	1,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ProposalOpenContractRespProposalOpenContractIsValidToSell) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ProposalOpenContractRespProposalOpenContractIsValidToSell {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ProposalOpenContractRespProposalOpenContractIsValidToSell, v)
	}
	*j = ProposalOpenContractRespProposalOpenContractIsValidToSell(v)
	return nil
}

// [Optional] Indicator whether take profit, stop loss, and/or stop out is allowed
// to be updated.
type ProposalOpenContractRespProposalOpenContractIsValidToUpdate struct {
	// [Optional] 1 if stop loss is allowed to be updated and 0 if otherwise. This
	// field is undefined if stop loss functionality is not supported by the contract.
	StopLoss *ProposalOpenContractRespProposalOpenContractIsValidToUpdateStopLoss `json:"stop_loss,omitempty"`

	// [Optional] 1 if stop out is allowed to be updated and 0 if otherwise. This
	// field is undefined if stop out functionality is not supported by the contract.
	StopOut *ProposalOpenContractRespProposalOpenContractIsValidToUpdateStopOut `json:"stop_out,omitempty"`

	// [Optional] 1 if take profit is allowed to be updated and 0 if otherwise. This
	// field is undefined if take profit functionality is not supported by the
	// contract.
	TakeProfit *ProposalOpenContractRespProposalOpenContractIsValidToUpdateTakeProfit `json:"take_profit,omitempty"`
}

type ProposalOpenContractRespProposalOpenContractIsValidToUpdateStopLoss float64

var enumValues_ProposalOpenContractRespProposalOpenContractIsValidToUpdateStopLoss = []interface{}{
	0.0,
	1.0,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ProposalOpenContractRespProposalOpenContractIsValidToUpdateStopLoss) UnmarshalJSON(b []byte) error {
	var v float64
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ProposalOpenContractRespProposalOpenContractIsValidToUpdateStopLoss {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ProposalOpenContractRespProposalOpenContractIsValidToUpdateStopLoss, v)
	}
	*j = ProposalOpenContractRespProposalOpenContractIsValidToUpdateStopLoss(v)
	return nil
}

type ProposalOpenContractRespProposalOpenContractIsValidToUpdateStopOut float64

var enumValues_ProposalOpenContractRespProposalOpenContractIsValidToUpdateStopOut = []interface{}{
	0.0,
	1.0,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ProposalOpenContractRespProposalOpenContractIsValidToUpdateStopOut) UnmarshalJSON(b []byte) error {
	var v float64
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ProposalOpenContractRespProposalOpenContractIsValidToUpdateStopOut {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ProposalOpenContractRespProposalOpenContractIsValidToUpdateStopOut, v)
	}
	*j = ProposalOpenContractRespProposalOpenContractIsValidToUpdateStopOut(v)
	return nil
}

type ProposalOpenContractRespProposalOpenContractIsValidToUpdateTakeProfit float64

var enumValues_ProposalOpenContractRespProposalOpenContractIsValidToUpdateTakeProfit = []interface{}{
	0.0,
	1.0,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ProposalOpenContractRespProposalOpenContractIsValidToUpdateTakeProfit) UnmarshalJSON(b []byte) error {
	var v float64
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ProposalOpenContractRespProposalOpenContractIsValidToUpdateTakeProfit {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ProposalOpenContractRespProposalOpenContractIsValidToUpdateTakeProfit, v)
	}
	*j = ProposalOpenContractRespProposalOpenContractIsValidToUpdateTakeProfit(v)
	return nil
}

// Orders are applicable to `MULTUP` and `MULTDOWN` contracts only.
type ProposalOpenContractRespProposalOpenContractLimitOrder struct {
	// Contains information where the contract will be closed automatically at the
	// loss specified by the user.
	StopLoss *ProposalOpenContractRespProposalOpenContractLimitOrderStopLoss `json:"stop_loss,omitempty"`

	// Contains information where the contract will be closed automatically when the
	// value of the contract is close to zero. This is set by the us.
	StopOut *ProposalOpenContractRespProposalOpenContractLimitOrderStopOut `json:"stop_out,omitempty"`

	// Contain information where the contract will be closed automatically at the
	// profit specified by the user.
	TakeProfit *ProposalOpenContractRespProposalOpenContractLimitOrderTakeProfit `json:"take_profit,omitempty"`
}

// Contains information where the contract will be closed automatically at the loss
// specified by the user.
type ProposalOpenContractRespProposalOpenContractLimitOrderStopLoss struct {
	// Localized display name
	DisplayName *string `json:"display_name,omitempty"`

	// Stop loss amount for display purpose.
	DisplayOrderAmount *string `json:"display_order_amount,omitempty"`

	// Stop loss amount. Will be deprecated soon. Please use [display_order_amount].
	OrderAmount *float64 `json:"order_amount,omitempty"`

	// Stop loss order epoch
	OrderDate *int `json:"order_date,omitempty"`

	// Pip-sized barrier value
	Value *string `json:"value,omitempty"`
}

// Contains information where the contract will be closed automatically when the
// value of the contract is close to zero. This is set by the us.
type ProposalOpenContractRespProposalOpenContractLimitOrderStopOut struct {
	// Localized display name
	DisplayName *string `json:"display_name,omitempty"`

	// Stop out amount for display purpose.
	DisplayOrderAmount *string `json:"display_order_amount,omitempty"`

	// Stop out amount. Will be deprecated soon. Please use [display_order_amount].
	OrderAmount *float64 `json:"order_amount,omitempty"`

	// Stop out order epoch
	OrderDate *int `json:"order_date,omitempty"`

	// Pip-sized barrier value
	Value *string `json:"value,omitempty"`
}

// Contain information where the contract will be closed automatically at the
// profit specified by the user.
type ProposalOpenContractRespProposalOpenContractLimitOrderTakeProfit struct {
	// Localized display name
	DisplayName *string `json:"display_name,omitempty"`

	// Take profit amount for display purpose.
	DisplayOrderAmount *string `json:"display_order_amount,omitempty"`

	// Take profit amount. Will be deprecated soon. Please use [display_order_amount].
	OrderAmount *float64 `json:"order_amount,omitempty"`

	// Take profit order epoch
	OrderDate *int `json:"order_date,omitempty"`

	// Pip-sized barrier value
	Value *string `json:"value,omitempty"`
}

type ProposalOpenContractRespProposalOpenContractStatus struct {
	Value interface{}
}

// MarshalJSON implements json.Marshaler.
func (j *ProposalOpenContractRespProposalOpenContractStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Value)
}

var enumValues_ProposalOpenContractRespProposalOpenContractStatus = []interface{}{
	"open",
	"sold",
	"won",
	"lost",
	"cancelled",
	nil,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ProposalOpenContractRespProposalOpenContractStatus) UnmarshalJSON(b []byte) error {
	var v struct {
		Value interface{}
	}
	if err := json.Unmarshal(b, &v.Value); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ProposalOpenContractRespProposalOpenContractStatus {
		if reflect.DeepEqual(v.Value, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ProposalOpenContractRespProposalOpenContractStatus, v.Value)
	}
	*j = ProposalOpenContractRespProposalOpenContractStatus(v)
	return nil
}

type ProposalOpenContractRespProposalOpenContractTickStreamElem struct {
	// Epoch time of a tick or the contract start or end time.
	Epoch *int `json:"epoch,omitempty"`

	// The spot value at the given epoch.
	Tick *float64 `json:"tick,omitempty"`

	// The spot value with the correct precision at the given epoch.
	TickDisplayValue *string `json:"tick_display_value,omitempty"`
}

type ProposalOpenContractRespProposalOpenContractTradeRiskProfile string

const ProposalOpenContractRespProposalOpenContractTradeRiskProfileHigh ProposalOpenContractRespProposalOpenContractTradeRiskProfile = "high"
const ProposalOpenContractRespProposalOpenContractTradeRiskProfileLow ProposalOpenContractRespProposalOpenContractTradeRiskProfile = "low"
const ProposalOpenContractRespProposalOpenContractTradeRiskProfileMedium ProposalOpenContractRespProposalOpenContractTradeRiskProfile = "medium"

var enumValues_ProposalOpenContractRespProposalOpenContractTradeRiskProfile = []interface{}{
	"low",
	"medium",
	"high",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ProposalOpenContractRespProposalOpenContractTradeRiskProfile) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ProposalOpenContractRespProposalOpenContractTradeRiskProfile {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ProposalOpenContractRespProposalOpenContractTradeRiskProfile, v)
	}
	*j = ProposalOpenContractRespProposalOpenContractTradeRiskProfile(v)
	return nil
}

// Every contract has buy and sell transaction ids, i.e. when you purchase a
// contract we associate it with buy transaction id, and if contract is already
// sold we associate that with sell transaction id.
type ProposalOpenContractRespProposalOpenContractTransactionIds struct {
	// Buy transaction ID for that contract
	Buy *int `json:"buy,omitempty"`

	// Sell transaction ID for that contract, only present when contract is already
	// sold.
	Sell *int `json:"sell,omitempty"`
}

// Contains contract validation information.
type ProposalOpenContractRespProposalOpenContractValidationParams struct {
	// [Only for Accumulators] Maximum payout for the contract.
	MaxPayout *string `json:"max_payout,omitempty"`

	// [Only for Accumulators] Maximum ticks for the contract.
	MaxTicks *int `json:"max_ticks,omitempty"`

	// Contains information for minimum and maximum stake amount for the contract.
	Stake *ProposalOpenContractRespProposalOpenContractValidationParamsStake `json:"stake,omitempty"`

	// [Only for Multipliers] Contains information for minimum and maximum stop loss
	// amount for the contract.
	StopLoss *ProposalOpenContractRespProposalOpenContractValidationParamsStopLoss `json:"stop_loss,omitempty"`

	// Contains information for minimum and maximum take profit amount for the
	// contract.
	TakeProfit *ProposalOpenContractRespProposalOpenContractValidationParamsTakeProfit `json:"take_profit,omitempty"`
}

// Contains information for minimum and maximum stake amount for the contract.
type ProposalOpenContractRespProposalOpenContractValidationParamsStake struct {
	// Maximum stakes allowed
	Max *string `json:"max,omitempty"`

	// Minimum stakes allowed
	Min *string `json:"min,omitempty"`
}

// [Only for Multipliers] Contains information for minimum and maximum stop loss
// amount for the contract.
type ProposalOpenContractRespProposalOpenContractValidationParamsStopLoss struct {
	// Maximum stop loss amount
	Max *string `json:"max,omitempty"`

	// Minimum stop loss amount
	Min *string `json:"min,omitempty"`
}

// Contains information for minimum and maximum take profit amount for the
// contract.
type ProposalOpenContractRespProposalOpenContractValidationParamsTakeProfit struct {
	// Maximum take profit amount
	Max *string `json:"max,omitempty"`

	// Minimum take profit amount
	Min *string `json:"min,omitempty"`
}

// For subscription requests only.
type ProposalOpenContractRespSubscription struct {
	// A per-connection unique identifier. Can be passed to the `forget` API call to
	// unsubscribe.
	Id string `json:"id"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ProposalOpenContractRespSubscription) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["id"]; raw != nil && !ok {
		return fmt.Errorf("field id in ProposalOpenContractRespSubscription: required")
	}
	type Plain ProposalOpenContractRespSubscription
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ProposalOpenContractRespSubscription(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ProposalOpenContractResp) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["echo_req"]; raw != nil && !ok {
		return fmt.Errorf("field echo_req in ProposalOpenContractResp: required")
	}
	type Plain ProposalOpenContractResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ProposalOpenContractResp(plain)
	return nil
}
