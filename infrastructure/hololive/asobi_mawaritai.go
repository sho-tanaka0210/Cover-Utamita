package infrastructure

import (
	"cover-utamita/consts"
	"cover-utamita/domain"
	"cover-utamita/infrastructure"
	"os"
)

type AsobiMawaritai struct {
	Members []consts.Constant
}

func (g AsobiMawaritai) SearchUtamita() (results []domain.Result, err error) {

	searcher := infrastructure.UtamitaSearcher{ApiKey: os.Getenv("YOUTUBE_API_KEY")}
	return searcher.SearchUtamita(g.Members)
}
