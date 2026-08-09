package infrastructure

import (
	"cover-utamita/consts"
	"cover-utamita/domain"
	"fmt"
)

type Channel struct {
	DiscordId string
	Url       string
}

func (c Channel) SendMessage(d domain.DiscordMessenger) error {
	videoUrl := fmt.Sprintf("%s%s", consts.VideoUrl, c.Url)
	messages, err := d.ChannelMessages(c.DiscordId, 100, "", "", "")
	if err != nil {
		return fmt.Errorf("Discordの投稿履歴取得に失敗しました: %w", err)
	}
	// 再試行時に同じ動画を投稿しないよう、直近の投稿内容を先に確認する。
	for _, message := range messages {
		if message.Content == videoUrl {
			return nil
		}
	}

	_, err = d.ChannelMessageSend(c.DiscordId, videoUrl)
	if err != nil {
		return fmt.Errorf("Discordへの投稿に失敗しました: %w", err)
	}

	return nil
}
