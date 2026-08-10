package domain

import (
	consts "cover-utamita/consts/hololive/channel_ids"
	"cover-utamita/domain"
	infrastructure "cover-utamita/infrastructure/hololive"
)

// 指定の所属グループの歌ってみた動画を取得する。
//
// 範囲: 日本時間の前日0時以上、当日0時未満
//
// 結果: 各チャンネルの歌ってみた動画の一覧
func SearchVideoes() (results []domain.Result, quotaUsage domain.QuotaUsage, err error) {

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
		infrastructure.FlowGlow{Members: consts.FlowGlows},

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
		r, groupQuotaUsage, err := members.SearchUtamita()
		quotaUsage.SearchListRequests += groupQuotaUsage.SearchListRequests
		quotaUsage.Units += groupQuotaUsage.Units
		if err != nil {
			return nil, quotaUsage, err
		}

		results = append(results, r...)
	}

	return results, quotaUsage, nil
}
