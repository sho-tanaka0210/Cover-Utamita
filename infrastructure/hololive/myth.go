package infrastructure

import (
	"cover-utamita/consts"
	"cover-utamita/infrastructure"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Myth struct {
	ChannelIds []string
}

// 歌ってみたの検索
func (g Myth) SearchUtamita(url string) (*[]infrastructure.Response, error) {

	apiKey := os.Getenv("YOUTUBE_API_KEY")

	// 前日の日付を取得
	yesterday := time.Now().AddDate(0, 0, -1)
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}
	t := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, jst)
	publishedAfter := t.Unix()

	var responses []infrastructure.Response

	for _, channelId := range g.ChannelIds {
		endpoint := fmt.Sprintf("%s?key=%s&part=snippet&channel_id=%s&publishedAfter=%d&maxResults=10&order=date&q=%s", url, apiKey, channelId, publishedAfter, consts.Query)
		bytes, err := infrastructure.Get(endpoint)
		if err != nil {
			return nil, err
		}

		var response infrastructure.Response
		err = json.Unmarshal(bytes, &response)
		if err != nil {
			return nil, err
		}
		responses = append(responses, response)
	}

	return &responses, nil
}
