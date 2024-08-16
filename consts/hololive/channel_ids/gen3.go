package consts

import "cover-utamita/consts"

type Gen3 struct {
	channelId string
	discordId string
}

// ChannelId implements consts.Constant.
func (h Gen3) ChannelId() string {
	return h.channelId
}

func (h Gen3) DiscordId() string {
	return h.discordId
}

func (h Gen3) GetDiscordId(channelId string) string {
	for _, v := range gen3s {
		if v.channelId == h.channelId {
			return h.discordId
		}
	}
	return ""
}

var (
	UsadaPekora = Gen3{
		channelId: "UC1DCedRgGHBdm81E1llLhOQ",
		discordId: "1253705706926899252",
	}

	UruhaRusia = Gen3{
		channelId: "UCl_gCybOJRIgOXw6Qb4qJzQ",
		discordId: "1253705730054291466",
	}

	ShiranuiFlare = Gen3{
		channelId: "UCvInZx9h3jC2JzsIzoOebWg",
		discordId: "1253705759506825249",
	}

	ShiroganeNoel = Gen3{
		channelId: "UCdyqAaZDKHXg4Ahi7VENThQ",
		discordId: "1253705797938974881",
	}

	HoushouMarine = Gen3{
		channelId: "UCCzUftO8KOVkV4wQG1vkUvg",
		discordId: "1253705837600313444",
	}

	gen3s = []Gen3{
		UsadaPekora, UruhaRusia, ShiranuiFlare, ShiroganeNoel, HoushouMarine}
	Gen3s = []consts.Constant{
		UsadaPekora, UruhaRusia, ShiranuiFlare, ShiroganeNoel, HoushouMarine}
)
