package consts

import "cover-utamita/consts"

type Gamers struct {
	channelId string
	discordId string
}

// ChannelId implements consts.Constant.
func (h Gamers) ChannelId() string {
	return h.channelId
}

func (h Gamers) DiscordId() string {
	return h.discordId
}

func (h Gamers) GetDiscordId(channelId string) string {
	for _, v := range gamers {
		if v.channelId == h.channelId {
			return h.discordId
		}
	}
	return ""
}

var (
	OokamiMio = Gamers{
		channelId: "UCp-5t9SrOQwXMU7iIjQfARg",
		discordId: "1253705472557584405",
	}

	NekomataOkayu = Gamers{
		channelId: "UCvaTdHTWBGv3MKj3KVqJVCw",
		discordId: "1253705501443883139",
	}

	InugamiKorone = Gamers{
		channelId: "UChAnqc_AY5_I3Px5dig3X1Q",
		discordId: "1253705526739468390",
	}

	gamers   = []Gamers{OokamiMio, NekomataOkayu, InugamiKorone}
	Gamerses = []consts.Constant{OokamiMio, NekomataOkayu, InugamiKorone}
)
