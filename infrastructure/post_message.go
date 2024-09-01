package infrastructure

import (
	"cover-utamita/consts"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Channel struct {
	DiscordId string
	Url       string
}

func (c Channel) SendMessage(d *discordgo.Session) error {

	videoUrl := fmt.Sprintf("%s%s", consts.VideoUrl, c.Url)
	_, err := d.ChannelMessageSend(c.DiscordId, videoUrl)
	if err != nil {
		fmt.Printf("Discordへの投稿に失敗しました。 : %v", err)
		return err
	}

	return nil
}
