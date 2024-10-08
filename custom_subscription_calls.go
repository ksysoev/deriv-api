package deriv

import (
	"context"

	"github.com/ksysoev/deriv-api/schema"
)

// SubscribeTicksHistory Get historic tick data for a given symbol.
//
//nolint:gocritic // don't want to break backward compatibility for now
func (api *Client) SubscribeTicksHistory(ctx context.Context, r schema.TicksHistory) (rsp schema.TicksHistoryResp, s *Subsciption[schema.TicksHistoryResp, schema.TicksResp], err error) {
	var f schema.TicksHistorySubscribe = 1

	id := api.getNextRequestID()
	r.ReqId = &id
	r.Subscribe = &f
	r.Style = schema.TicksHistoryStyleTicks
	s = NewSubcription[schema.TicksHistoryResp, schema.TicksResp](ctx, api)
	rsp, err = s.Start(id, r)

	return
}

// SubscribeTicksHistory Get historic candles data for a given symbol.
//
//nolint:gocritic // don't want to break backward compatibility for now
func (api *Client) SubscribeCandlesHistory(ctx context.Context, r schema.TicksHistory) (rsp schema.TicksHistoryResp, s *Subsciption[schema.TicksHistoryResp, schema.TicksHistoryResp], err error) {
	var f schema.TicksHistorySubscribe = 1

	id := api.getNextRequestID()
	r.ReqId = &id
	r.Subscribe = &f
	r.Style = schema.TicksHistoryStyleCandles
	s = NewSubcription[schema.TicksHistoryResp, schema.TicksHistoryResp](ctx, api)
	rsp, err = s.Start(id, r)

	return
}

// SubscribeTransaction Subscribe to transaction notifications
func (api *Client) SubscribeTransaction(ctx context.Context, r schema.Transaction) (rsp schema.TransactionResp, s *Subsciption[schema.TransactionResp, schema.TransactionResp], err error) {
	var f schema.TransactionSubscribe = 1

	id := api.getNextRequestID()
	r.ReqId = &id
	r.Subscribe = f
	s = NewSubcription[schema.TransactionResp, schema.TransactionResp](ctx, api)
	rsp, err = s.Start(id, r)

	return
}
