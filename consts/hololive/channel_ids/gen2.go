package consts

import "cover-utamita/consts"

type Gen2 struct {
	channelId string
	discordId string
}

// ChannelId implements consts.Constant.
func (h Gen2) ChannelId() string {
	return h.channelId
}

func (h Gen2) DiscordId() string {
	return h.discordId
}

func (h Gen2) GetDiscordId(channelId string) string {
	for _, v := range gen2s {
		if v.channelId == h.channelId {
			return h.discordId
		}
	}
	return ""
}

var (
	MinatoAqua = Gen2{
		channelId: "UC1opHUrw8rvnsadT-iGp7Cg",
		discordId: "1253705273919410226",
	}

	MurasakiShion = Gen2{
		channelId: "UCXTpFs_3PqI41qX2d9tL2Rw",
		discordId: "1253705307083898891",
	}

	NakiriAyame = Gen2{
		channelId: "UC7fk0CB07ly8oSl0aqKkqFg",
		discordId: "1253705349253435485",
	}

	YudukiChoko = Gen2{
		channelId: "UC1suqwovbL1kzsoaZgFZLKg",
		discordId: "1253705381486661653",
	}

	OozoraSubaru = Gen2{
		channelId: "UCvzGlP9oQwU--Y0r9id_jnA",
		discordId: "1253705410272170005",
	}

	gen2s = []Gen2{
		MinatoAqua, MurasakiShion, NakiriAyame, YudukiChoko, OozoraSubaru}
	Gen2s = []consts.Constant{
		MinatoAqua, MurasakiShion, NakiriAyame, YudukiChoko, OozoraSubaru}
)
