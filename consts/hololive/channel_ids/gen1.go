package consts

import "cover-utamita/consts"

type Gen1 struct {
	channelId string
	discordId string
}

// ChannelId implements consts.Constant.
func (h Gen1) ChannelId() string {
	return h.channelId
}

func (h Gen1) DiscordId() string {
	return h.discordId
}

func (h Gen1) GetDiscordId(channelId string) string {
	for _, v := range gen1s {
		if v.channelId == h.channelId {
			return h.discordId
		}
	}
	return ""
}

var (
	YozoraMel = Gen1{
		channelId: "UCD8HOxPs4Xvsm8H0ZxXGiBw",
		discordId: "1253705014157770843",
	}

	AkiRosenthal = Gen1{
		channelId: "UCFTLzh12_nrtzqBPsTCqenA",
		discordId: "1253705047636709386",
	}

	AkaiHaato = Gen1{
		channelId: "UC1CfXB_kRs3C-zaeTG3oGyg",
		discordId: "1253705093790961805",
	}

	ShirakamiFubuki = Gen1{
		channelId: "UCdn5BQ06XqgXoAxIhbqw5Rg",
		discordId: "1253705132697325679",
	}

	NatsuiroMatsuri = Gen1{
		channelId: "UCQ0UDLQCjY0rmuxCDE38FGg",
		discordId: "1253705170949243007",
	}

	gen1s = []Gen1{YozoraMel, AkiRosenthal, AkaiHaato, ShirakamiFubuki, NatsuiroMatsuri}
	Gen1s = []consts.Constant{YozoraMel, AkiRosenthal, AkaiHaato, ShirakamiFubuki, NatsuiroMatsuri}
)
