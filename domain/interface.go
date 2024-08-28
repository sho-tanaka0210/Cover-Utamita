package domain

type Result struct {
	DiscordId string
	ChannelId string
	Url       string
}

type Group interface {
	// 歌ってみたの検索
	SearchUtamita() (results []Result, err error)
}

type Post interface {
	SendMessage() error
}
