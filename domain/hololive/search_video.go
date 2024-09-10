package domain

import (
	consts "cover-utamita/consts/hololive/channel_ids"
	"cover-utamita/domain"
	infrastructure "cover-utamita/infrastructure/hololive"
)

// 指定の所属グループの歌ってみた動画を取得する。
//
// 範囲: 実行日から24時間の範囲
//
// 結果: 各チャンネルの歌ってみた動画の一覧
func SearchVideoes() (results []domain.Result, err error) {

	hololive := []domain.Group{
		// ALL
		infrastructure.HololiveAll{Members: consts.Whole},

		// JP
		infrastructure.Gen0{Members: consts.Gen0s},
		infrastructure.Gen1{Members: consts.Gen1s},
		infrastructure.Gamers{Members: consts.Gamerses},
		infrastructure.Gen2{Members: consts.Gen2s},
		infrastructure.Gen3{Members: consts.Gen3s},
		infrastructure.Gen4{Members: consts.Gen4s},
		infrastructure.Gen5{Members: consts.Gen5s},

		infrastructure.HoloX{Members: consts.Holoxes},

		infrastructure.ReGross{Members: consts.ReGrosses},

		// ID
		infrastructure.IdGen1{Members: consts.IdGen1s},
		infrastructure.IdGen2{Members: consts.IdGen2s},
		infrastructure.IdGen3{Members: consts.IdGen3s},

		// EN
		infrastructure.Myth{Members: consts.Mythes},
		infrastructure.Promise{Members: consts.Promises},
		infrastructure.Justice{Members: consts.Jutsices},
		infrastructure.Advent{Members: consts.Advents},
	}

	for _, members := range hololive {
		r, err := members.SearchUtamita()
		if err != nil {
			return nil, err
		}

		results = append(results, r...)
	}

	return results, nil
}
