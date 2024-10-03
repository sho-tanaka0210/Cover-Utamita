package domain

import (
	"cover-utamita/consts"
	"fmt"
	"strings"

	"google.golang.org/api/youtube/v3"
)

// YouTubeAPIを用いて動画検索を行う
//
//	service:        YouTube Serviceクライアント
//	channelId:      指定のチャンネルID
//	publishedAfter: どれくらい前の日付から検索するか
//	maxResults:     検索結果の最大数
//
//	results:        検索結果
//	err:            エラー
func SearchVideoes(service *youtube.Service, channelId string, publishedAfter string, maxResults int64) (results *youtube.SearchListResponse, err error) {

	call := service.Search.List([]string{"id", "snippet"}).ChannelId(channelId).PublishedAfter(publishedAfter).MaxResults(maxResults)
	results, err = call.Do()
	if err != nil {
		fmt.Printf("APIリクエストに失敗しました: %v", err)
		return nil, err
	}

	return results, nil
}

// ChannelIdやPublishedAfterからとってきた動画から、動画タイトルによる動画抽出を行う。
//
//	items: 検索結果一覧
//	member: 検索をしているメンバー情報
//
//	results: 抽出結果
func VideoRetrieval(items []*youtube.SearchResult, member consts.Constant) (results []Result) {
	for _, item := range items {
		if item.Id.Kind == "youtube#video" {
			title := item.Snippet.Title
			if titleRetrieval(title) {

				results = append(results, Result{ChannelId: item.Id.ChannelId, Url: item.Id.VideoId, DiscordId: member.DiscordId()})
			}
		}
	}

	return results
}

func titleRetrieval(title string) bool {

	return strings.Contains(title, consts.Utattemita) ||
		strings.Contains(strings.ToLower(title), consts.Cover) ||
		strings.Contains(title, consts.OriginalSong) ||
		strings.Contains(title, consts.Original) ||
		strings.Contains(strings.ToLower(title), consts.CoveredBy) ||
		strings.Contains(title, consts.Mv) ||
		strings.Contains(strings.ToLower(title), consts.Official)
}
