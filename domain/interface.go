package domain

import "cover-utamita/infrastructure"

type Group interface {
	// 歌ってみたの検索
	SearchUtamita(url string) (*[]infrastructure.Response, error)
}
