package consts

import "cover-utamita/consts"

type Gen5 struct {
	channelId string
	discordId string
}

// ChannelId implements consts.Constant.
func (h Gen5) ChannelId() string {
	return h.channelId
}

func (h Gen5) DiscordId() string {
	return h.discordId
}

func (h Gen5) GetDiscordId(channelId string) string {
	for _, v := range gen5s {
		if v.channelId == h.channelId {
			return h.discordId
		}
	}
	return ""
}

var (
	YukihanaLamy = Gen5{
		channelId: "UCFKOVgVbGmX65RxO3EtH3iw",
		discordId: "1253706098716704768",
	}

	MomosuzuNene = Gen5{
		channelId: "UCAWSyEs_Io8MtpY3m-zqILA",
		discordId: "1253706127212937216",
	}

	ShishiroBotan = Gen5{
		channelId: "UCUKD-uaobj9jiqB-VXt71mA",
		discordId: "1253706161140666418",
	}

	ManoAroe = Gen5{
		channelId: "UCgZuwn-O7Szh9cAgHqJ6vjw",
		discordId: "1253706191381725196",
	}

	OmaruPolka = Gen5{
		channelId: "UCK9V2B22uJYu3N7eR_BT9QA",
		discordId: "1253706224386572403",
	}

	gen5s = []Gen5{YukihanaLamy, MomosuzuNene, ShishiroBotan, ManoAroe, OmaruPolka}
	Gen5s = []consts.Constant{YukihanaLamy, MomosuzuNene, ShishiroBotan, ManoAroe, OmaruPolka}
)
