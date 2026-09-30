package infrastructure

import (
	"context"
	"cover-utamita/consts"
	channelids "cover-utamita/consts/hololive/channel_ids"
	"cover-utamita/domain"
	"fmt"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type UtamitaSearcher struct {
	ApiKey string
}

// 歌ってみたの検索
func (g UtamitaSearcher) SearchUtamita(members []channelids.Member) (results []domain.Result, quotaUsage domain.QuotaUsage, err error) {

	service, err := g.prepareService()
	if err != nil {
		return nil, quotaUsage, err
	}

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, quotaUsage, err
	}
	publishedAfter, publishedBefore := previousDayPeriod(time.Now(), jst)

	for _, member := range members {
		items, memberQuotaUsage, err := domain.SearchVideos(
			service,
			member.YouTubeChannelID,
			publishedAfter.Format(time.RFC3339),
			publishedBefore.Format(time.RFC3339),
			consts.MaxResults,
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

func previousDayPeriod(now time.Time, location *time.Location) (time.Time, time.Time) {
	today := now.In(location)
	publishedBefore := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, location)
	publishedAfter := publishedBefore.AddDate(0, 0, consts.BeforeDay)
	return publishedAfter, publishedBefore
}

// YouTube検索のためのサービス
func (g UtamitaSearcher) prepareService() (*youtube.Service, error) {

	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(g.ApiKey))
	if err != nil {
		fmt.Printf("YouTubeサービスの作成に失敗しました: %v", err)
		return nil, err
	}

	return service, nil
}
