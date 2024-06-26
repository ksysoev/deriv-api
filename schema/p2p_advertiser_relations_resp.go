// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

// Returns information about favourite and blocked advertisers.
type P2PAdvertiserRelationsResp struct {
	// Echo of the request made.
	EchoReq P2PAdvertiserRelationsRespEchoReq `json:"echo_req"`

	// Action name of the request made.
	MsgType P2PAdvertiserRelationsRespMsgType `json:"msg_type"`

	// P2P advertiser relations information.
	P2PAdvertiserRelations *P2PAdvertiserRelationsRespP2PAdvertiserRelations `json:"p2p_advertiser_relations,omitempty"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`
}

// Echo of the request made.
type P2PAdvertiserRelationsRespEchoReq map[string]interface{}

type P2PAdvertiserRelationsRespMsgType string

const P2PAdvertiserRelationsRespMsgTypeP2PAdvertiserRelations P2PAdvertiserRelationsRespMsgType = "p2p_advertiser_relations"

var enumValues_P2PAdvertiserRelationsRespMsgType = []interface{}{
	"p2p_advertiser_relations",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertiserRelationsRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PAdvertiserRelationsRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PAdvertiserRelationsRespMsgType, v)
	}
	*j = P2PAdvertiserRelationsRespMsgType(v)
	return nil
}

// P2P advertiser relations information.
type P2PAdvertiserRelationsRespP2PAdvertiserRelations struct {
	// List of advertisers blocked by the current user.
	BlockedAdvertisers []P2PAdvertiserRelationsRespP2PAdvertiserRelationsBlockedAdvertisersElem `json:"blocked_advertisers"`

	// Favourite advertisers of the current user.
	FavouriteAdvertisers []P2PAdvertiserRelationsRespP2PAdvertiserRelationsFavouriteAdvertisersElem `json:"favourite_advertisers"`
}

// Advertiser details.
type P2PAdvertiserRelationsRespP2PAdvertiserRelationsBlockedAdvertisersElem struct {
	// The epoch time that the advertiser was blocked.
	CreatedTime *int `json:"created_time,omitempty"`

	// Advertiser unique identifer.
	Id *string `json:"id,omitempty"`

	// Advertiser displayed name.
	Name *string `json:"name,omitempty"`
}

// Advertiser details.
type P2PAdvertiserRelationsRespP2PAdvertiserRelationsFavouriteAdvertisersElem struct {
	// The epoch time that the advertiser was set as favourite.
	CreatedTime *int `json:"created_time,omitempty"`

	// Advertiser unique identifer.
	Id *string `json:"id,omitempty"`

	// Advertiser displayed name.
	Name *string `json:"name,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertiserRelationsRespP2PAdvertiserRelations) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["blocked_advertisers"]; raw != nil && !ok {
		return fmt.Errorf("field blocked_advertisers in P2PAdvertiserRelationsRespP2PAdvertiserRelations: required")
	}
	if _, ok := raw["favourite_advertisers"]; raw != nil && !ok {
		return fmt.Errorf("field favourite_advertisers in P2PAdvertiserRelationsRespP2PAdvertiserRelations: required")
	}
	type Plain P2PAdvertiserRelationsRespP2PAdvertiserRelations
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = P2PAdvertiserRelationsRespP2PAdvertiserRelations(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PAdvertiserRelationsResp) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["echo_req"]; raw != nil && !ok {
		return fmt.Errorf("field echo_req in P2PAdvertiserRelationsResp: required")
	}
	if _, ok := raw["msg_type"]; raw != nil && !ok {
		return fmt.Errorf("field msg_type in P2PAdvertiserRelationsResp: required")
	}
	type Plain P2PAdvertiserRelationsResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = P2PAdvertiserRelationsResp(plain)
	return nil
}
