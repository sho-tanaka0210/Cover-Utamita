package infrastructure

import (
	"cover-utamita/domain"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type DiscordMessenger interface {
	ChannelMessages(channelID string, limit int, beforeID, afterID, aroundID string, options ...discordgo.RequestOption) ([]*discordgo.Message, error)
	ChannelMessageSend(channelID string, content string, options ...discordgo.RequestOption) (*discordgo.Message, error)
}

type Channel struct {
	DiscordID string
	URL       string
}

func (c Channel) SendMessage(d DiscordMessenger) error {
	videoURL := domain.VideoURL(c.URL)
	messages, err := d.ChannelMessages(c.DiscordID, 100, "", "", "")
	if err != nil {
		return fmt.Errorf("Discordの投稿履歴取得に失敗しました: %w", err)
	}
	// 再試行時に同じ動画を投稿しないよう、直近の投稿内容を先に確認する。
	for _, message := range messages {
		if message.Content == videoURL {
			return nil
		}
	}

	_, err = d.ChannelMessageSend(c.DiscordID, videoURL)
	if err != nil {
		return fmt.Errorf("Discordへの投稿に失敗しました: %w", err)
	}

	return nil
}
