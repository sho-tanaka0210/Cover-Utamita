package infrastructure

import (
	"context"
	channelids "cover-utamita/consts/hololive/channel_ids"
	"cover-utamita/domain"
	"fmt"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type UtamitaSearcher struct {
	APIKey string
}

// 歌ってみたの検索
func (g UtamitaSearcher) SearchUtamita(members []channelids.Member) (results []domain.Result, quotaUsage domain.QuotaUsage, err error) {

	service, err := g.prepareService()
	if err != nil {
		return nil, quotaUsage, err
	}

	publishedAfter, publishedBefore := domain.PreviousDayPeriod(time.Now())

	for _, member := range members {
		items, memberQuotaUsage, err := domain.SearchVideos(
			service,
			member.YouTubeChannelID,
			publishedAfter.Format(time.RFC3339),
			publishedBefore.Format(time.RFC3339),
		)
		quotaUsage.SearchListRequests += memberQuotaUsage.SearchListRequests
		quotaUsage.Units += memberQuotaUsage.Units
		if err != nil {
			return nil, quotaUsage, err
		}

		results = append(results, domain.VideoRetrieval(items, member)...)
	}

	return results, quotaUsage, nil
}

// YouTube検索のためのサービス
func (g UtamitaSearcher) prepareService() (*youtube.Service, error) {

	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(g.APIKey))
	if err != nil {
		fmt.Printf("YouTubeサービスの作成に失敗しました: %v", err)
		return nil, err
	}

	return service, nil
}
