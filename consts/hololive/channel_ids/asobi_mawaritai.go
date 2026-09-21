package consts

import "cover-utamita/consts"

type AsobiMawaritai struct {
	channelId string
	discordId string
}

func (h AsobiMawaritai) ChannelId() string {
	return h.channelId
}

func (h AsobiMawaritai) DiscordId() string {
	return h.discordId
}

func (h AsobiMawaritai) GetDiscordId(channelId string) string {
	for _, v := range asobiMawaritais {
		if v.channelId == h.channelId {
			return h.discordId
		}
	}
	return ""
}

var (
	AchichiMela = AsobiMawaritai{
		channelId: "UC8eitCE9Z6EwUCs-VUi1blg",
		discordId: "1551511235390869574",
	}

	SorashinaSopia = AsobiMawaritai{
		channelId: "UCROQtXcp2loQEmvpe5rhJzQ",
		discordId: "1551511280727105587",
	}

	SuzunaTsuzuri = AsobiMawaritai{
		channelId: "UCy9mgxB8pn2C4aNK_MPthDQ",
		discordId: "1551511320094838784",
	}

	HyakutoKyoko = AsobiMawaritai{
		channelId: "UCSjQDxud2HkAO2DVD3lwxmw",
		discordId: "1551511367402397737",
	}

	asobiMawaritais = []AsobiMawaritai{AchichiMela, SorashinaSopia, SuzunaTsuzuri, HyakutoKyoko}
	AsobiMawaritais = []consts.Constant{AchichiMela, SorashinaSopia, SuzunaTsuzuri, HyakutoKyoko}
)
