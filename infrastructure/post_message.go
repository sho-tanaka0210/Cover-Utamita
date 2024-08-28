package infrastructure

import "github.com/bwmarrin/discordgo"

type Channel struct {
	DiscordId string
	ChannelId string
	Url       string
}

func (c Channel) SendMessage(d *discordgo.Session) error {

	_, err := d.ChannelMessageSend(c.DiscordId, c.Url)
	if err != nil {
		return err
	}

	return nil
}
