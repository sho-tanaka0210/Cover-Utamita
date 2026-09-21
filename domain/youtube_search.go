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
//	publishedAfter: 検索対象期間の開始日時
//	publishedBefore: 検索対象期間の終了日時
//	maxResults:     検索結果の最大数
//
//	results:        検索結果
//	quotaUsage:     APIリクエスト数とクォータ消費量
//	err:            エラー
func SearchVideos(service *youtube.Service, channelId string, publishedAfter string, publishedBefore string, maxResults int64) (results []*youtube.SearchResult, quotaUsage QuotaUsage, err error) {
	pageToken := ""
	for {
		call := service.Search.List([]string{"snippet"}).
			ChannelId(channelId).
			PublishedAfter(publishedAfter).
			PublishedBefore(publishedBefore).
			MaxResults(maxResults).
			Order("date").
			Q(consts.Query).
			Type("video")
		if pageToken != "" {
			call.PageToken(pageToken)
		}

		// 失敗したリクエストもクォータ対象になるため、送信前に加算する。
		quotaUsage.SearchListRequests++
		quotaUsage.Units += consts.SearchListQuotaUnits
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
func VideoRetrieval(items []*youtube.SearchResult, member consts.Constant) (results []Result) {
	for _, item := range items {
		if item != nil && item.Id != nil && item.Snippet != nil && item.Id.Kind == "youtube#video" {
			title := item.Snippet.Title
			if titleRetrieval(title) {
				results = append(results, Result{ChannelId: item.Snippet.ChannelId, Url: item.Id.VideoId, DiscordId: member.DiscordId()})
			}
		}
	}

	return results
}

func titleRetrieval(title string) bool {
	title = strings.ToLower(title)
	return strings.Contains(title, consts.Utattemita) ||
		strings.Contains(title, consts.Cover) ||
		strings.Contains(title, consts.OriginalSong) ||
		strings.Contains(title, consts.Original) ||
		strings.Contains(title, consts.CoveredBy) ||
		strings.Contains(title, consts.Mv) ||
		strings.Contains(title, consts.Official) ||
		strings.Contains(title, consts.OriginalKyoku)
}
