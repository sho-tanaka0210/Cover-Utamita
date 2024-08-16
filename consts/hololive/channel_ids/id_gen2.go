package consts

import "cover-utamita/consts"

type IdGen2 struct {
	channelId string
	discordId string
}

// ChannelId implements consts.Constant.
func (h IdGen2) ChannelId() string {
	return h.channelId
}

func (h IdGen2) DiscordId() string {
	return h.discordId
}

func (h IdGen2) GetDiscordId(channelId string) string {
	for _, v := range idGen2s {
		if v.channelId == h.channelId {
			return h.discordId
		}
	}
	return ""
}

var (
	KureijiOllie = IdGen2{
		channelId: "UCYz_5n-uDuChHtLo7My1HnQ",
		discordId: "1253707100279017573",
	}

	AnyaMelfissa = IdGen2{
		channelId: "UC727SQYUvx5pDDGQpTICNWg",
		discordId: "1253707138534998178",
	}

	PavoliaReine = IdGen2{
		channelId: "UChgTyjG-pdNvxxhdsXfHQ5Q",
		discordId: "1253707166175727676",
	}

	idGen2s = []IdGen2{KureijiOllie, AnyaMelfissa, PavoliaReine}
	IdGen2s = []consts.Constant{KureijiOllie, AnyaMelfissa, PavoliaReine}
)
