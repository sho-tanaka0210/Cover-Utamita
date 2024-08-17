package domain

type Result struct {
	ChannelId string
	Url       string
}

type Group interface {
	// 歌ってみたの検索
	SearchUtamita() (result []Result, err error)
}
