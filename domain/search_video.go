package domain

import (
	consts "cover-utamita/consts/hololive/channel_ids"
	infrastructure "cover-utamita/infrastructure/hololive"
)

// 指定の所属グループの歌ってみた動画を取得する。
//
// 範囲: 実行日から24時間の範囲
//
// 結果: 各チャンネルの歌ってみた動画の一覧
func SearchVideoes() (err error) {

	// ホロライブのメンバーのチャンネルを検索する
	hololive := []Group{
		// JP
		infrastructure.Gen0s{ChannelIds: consts.Gen0s},
		infrastructure.Gen1s{ChannelIds: consts.Gen1s},
		infrastructure.Gamers{ChannelIds: consts.Gamers},
		infrastructure.Gen2s{ChannelIds: consts.Gen2s},
		infrastructure.Gen3s{ChannelIds: consts.Gen3s},
		infrastructure.Gen4s{ChannelIds: consts.Gen4s},
		infrastructure.Gen5s{ChannelIds: consts.Gen5s},
		infrastructure.Holox{ChannelIds: consts.HoloX},

		infrastructure.ReGross{ChannelIds: consts.ReGrosses},

		// ID
		infrastructure.IdGen1{ChannelIds: consts.IdGen1s},
		infrastructure.IdGen2{ChannelIds: consts.IdGen2s},
		infrastructure.IdGen3{ChannelIds: consts.IdGen3s},

		// EN
		infrastructure.Myth{ChannelIds: consts.Mythes},
		infrastructure.Promise{ChannelIds: consts.Promises},
		infrastructure.Advent{ChannelIds: consts.Advents},
		infrastructure.Justice{ChannelIds: consts.Justices},
		infrastructure.Council{ChannelIds: consts.Councils},
	}

	for _, member := range hololive {
		member.SearchUtamita()
	}

	return nil
}
