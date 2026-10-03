package domain

import (
	channelids "cover-utamita/consts/hololive/channel_ids"
	"fmt"

	"google.golang.org/api/youtube/v3"
)

// YouTubeAPIを用いて動画検索を行う
//
//	service:        YouTube Serviceクライアント
//	channelID:      指定のチャンネルID
//	publishedAfter: 検索対象期間の開始日時
//	publishedBefore: 検索対象期間の終了日時
//
//	results:        検索結果
//	quotaUsage:     APIリクエスト数とクォータ消費量
//	err:            エラー
func SearchVideos(service *youtube.Service, channelID string, publishedAfter string, publishedBefore string) (results []*youtube.SearchResult, quotaUsage QuotaUsage, err error) {
	// APIの上限まで取得してページ数を抑える。
	const maxResults int64 = 50
	// search.list専用枠はリクエスト回数と同じ単位で集計する。
	const searchListQuotaUnits int64 = 1
	pageToken := ""
	for {
		call := service.Search.List([]string{"snippet"}).
			ChannelId(channelID).
			PublishedAfter(publishedAfter).
			PublishedBefore(publishedBefore).
			MaxResults(maxResults).
			Order("date").
			Q(searchQuery()).
			Type("video")
		if pageToken != "" {
			call.PageToken(pageToken)
		}

		// 失敗したリクエストもクォータ対象になるため、送信前に加算する。
		quotaUsage.SearchListRequests++
		quotaUsage.Units += searchListQuotaUnits
		response, requestErr := call.Do()
		if requestErr != nil {
			return nil, quotaUsage, fmt.Errorf("APIリクエストに失敗しました: %w", requestErr)
		}

		for _, item := range response.Items {
			if item != nil && item.Snippet != nil && item.Snippet.LiveBroadcastContent != "upcoming" {
				results = append(results, item)
			}
		}
		if response.NextPageToken == "" {
			break
		}
		pageToken = response.NextPageToken
	}

	return results, quotaUsage, nil
}

// ChannelIdやPublishedAfterからとってきた動画から、動画タイトルによる動画抽出を行う。
//
//	items: 検索結果一覧
//	member: 検索をしているメンバー情報
//
//	results: 抽出結果
func VideoRetrieval(items []*youtube.SearchResult, member channelids.Member) (results []Result) {
	for _, item := range items {
		if item != nil && item.Id != nil && item.Snippet != nil && item.Id.Kind == "youtube#video" {
			title := item.Snippet.Title
			if titleRetrieval(title) {
				results = append(results, Result{ChannelID: item.Snippet.ChannelId, URL: item.Id.VideoId, DiscordID: member.DiscordChannelID})
			}
		}
	}

	return results
}
