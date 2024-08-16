package consts

import "cover-utamita/consts"

type Myth struct {
	channelId string
	discordId string
}

// ChannelId implements consts.Constant.
func (h Myth) ChannelId() string {
	return h.channelId
}

func (h Myth) DiscordId() string {
	return h.discordId
}

func (h Myth) GetDiscordId(channelId string) string {
	for _, v := range mythes {
		if v.channelId == h.channelId {
			return h.discordId
		}
	}
	return ""
}

var (
	MoriCalliope = Myth{
		channelId: "UCL_qhgtOy0dy1Agp8vkySQg",
		discordId: "1253708092630106233",
	}

	TakanashiKiara = Myth{
		channelId: "UCHsx4Hqa-1ORjQTh9TYDhww",
		discordId: "1253708129271808104",
	}

	NinomaeInanis = Myth{
		channelId: "UCMwGHR0BTZuLsmjY_NT5Pwg",
		discordId: "1253708164323475538",
	}

	GawrGura = Myth{
		channelId: "UCoSrY_IQQVpmIRZ9Xf-y93g",
		discordId: "1253708198616240149",
	}

	WatsonAmelia = Myth{
		channelId: "UCyl1z3jo3XHR1riLFKG5UAg",
		discordId: "1253708226302840893",
	}

	mythes = []Myth{MoriCalliope, TakanashiKiara, NinomaeInanis, GawrGura, WatsonAmelia}
	Mythes = []consts.Constant{MoriCalliope, TakanashiKiara, NinomaeInanis, GawrGura, WatsonAmelia}
)
