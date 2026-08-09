package infrastructure

import (
	"errors"
	"testing"

	"github.com/bwmarrin/discordgo"
)

type fakeDiscordMessenger struct {
	messages []*discordgo.Message
	sent     []string
	readErr  error
}

func (f *fakeDiscordMessenger) ChannelMessages(_ string, _ int, _, _, _ string, _ ...discordgo.RequestOption) ([]*discordgo.Message, error) {
	return f.messages, f.readErr
}

func (f *fakeDiscordMessenger) ChannelMessageSend(_ string, content string, _ ...discordgo.RequestOption) (*discordgo.Message, error) {
	f.sent = append(f.sent, content)
	return &discordgo.Message{Content: content}, nil
}

func TestSendMessageSkipsVideoAlreadyPosted(t *testing.T) {
	const videoURL = "https://www.youtube.com/watch?v=video-id"
	discord := &fakeDiscordMessenger{
		messages: []*discordgo.Message{{Content: videoURL}},
	}

	if err := (Channel{DiscordId: "channel", Url: "video-id"}).SendMessage(discord); err != nil {
		t.Fatal(err)
	}
	if len(discord.sent) != 0 {
		t.Fatalf("sent %d messages, want 0", len(discord.sent))
	}
}

func TestSendMessagePostsNewVideo(t *testing.T) {
	discord := &fakeDiscordMessenger{}

	if err := (Channel{DiscordId: "channel", Url: "video-id"}).SendMessage(discord); err != nil {
		t.Fatal(err)
	}
	if len(discord.sent) != 1 {
		t.Fatalf("sent %d messages, want 1", len(discord.sent))
	}
}

func TestSendMessageDoesNotPostWhenHistoryCannotBeChecked(t *testing.T) {
	discord := &fakeDiscordMessenger{readErr: errors.New("unavailable")}

	if err := (Channel{DiscordId: "channel", Url: "video-id"}).SendMessage(discord); err == nil {
		t.Fatal("SendMessage() error = nil, want an error")
	}
	if len(discord.sent) != 0 {
		t.Fatalf("sent %d messages, want 0", len(discord.sent))
	}
}
