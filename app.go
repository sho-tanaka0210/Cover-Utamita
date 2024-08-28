package main

import (
	domain "cover-utamita/domain/hololive"
	"cover-utamita/infrastructure"

	"github.com/bwmarrin/discordgo"
)

func App(d *discordgo.Session) error {
	// hololive
	r, err := domain.SearchVideoes()
	if err != nil {
		return err
	}

	for _, v := range r {
		c := infrastructure.Channel{ChannelId: v.ChannelId, Url: v.Url, DiscordId: v.DiscordId}
		err = c.SendMessage(d)
		if err != nil {
			return err
		}
	}

	return nil
}
