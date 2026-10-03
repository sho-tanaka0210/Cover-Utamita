package domain

import "github.com/bwmarrin/discordgo"

type Result struct {
	DiscordID string `json:"discordId"`
	ChannelID string `json:"channelId"`
	URL       string `json:"url"`
}

type QuotaUsage struct {
	SearchListRequests int64
	Units              int64
}

type Post interface {
	SendMessage(d DiscordMessenger) error
}

type DiscordMessenger interface {
	ChannelMessages(channelID string, limit int, beforeID, afterID, aroundID string, options ...discordgo.RequestOption) ([]*discordgo.Message, error)
	ChannelMessageSend(channelID string, content string, options ...discordgo.RequestOption) (*discordgo.Message, error)
}
