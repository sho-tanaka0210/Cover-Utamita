package infrastructure

import (
	"cover-utamita/consts"
	"cover-utamita/domain"
	"cover-utamita/infrastructure"
	"os"
)

type Advent struct {
	Members []consts.Constant
}

// 歌ってみたの検索
func (g Advent) SearchUtamita() (results []domain.Result, err error) {

	searcher := infrastructure.UtamitaSearcher{ApiKey: os.Getenv("YOUTUBE_API_KEY")}
	return searcher.SearchUtamita(g.Members)
}
