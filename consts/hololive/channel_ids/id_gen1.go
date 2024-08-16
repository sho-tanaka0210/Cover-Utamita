package consts

import "cover-utamita/consts"

type IdGen1 struct {
	channelId string
	discordId string
}

// ChannelId implements consts.Constant.
func (h IdGen1) ChannelId() string {
	return h.channelId
}

func (h IdGen1) DiscordId() string {
	return h.discordId
}

func (h IdGen1) GetDiscordId(channelId string) string {
	for _, v := range idGen1s {
		if v.channelId == h.channelId {
			return h.discordId
		}
	}
	return ""
}

var (
	AyundaRisu = IdGen1{
		channelId: "UCOyYb1c43VlX9rc_lT6NKQw",
		discordId: "1253706969030725725",
	}

	MoonaHoshinova = IdGen1{
		channelId: "UCP0BspO_AMEe3aQqqpo89Dg",
		discordId: "1253707018280239104",
	}

	AiraniIofifteen = IdGen1{
		channelId: "UCAoy6rzhSf4ydcYjJw3WoVg",
		discordId: "1253707057782194319",
	}

	idGen1s = []IdGen1{AyundaRisu, MoonaHoshinova, AiraniIofifteen}
	IdGen1s = []consts.Constant{AyundaRisu, MoonaHoshinova, AiraniIofifteen}
)
