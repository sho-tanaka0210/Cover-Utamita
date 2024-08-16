package consts

import "cover-utamita/consts"

type Advent struct {
	channelId string
	discordId string
}

// ChannelId implements consts.Constant.
func (h Advent) ChannelId() string {
	return h.channelId
}

func (h Advent) DiscordId() string {
	return h.discordId
}

func (h Advent) GetDiscordId(channelId string) string {
	for _, v := range advents {
		if v.channelId == h.channelId {
			return h.discordId
		}
	}
	return ""
}

var (
	ShioriNovella = Advent{
		channelId: "UCgnfPPb9JI3e9A4cXHnWbyg",
		discordId: "1253710291770802308",
	}

	KosekiBijou = Advent{
		channelId: "UC9p_lqQ0FEDz327Vgf5JwqA",
		discordId: "1253710317616103424",
	}

	NerissaRavencroft = Advent{
		channelId: "UC_sFNM0z0MWm9A6WlKPuMMg",
		discordId: "1253710344476561499",
	}

	FuwaMoco = Advent{
		channelId: "UCt9H_RpQzhxzlyBxFqrdHqA",
		discordId: "1253710376109867068",
	}

	advents = []Advent{ShioriNovella, KosekiBijou, NerissaRavencroft, FuwaMoco}
	Advents = []consts.Constant{ShioriNovella, KosekiBijou, NerissaRavencroft, FuwaMoco}
)
