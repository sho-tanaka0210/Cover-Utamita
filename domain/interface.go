package domain

import "github.com/bwmarrin/discordgo"

type Result struct {
	DiscordId string `json:"discordId"`
	ChannelId string `json:"channelId"`
	Url       string `json:"url"`
}

type Group interface {
	// 歌ってみたの検索
	SearchUtamita() (results []Result, err error)
}

type Post interface {
	SendMessage(d *discordgo.Session) error
}
