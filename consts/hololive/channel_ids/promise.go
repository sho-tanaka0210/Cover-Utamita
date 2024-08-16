package consts

import "cover-utamita/consts"

type Promise struct {
	channelId string
	discordId string
}

// ChannelId implements consts.Constant.
func (h Promise) ChannelId() string {
	return h.channelId
}

func (h Promise) DiscordId() string {
	return h.discordId
}

func (h Promise) GetDiscordId(channelId string) string {
	for _, v := range promises {
		if v.channelId == h.channelId {
			return h.discordId
		}
	}
	return ""
}

var (
	IRyS = Promise{
		channelId: "UC8rcEBzJSleTkf_-agPM20g",
		discordId: "1253708273119399997",
	}

	CeresFauna = Promise{
		channelId: "UCO_aKKYxn4tvrqPjcTzZ6EQ",
		discordId: "1253709361298018304",
	}

	OuroKronii = Promise{
		channelId: "UCmbs8T6MWqUHP1tIQvSgKrg",
		discordId: "1253710181012078624",
	}

	NanashiMumei = Promise{
		channelId: "UC3n5uGu18FoCy23ggWWp8tA",
		discordId: "1253710213635112960",
	}

	HakosBaelz = Promise{
		channelId: "UCgmPnx-EEeOrZSg5Tiw7ZRQ",
		discordId: "1253710247831277752",
	}

	promises = []Promise{IRyS, CeresFauna, OuroKronii, NanashiMumei, HakosBaelz}
	Promises = []consts.Constant{IRyS, CeresFauna, OuroKronii, NanashiMumei, HakosBaelz}
)
