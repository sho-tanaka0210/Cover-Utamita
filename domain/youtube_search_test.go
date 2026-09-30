package domain

import (
	"context"
	"cover-utamita/consts"
	channelids "cover-utamita/consts/hololive/channel_ids"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func TestSearchVideosUsesConditionsAndRetrievesEveryPage(t *testing.T) {
	const (
		publishedAfter  = "2026-08-09T00:00:00+09:00"
		publishedBefore = "2026-08-10T00:00:00+09:00"
	)

	var queries []url.Values
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		queries = append(queries, r.URL.Query())
		w.Header().Set("Content-Type", "application/json")

		response := map[string]any{
			"items": []map[string]any{
				{
					"id":      map[string]any{"kind": "youtube#video", "videoId": "first"},
					"snippet": map[string]any{"channelId": "channel", "liveBroadcastContent": "none", "title": "first cover"},
				},
				{
					"id":      map[string]any{"kind": "youtube#video", "videoId": "upcoming"},
					"snippet": map[string]any{"channelId": "channel", "liveBroadcastContent": "upcoming", "title": "upcoming cover"},
				},
			},
		}
		if r.URL.Query().Get("pageToken") == "next-page" {
			response["items"] = []map[string]any{
				{
					"id":      map[string]any{"kind": "youtube#video", "videoId": "second"},
					"snippet": map[string]any{"channelId": "channel", "liveBroadcastContent": "none", "title": "second cover"},
				},
			}
		} else {
			response["nextPageToken"] = "next-page"
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	service, err := youtube.NewService(
		context.Background(),
		option.WithHTTPClient(server.Client()),
		option.WithoutAuthentication(),
	)
	if err != nil {
		t.Fatal(err)
	}
	service.BasePath = server.URL + "/"

	items, quotaUsage, err := SearchVideos(service, "channel", publishedAfter, publishedBefore, consts.MaxResults)
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 2 {
		t.Fatalf("len(items) = %d, want 2", len(items))
	}
	if items[0].Id.VideoId != "first" || items[1].Id.VideoId != "second" {
		t.Fatalf("video IDs = %q, %q, want first, second", items[0].Id.VideoId, items[1].Id.VideoId)
	}
	if quotaUsage.SearchListRequests != 2 || quotaUsage.Units != 2*consts.SearchListQuotaUnits {
		t.Fatalf("quota usage = %+v, want 2 requests and %d units", quotaUsage, 2*consts.SearchListQuotaUnits)
	}
	if len(queries) != 2 {
		t.Fatalf("request count = %d, want 2", len(queries))
	}

	wantConditions := map[string]string{
		"channelId":       "channel",
		"part":            "snippet",
		"publishedAfter":  publishedAfter,
		"publishedBefore": publishedBefore,
		"maxResults":      "50",
		"order":           "date",
		"q":               consts.Query,
		"type":            "video",
	}
	for key, want := range wantConditions {
		if got := queries[0].Get(key); got != want {
			t.Errorf("%s = %q, want %q", key, got, want)
		}
	}
	if got := queries[0].Get("pageToken"); got != "" {
		t.Errorf("first pageToken = %q, want empty", got)
	}
	if got := queries[1].Get("pageToken"); got != "next-page" {
		t.Errorf("second pageToken = %q, want next-page", got)
	}
}

func TestVideoRetrievalUsesSnippetChannelIDAndFiltersTitle(t *testing.T) {
	member := channelids.Member{YouTubeChannelID: "requested-channel", DiscordChannelID: "discord"}
	items := []*youtube.SearchResult{
		{
			Id:      &youtube.ResourceId{Kind: "youtube#video", VideoId: "matched", ChannelId: "wrong-channel"},
			Snippet: &youtube.SearchResultSnippet{ChannelId: "source-channel", Title: "NEW SONG COVER"},
		},
		{
			Id:      &youtube.ResourceId{Kind: "youtube#video", VideoId: "unmatched"},
			Snippet: &youtube.SearchResultSnippet{ChannelId: "source-channel", Title: "雑談配信"},
		},
	}

	results := VideoRetrieval(items, member)
	if len(results) != 1 {
		t.Fatalf("len(results) = %d, want 1", len(results))
	}
	want := Result{ChannelId: "source-channel", Url: "matched", DiscordId: "discord"}
	if results[0] != want {
		t.Fatalf("result = %+v, want %+v", results[0], want)
	}
}

func TestTitleRetrieval(t *testing.T) {
	tests := []struct {
		name  string
		title string
		want  bool
	}{
		{name: "Japanese", title: "新作を歌ってみた", want: true},
		{name: "case insensitive", title: "NEW SONG COVER", want: true},
		{name: "original song", title: "My Original Song", want: true},
		{name: "unmatched", title: "ゲーム配信のお知らせ", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleRetrieval(tt.title); got != tt.want {
				t.Fatalf("titleRetrieval(%q) = %t, want %t", tt.title, got, tt.want)
			}
		})
	}
}
