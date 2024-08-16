package consts

import "cover-utamita/consts"

type ReGross struct {
	channelId string
	discordId string
}

// ChannelId implements consts.Constant.
func (h ReGross) ChannelId() string {
	return h.channelId
}

func (h ReGross) DiscordId() string {
	return h.discordId
}

func (h ReGross) GetDiscordId(channelId string) string {
	for _, v := range regrosses {
		if v.channelId == h.channelId {
			return h.discordId
		}
	}
	return ""
}

var (
	HiodoshiAo = ReGross{
		channelId: "UCMGfV7TVTmHhEErVJg1oHBQ",
		discordId: "1253706526250766472",
	}

	OtonoseKanade = ReGross{
		channelId: "UCWQtYtq9EOB4-I5P-3fh8lA",
		discordId: "1253706561705082921",
	}

	IchijouRirika = ReGross{
		channelId: "UCtyWhCj3AqKh2dXctLkDtng",
		discordId: "1253706585792974888",
	}

	JuufuuteiRaden = ReGross{
		channelId: "UCdXAk5MpyLD8594lm_OvtGQ",
		discordId: "1253706618646823023",
	}

	TodorokiHajime = ReGross{
		channelId: "UC1iA6_NT4mtAcIII6ygrvCw",
		discordId: "1253706647495508018",
	}

	regrosses = []ReGross{HiodoshiAo, OtonoseKanade, IchijouRirika, JuufuuteiRaden, TodorokiHajime}
	ReGrosses = []consts.Constant{HiodoshiAo, OtonoseKanade, IchijouRirika, JuufuuteiRaden, TodorokiHajime}
)
