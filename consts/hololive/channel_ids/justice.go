package consts

import "cover-utamita/consts"

type Justice struct {
	channelId string
	discordId string
}

// ChannelId implements consts.Constant.
func (h Justice) ChannelId() string {
	return h.channelId
}

func (h Justice) DiscordId() string {
	return h.discordId
}

func (h Justice) GetDiscordId(channelId string) string {
	for _, v := range justices {
		if v.channelId == h.channelId {
			return h.discordId
		}
	}
	return ""
}

var (
	ElizabethRoseBloodflame = Justice{
		channelId: "UCW5uhrG1eCBYditmhL0Ykjw",
		discordId: "1253710511023849555",
	}

	GigiMurin = Justice{
		channelId: "UCDHABijvPBnJm7F-KlNME3w",
		discordId: "1253710543617917041",
	}

	CeciliaImmergreen = Justice{
		channelId: "UCvN5h1ShZtc7nly3pezRayg",
		discordId: "1253710570134179911",
	}

	RaoraPanthera = Justice{
		channelId: "UCl69AEx4MdqMZH7Jtsm7Tig",
		discordId: "1253710594930901055",
	}

	justices = []Justice{ElizabethRoseBloodflame, GigiMurin, CeciliaImmergreen, RaoraPanthera}
	Jutsices = []consts.Constant{ElizabethRoseBloodflame, GigiMurin, CeciliaImmergreen, RaoraPanthera}
)
