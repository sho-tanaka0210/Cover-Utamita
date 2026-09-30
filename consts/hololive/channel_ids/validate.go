package consts

import (
	"fmt"
	"strings"
)

func ValidateMembers(groups [][]Member) error {
	youtubeIDs := make(map[string]struct{})
	discordIDs := make(map[string]struct{})
	for groupIndex, group := range groups {
		for memberIndex, member := range group {
			if strings.TrimSpace(member.YouTubeChannelID) == "" || strings.TrimSpace(member.DiscordChannelID) == "" {
				return fmt.Errorf("group %d member %d: empty channel ID", groupIndex, memberIndex)
			}
			if _, exists := youtubeIDs[member.YouTubeChannelID]; exists {
				return fmt.Errorf("group %d member %d: duplicate YouTube channel ID %q", groupIndex, memberIndex, member.YouTubeChannelID)
			}
			if _, exists := discordIDs[member.DiscordChannelID]; exists {
				return fmt.Errorf("group %d member %d: duplicate Discord channel ID %q", groupIndex, memberIndex, member.DiscordChannelID)
			}
			youtubeIDs[member.YouTubeChannelID] = struct{}{}
			discordIDs[member.DiscordChannelID] = struct{}{}
		}
	}
	return nil
}
