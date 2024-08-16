package consts

import "cover-utamita/consts"

type Council struct {
	channelId string
	discordId string
}

// ChannelId implements consts.Constant.
func (h Council) ChannelId() string {
	return h.channelId
}

func (h Council) DiscordId() string {
	return h.discordId
}

func (h Council) GetDiscordId(channelId string) string {
	for _, v := range council {
		if v.channelId == h.channelId {
			return h.discordId
		}
	}
	return ""
}

// Council
var (
	// Tukumo Sana
	TukumoSana = Council{
		channelId: "UCsUj0dszADCGbF3gNrQEuSQ",
		discordId: "1253710759146295377",
	}

	council  = []Council{TukumoSana}
	Councils = []consts.Constant{TukumoSana}
)
