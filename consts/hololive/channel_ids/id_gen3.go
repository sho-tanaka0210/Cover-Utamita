package consts

import "cover-utamita/consts"

type IdGen3 struct {
	channelId string
	discordId string
}

// ChannelId implements consts.Constant.
func (h IdGen3) ChannelId() string {
	return h.channelId
}

func (h IdGen3) DiscordId() string {
	return h.discordId
}

func (h IdGen3) GetDiscordId(channelId string) string {
	for _, v := range idGen3s {
		if v.channelId == h.channelId {
			return h.discordId
		}
	}
	return ""
}

var (
	VestiaZeta = IdGen3{
		channelId: "UCTvHWSfBZgtxE4sILOaurIQ",
		discordId: "1253707215651471391",
	}

	KaelaKovalskia = IdGen3{
		channelId: "UCZLZ8Jjx_RN2CXloOmgTHVg",
		discordId: "1253707243652644885",
	}

	KoboKanaeru = IdGen3{
		channelId: "UCjLEmnpCNeisMxy134KPwWw",
		discordId: "1253708044932616274",
	}

	idGen3s = []IdGen3{VestiaZeta, KaelaKovalskia, KoboKanaeru}
	IdGen3s = []consts.Constant{VestiaZeta, KaelaKovalskia, KoboKanaeru}
)
