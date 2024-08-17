package infrastructure

import (
	"cover-utamita/consts"
	"cover-utamita/domain"
	"cover-utamita/infrastructure"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Justice struct {
	Members []consts.Constant
}

// 歌ってみたの検索
func (g Justice) SearchUtamita() (results []domain.Result, err error) {

	apiKey := os.Getenv("YOUTUBE_API_KEY")

	// 前日の日付を取得
	yesterday := time.Now().AddDate(0, 0, -1)
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}
	t := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, jst)
	publishedAfter := t.Unix()

	for _, member := range g.Members {
		endpoint := fmt.Sprintf("%s?key=%s&part=snippet&channel_id=%s&publishedAfter=%d&maxResults=10&order=date&q=%s", consts.SearchVideo, apiKey, member.ChannelId(), publishedAfter, consts.Query)
		bytes, err := infrastructure.Get(endpoint)
		if err != nil {
			return nil, err
		}

		var response infrastructure.Response
		err = json.Unmarshal(bytes, &response)
		if err != nil {
			return nil, err
		}

		for _, r := range response.Items {
			results = append(results, domain.Result{ChannelId: r.Snippet.ChannelId, Url: r.Id.VideoId})
		}
	}

	return results, nil
}
