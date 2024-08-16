package consts

import "cover-utamita/consts"

type HololiveWhole struct {
	channelId string
	discordId string
}

// ChannelId implements consts.Constant.
func (h HololiveWhole) ChannelId() string {
	return h.channelId
}

func (h HololiveWhole) DiscordId() string {
	return h.discordId
}

func (h HololiveWhole) GetDiscordId(channelId string) string {
	for _, v := range whole {
		if v.channelId == h.channelId {
			return h.discordId
		}
	}
	return ""
}

var (
	HololiveAll = HololiveWhole{
		channelId: "",
		discordId: "1253710858240921740",
	}

	HololiveENAll = HololiveWhole{
		channelId: "",
		discordId: "1259585771833135144",
	}

	HololiveIDAll = HololiveWhole{
		channelId: "",
		discordId: "1259585707219877931",
	}

	HololiveDEVIS = HololiveWhole{
		channelId: "",
		discordId: "1259586334935482468",
	}

	whole = []HololiveWhole{HololiveAll, HololiveENAll, HololiveIDAll, HololiveDEVIS}
	Whole = []consts.Constant{HololiveAll, HololiveENAll, HololiveIDAll, HololiveDEVIS}
)
