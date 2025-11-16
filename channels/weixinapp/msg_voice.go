package weixinapp

import "github.com/midy177/webhook-adapter/models"

func init() {
	Payload2MsgFnMap[MsgTypeVoice] = NewMsgVoiceFromPayload
}

type Voice struct {
	MediaID string `json:"media_id"`
}

// Todo
func NewMsgVoiceFromPayload(payload *models.Payload) *Msg {
	return &Msg{}
}
