// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "fmt"
import "reflect"
import "encoding/json"

// Echo of the request made.
type P2PChatCreateRespEchoReq map[string]interface{}

type P2PChatCreateRespMsgType string

var enumValues_P2PChatCreateRespMsgType = []interface{}{
	"p2p_chat_create",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PChatCreateRespMsgType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_P2PChatCreateRespMsgType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_P2PChatCreateRespMsgType, v)
	}
	*j = P2PChatCreateRespMsgType(v)
	return nil
}

const P2PChatCreateRespMsgTypeP2PChatCreate P2PChatCreateRespMsgType = "p2p_chat_create"

// Information of the P2P chat.
type P2PChatCreateRespP2PChatCreate struct {
	// The URL to be used to initialise the chat for the requested order.
	ChannelUrl string `json:"channel_url"`

	// The unique identifier for the order that the chat belongs to.
	OrderId string `json:"order_id"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PChatCreateRespP2PChatCreate) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["channel_url"]; !ok || v == nil {
		return fmt.Errorf("field channel_url: required")
	}
	if v, ok := raw["order_id"]; !ok || v == nil {
		return fmt.Errorf("field order_id: required")
	}
	type Plain P2PChatCreateRespP2PChatCreate
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = P2PChatCreateRespP2PChatCreate(plain)
	return nil
}

// Information of the created P2P chat.
type P2PChatCreateResp struct {
	// Echo of the request made.
	EchoReq P2PChatCreateRespEchoReq `json:"echo_req"`

	// Action name of the request made.
	MsgType P2PChatCreateRespMsgType `json:"msg_type"`

	// Information of the P2P chat.
	P2PChatCreate *P2PChatCreateRespP2PChatCreate `json:"p2p_chat_create,omitempty"`

	// Optional field sent in request to map to response, present only when request
	// contains `req_id`.
	ReqId *int `json:"req_id,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *P2PChatCreateResp) UnmarshalJSON(b []byte) error {
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
	type Plain P2PChatCreateResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = P2PChatCreateResp(plain)
	return nil
}