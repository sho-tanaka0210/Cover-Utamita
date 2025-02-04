package infrastructure

import (
	"context"
	"cover-utamita/consts"
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
func (g UtamitaSearcher) SearchUtamita(members []consts.Constant) (results []domain.Result, err error) {

	service, err := g.prepareService()
	if err != nil {
		return nil, err
	}

	// 前日の日付を取得
	loc, _ := time.LoadLocation("Asia/Tokyo")
	yesterday := time.Now().In(loc).AddDate(0, 0, consts.BeforeDay)
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}
	t := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, jst)

	for _, member := range members {
		response, err := domain.SearchVideoes(service, member.ChannelId(), t.Format(time.RFC3339), consts.MaxResults)
		if err != nil {
			fmt.Printf("APIリクエストに失敗しました: %v", err)
			return nil, err
		}

		results = append(results, domain.VideoRetrieval(response.Items, member)...)
	}

	return results, nil
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
