package domain

import "github.com/bwmarrin/discordgo"

type Result struct {
	DiscordId string `json:"discordId"`
	ChannelId string `json:"channelId"`
	Url       string `json:"url"`
}

type QuotaUsage struct {
	SearchListRequests int64
	Units              int64
}

type Group interface {
	// 歌ってみたの検索
	SearchUtamita() (results []Result, quotaUsage QuotaUsage, err error)
}

type Post interface {
	SendMessage(d DiscordMessenger) error
}

type DiscordMessenger interface {
	ChannelMessages(channelID string, limit int, beforeID, afterID, aroundID string, options ...discordgo.RequestOption) ([]*discordgo.Message, error)
	ChannelMessageSend(channelID string, content string, options ...discordgo.RequestOption) (*discordgo.Message, error)
}
