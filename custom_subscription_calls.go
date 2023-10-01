package deriv

import "github.com/ksysoev/deriv-api/schema"

// SubscribeTicksHistory Get historic tick data for a given symbol.
func (a *DerivAPI) SubscribeTicksHistory(r schema.TicksHistory) (rsp schema.TicksHistoryResp, s *Subsciption[schema.TicksHistoryResp, schema.TicksResp], err error) {
	id := a.getNextRequestID()
	var f schema.TicksHistorySubscribe = 1
	r.ReqId = &id
	r.Subscribe = &f
	r.Style = schema.TicksHistoryStyleTicks
	s = NewSubcription[schema.TicksHistoryResp, schema.TicksResp](a)
	rsp, err = s.Start(id, r)
	return
}

// SubscribeTicksHistory Get historic candles data for a given symbol.
func (a *DerivAPI) SubscribeCandlesHistory(r schema.TicksHistory) (rsp schema.TicksHistoryResp, s *Subsciption[schema.TicksHistoryResp, schema.TicksHistoryResp], err error) {
	id := a.getNextRequestID()
	var f schema.TicksHistorySubscribe = 1
	r.ReqId = &id
	r.Subscribe = &f
	r.Style = schema.TicksHistoryStyleCandles
	s = NewSubcription[schema.TicksHistoryResp, schema.TicksHistoryResp](a)
	rsp, err = s.Start(id, r)
	return
}

// SubscribeTransaction Subscribe to transaction notifications
func (a *DerivAPI) SubscribeTransaction(r schema.Transaction) (rsp schema.TransactionResp, s *Subsciption[schema.TransactionResp, schema.TransactionResp], err error) {
	id := a.getNextRequestID()
	var f schema.TransactionSubscribe = 1
	r.ReqId = &id
	r.Subscribe = f
	s = NewSubcription[schema.TransactionResp, schema.TransactionResp](a)
	rsp, err = s.Start(id, r)
	return
}
