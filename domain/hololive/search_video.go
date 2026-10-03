package domain

import (
	channelids "cover-utamita/consts/hololive/channel_ids"
	"cover-utamita/domain"
	"cover-utamita/infrastructure"
	"os"
)

// 指定の所属グループの歌ってみた動画を取得する。
//
// 範囲: 日本時間の前日0時以上、当日0時未満
//
// 結果: 各チャンネルの歌ってみた動画の一覧
func SearchVideoes() (results []domain.Result, quotaUsage domain.QuotaUsage, err error) {
	if err := channelids.ValidateMembers(channelids.SearchGroups); err != nil {
		return nil, quotaUsage, err
	}

	searcher := infrastructure.UtamitaSearcher{APIKey: os.Getenv("YOUTUBE_API_KEY")}
	for _, members := range channelids.SearchGroups {
		r, groupQuotaUsage, err := searcher.SearchUtamita(members)
		quotaUsage.SearchListRequests += groupQuotaUsage.SearchListRequests
		quotaUsage.Units += groupQuotaUsage.Units
		if err != nil {
			return nil, quotaUsage, err
		}
		results = append(results, r...)
	}

	return results, quotaUsage, nil
}
