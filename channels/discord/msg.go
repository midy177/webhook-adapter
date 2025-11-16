package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/midy177/webhook-adapter/models"
)

const (
	ChannelTypeDiscordWebhook string = "discord-webhook"
)

type MsgWebhook struct {
	discordgo.WebhookParams `json:"inline"`
}

func NewMsgWebhookFromPayload(payload *models.Payload) *MsgWebhook {
	msg := &MsgWebhook{
		discordgo.WebhookParams{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       payload.Title,
					Description: payload.Markdown,
				},
			},
		},
	}

	return msg
}
