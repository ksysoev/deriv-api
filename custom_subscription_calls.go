package deriv

import "github.com/ksysoev/deriv-api/schema"

// SubscribeTicksHistory Get historic tick data for a given symbol.
//
//nolint:gocritic // don't want to break backward compatibility for now
func (api *DerivAPI) SubscribeTicksHistory(r schema.TicksHistory) (rsp schema.TicksHistoryResp, s *Subsciption[schema.TicksHistoryResp, schema.TicksResp], err error) {
	var f schema.TicksHistorySubscribe = 1

	id := api.getNextRequestID()
	r.ReqId = &id
	r.Subscribe = &f
	r.Style = schema.TicksHistoryStyleTicks
	s = NewSubcription[schema.TicksHistoryResp, schema.TicksResp](api)
	rsp, err = s.Start(id, r)

	return
}

// SubscribeTicksHistory Get historic candles data for a given symbol.
//
//nolint:gocritic // don't want to break backward compatibility for now
func (api *DerivAPI) SubscribeCandlesHistory(r schema.TicksHistory) (rsp schema.TicksHistoryResp, s *Subsciption[schema.TicksHistoryResp, schema.TicksHistoryResp], err error) {
	var f schema.TicksHistorySubscribe = 1

	id := api.getNextRequestID()
	r.ReqId = &id
	r.Subscribe = &f
	r.Style = schema.TicksHistoryStyleCandles
	s = NewSubcription[schema.TicksHistoryResp, schema.TicksHistoryResp](api)
	rsp, err = s.Start(id, r)

	return
}

// SubscribeTransaction Subscribe to transaction notifications
func (api *DerivAPI) SubscribeTransaction(r schema.Transaction) (rsp schema.TransactionResp, s *Subsciption[schema.TransactionResp, schema.TransactionResp], err error) {
	var f schema.TransactionSubscribe = 1

	id := api.getNextRequestID()
	r.ReqId = &id
	r.Subscribe = f
	s = NewSubcription[schema.TransactionResp, schema.TransactionResp](api)
	rsp, err = s.Start(id, r)

	return
}
