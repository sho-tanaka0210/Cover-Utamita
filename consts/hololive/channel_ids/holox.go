package consts

import "cover-utamita/consts"

type HoloX struct {
	channelId string
	discordId string
}

// ChannelId implements consts.Constant.
func (h HoloX) ChannelId() string {
	return h.channelId
}

func (h HoloX) DiscordId() string {
	return h.discordId
}

func (h HoloX) GetDiscordId(channelId string) string {
	for _, v := range holoxes {
		if v.channelId == h.channelId {
			return h.discordId
		}
	}
	return ""
}

var (
	LaplusDarknesss = HoloX{
		channelId: "UCENwRMx5Yh42zWpzURebzTw",
		discordId: "1253706274307182613",
	}

	TakaneLui = HoloX{
		channelId: "UCs9_O1tRPMQTHQ-N_L6FU2g",
		discordId: "1253706305206751292",
	}

	HakuiKoyori = HoloX{
		channelId: "UC6eWCld0KwmyHFbAqK3V-Rw",
		discordId: "1253706337464877179",
	}

	SakamataKuroe = HoloX{
		channelId: "UCIBY1ollUsauvVi4hW4cumw",
		discordId: "1253706383967125525",
	}

	KazamaIroha = HoloX{
		channelId: "UC_vMYWcDjmfdpH6r4TTn1MQ",
		discordId: "1253706414862368841",
	}

	holoxes = []HoloX{LaplusDarknesss, TakaneLui, HakuiKoyori, SakamataKuroe, KazamaIroha}
	Holoxes = []consts.Constant{LaplusDarknesss, TakaneLui, HakuiKoyori, SakamataKuroe, KazamaIroha}
)
