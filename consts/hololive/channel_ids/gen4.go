package consts

import "cover-utamita/consts"

type Gen4 struct {
	channelId string
	discordId string
}

// ChannelId implements consts.Constant.
func (h Gen4) ChannelId() string {
	return h.channelId
}

func (h Gen4) DiscordId() string {
	return h.discordId
}

func (h Gen4) GetDiscordId(channelId string) string {
	for _, v := range gen4s {
		if v.channelId == h.channelId {
			return h.discordId
		}
	}
	return ""
}

var (
	AmaneKanata = Gen4{
		channelId: "UCZlDXzGoo7d44bwdNObFacg",
		discordId: "1253705955858976848",
	}

	KiryuuCoco = Gen4{
		channelId: "UCS9uQI-jC3DE0L4IpXyvr6w",
		discordId: "1253705982228430930",
	}

	TsunomakiWatame = Gen4{
		channelId: "UCqm3BQLlJfvkTsX_hvm0UmA",
		discordId: "1253706008052633651",
	}

	TokoyamiTowa = Gen4{
		channelId: "UC1uv2Oq6kNxgATlCiez59hw",
		discordId: "1253706034833260635",
	}

	HimemoriLuna = Gen4{
		channelId: "UCa9Y57gfeY0Zro_noHRVrnw",
		discordId: "1253706063786676265",
	}

	gen4s = []Gen4{
		AmaneKanata, KiryuuCoco, TsunomakiWatame, TokoyamiTowa, HimemoriLuna}
	Gen4s = []consts.Constant{
		AmaneKanata, KiryuuCoco, TsunomakiWatame, TokoyamiTowa, HimemoriLuna}
)
