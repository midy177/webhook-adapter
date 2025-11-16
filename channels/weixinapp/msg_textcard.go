package weixinapp

import "github.com/midy177/webhook-adapter/models"

func init() {
	Payload2MsgFnMap[MsgTypeTextCard] = NewMsgTextCardFromPayload
}

type TextCard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	BtnText     string `json:"btntext"`
}

// Todo
func NewMsgTextCardFromPayload(payload *models.Payload) *Msg {
	return &Msg{}
}
