// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "fmt"
import "reflect"
import "encoding/json"

// Response for creating a P2P order review.
type P2POrderReviewResp struct {
	// Echo of the request made.
	EchoReq P2POrderReviewRespEchoReq `json:"echo_req"`

	// Action name of the request made.
	MsgType P2POrderReviewRespMsgType `json:"msg_type"`

	// Details of the created order review.
	P2POrderReview *P2POrderReviewRespP2POrderReview `json:"p2p_order_review,omitempty"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`
}

// Echo of the request made.
type P2POrderReviewRespEchoReq map[string]interface{}

type P2POrderReviewRespMsgType string

const P2POrderReviewRespMsgTypeP2POrderReview P2POrderReviewRespMsgType = "p2p_order_review"

// Details of the created order review.
type P2POrderReviewRespP2POrderReview struct {
	// The reviewed advertiser's identification number.
	AdvertiserId string `json:"advertiser_id"`

	// The epoch time of the review.
	CreatedTime int `json:"created_time"`

	// The order identification number.
	OrderId string `json:"order_id"`

	// Rating for the transaction, 1 to 5.
	Rating int `json:"rating"`

	// `1` if the advertiser is recommended, `0` if not recommended.
	Recommended P2POrderReviewRespP2POrderReviewRecommended `json:"recommended"`
}

type P2POrderReviewRespP2POrderReviewRecommended struct {
	Value interface{}
}

var enumValues_P2POrderReviewRespP2POrderReviewRecommended = []interface{}{
	nil,
	0,
	1,
}

// MarshalJSON implements json.Marshaler.
func (j *P2POrderReviewRespP2POrderReviewRecommended) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2POrderReviewRespP2POrderReviewRecommended) UnmarshalJSON(b []byte) error {
	var v struct {
		Value interface{}
	}
	if err := json.Unmarshal(b, &v.Value); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2POrderReviewRespP2POrderReviewRecommended {
		if reflect.DeepEqual(v.Value, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2POrderReviewRespP2POrderReviewRecommended, v.Value)
	}
	*j = P2POrderReviewRespP2POrderReviewRecommended(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2POrderReviewRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2POrderReviewRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2POrderReviewRespMsgType, v)
	}
	*j = P2POrderReviewRespMsgType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2POrderReviewRespP2POrderReview) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["advertiser_id"]; !ok || v == nil {
		return fmt.Errorf("field advertiser_id: required")
	}
	if v, ok := raw["created_time"]; !ok || v == nil {
		return fmt.Errorf("field created_time: required")
	}
	if v, ok := raw["order_id"]; !ok || v == nil {
		return fmt.Errorf("field order_id: required")
	}
	if v, ok := raw["rating"]; !ok || v == nil {
		return fmt.Errorf("field rating: required")
	}
	if v, ok := raw["recommended"]; !ok || v == nil {
		return fmt.Errorf("field recommended: required")
	}
	type Plain P2POrderReviewRespP2POrderReview
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = P2POrderReviewRespP2POrderReview(plain)
	return nil
}

var enumValues_P2POrderReviewRespMsgType = []interface{}{
	"p2p_order_review",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2POrderReviewResp) UnmarshalJSON(b []byte) error {
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
	type Plain P2POrderReviewResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = P2POrderReviewResp(plain)
	return nil
}