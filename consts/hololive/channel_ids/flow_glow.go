package consts

import "cover-utamita/consts"

type FlowGlow struct {
	channelId string
	discordId string
}

func (h FlowGlow) ChannelId() string {
	return h.channelId
}

func (h FlowGlow) DiscordId() string {
	return h.discordId
}

func (h FlowGlow) GetDiscordId(channelId string) string {
	for _, v := range flowGlows {
		if v.channelId == h.channelId {
			return h.discordId
		}
	}
	return ""
}

var (
	IsakiRiona = FlowGlow{
		channelId: "UC9LSiN9hXI55svYEBrrK-tw",
		discordId: "1304345333085507584",
	}

	KoganeNiko = FlowGlow{
		channelId: "UCuI_opAVX6qbxZY-a-AxFuQ",
		discordId: "1304345368623841331",
	}

	MizumiyaSu = FlowGlow{
		channelId: "UCjk2nKmHzgH5Xy-C5qYRd5A",
		discordId: "1304345394914000916",
	}

	RindoChihaya = FlowGlow{
		channelId: "UCKMWFR6lAstLa7Vbf5dH7ig",
		discordId: "1304345418594910299",
	}

	KikiraraVivi = FlowGlow{
		channelId: "UCGzTVXqMQHa4AgJVJIVvtDQ",
		discordId: "1304345439499325492",
	}

	flowGlows = []FlowGlow{IsakiRiona, KoganeNiko, MizumiyaSu, RindoChihaya, KikiraraVivi}
	FlowGlows = []consts.Constant{IsakiRiona, KoganeNiko, MizumiyaSu, RindoChihaya, KikiraraVivi}
)
