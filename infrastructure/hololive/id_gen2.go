package infrastructure

import (
	"context"
	"cover-utamita/consts"
	"cover-utamita/domain"
	"log"
	"os"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type IdGen2 struct {
	Members []consts.Constant
}

// 歌ってみたの検索
func (g IdGen2) SearchUtamita() (results []domain.Result, err error) {

	service, err := g.prepareService()
	if err != nil {
		return nil, err
	}

	// 前日の日付を取得
	yesterday := time.Now().AddDate(0, 0, consts.BeforeDay)
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}
	t := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, jst)

	for _, member := range g.Members {
		call := service.Search.List([]string{"id", "snippet"}).Q(consts.Query).ChannelId(member.ChannelId()).PublishedAfter(t.Format(time.RFC3339)).MaxResults(10)
		response, err := call.Do()
		if err != nil {
			log.Fatalf("APIリクエストに失敗しました: %v", err)
			return nil, err
		}

		for _, item := range response.Items {
			if item.Id.Kind == "youtube#video" {
				results = append(results, domain.Result{ChannelId: item.Id.ChannelId, Url: item.Id.VideoId})
			}
		}
	}

	return results, nil
}

// YouTube検索のためのサービス
func (g IdGen2) prepareService() (*youtube.Service, error) {

	apiKey := os.Getenv("YOUTUBE_API_KEY")
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("YouTubeサービスの作成に失敗しました: %v", err)
		return nil, err
	}

	return service, nil
}
