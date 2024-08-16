package consts

import (
	"cover-utamita/consts"
)

type Gen0 struct {
	channelId string
	discordId string
}

// ChannelId implements consts.Constant.
func (h Gen0) ChannelId() string {
	return h.channelId
}

func (h Gen0) DiscordId() string {
	return h.discordId
}

func (h Gen0) GetDiscordId(channelId string) string {
	for _, v := range gen0s {
		if v.channelId == h.channelId {
			return h.discordId
		}
	}
	return ""
}

var (
	TokinoSora = Gen0{
		channelId: "UCp6993wxpyDPHUpavwDFqgg",
		discordId: "1253704712931246110",
	}

	RobocoSan = Gen0{
		channelId: "UCDqI2jOz0weumE8s7paEk6g",
		discordId: "1253704752173285406",
	}

	SakuraMiko = Gen0{
		channelId: "UC-hM6YJuNYVAmUWxeIr9FeA",
		discordId: "1253704783789817979",
	}

	HoshimachiSuisei = Gen0{
		channelId: "UC5CwaMl1eIgY8h02uZw7u8A",
		discordId: "1253704813846462475",
	}

	Azki = Gen0{
		channelId: "UC0TXe_LYZ4scaW2XMyi5_kw",
		discordId: "1253704874856812635",
	}

	gen0s = []Gen0{TokinoSora, RobocoSan, SakuraMiko, HoshimachiSuisei, Azki}
	Gen0s = []consts.Constant{TokinoSora, RobocoSan, SakuraMiko, HoshimachiSuisei, Azki}
)
