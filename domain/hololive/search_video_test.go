package domain

import (
	channelids "cover-utamita/consts/hololive/channel_ids"
	"cover-utamita/domain"
	"errors"
	"reflect"
	"testing"
)

type fakeUtamitaSearcher struct {
	groups  [][]channelids.Member
	results [][]domain.Result
	quotas  []domain.QuotaUsage
	errors  []error
	calls   int
}

func (f *fakeUtamitaSearcher) SearchUtamita(members []channelids.Member) ([]domain.Result, domain.QuotaUsage, error) {
	f.groups = append(f.groups, members)
	index := f.calls
	f.calls++
	return f.results[index], f.quotas[index], f.errors[index]
}

func TestSearchVideosCombinesGroupsAndQuota(t *testing.T) {
	groups := [][]channelids.Member{
		{{YouTubeChannelID: "youtube-1", DiscordChannelID: "discord-1"}},
		{{YouTubeChannelID: "youtube-2", DiscordChannelID: "discord-2"}},
	}
	searcher := &fakeUtamitaSearcher{
		results: [][]domain.Result{
			{{ChannelId: "youtube-1", DiscordId: "discord-1", Url: "video-1"}},
			{{ChannelId: "youtube-2", DiscordId: "discord-2", Url: "video-2"}},
		},
		quotas: []domain.QuotaUsage{{SearchListRequests: 2, Units: 2}, {SearchListRequests: 3, Units: 3}},
		errors: []error{nil, nil},
	}

	results, quota, err := SearchVideos(searcher, groups)
	if err != nil {
		t.Fatal(err)
	}
	wantResults := append(append([]domain.Result(nil), searcher.results[0]...), searcher.results[1]...)
	if !reflect.DeepEqual(results, wantResults) {
		t.Errorf("results = %+v, want %+v", results, wantResults)
	}
	if want := (domain.QuotaUsage{SearchListRequests: 5, Units: 5}); quota != want {
		t.Errorf("quota = %+v, want %+v", quota, want)
	}
	if !reflect.DeepEqual(searcher.groups, groups) {
		t.Errorf("searched groups = %+v, want %+v", searcher.groups, groups)
	}
}

func TestSearchVideosStopsOnErrorAndKeepsQuota(t *testing.T) {
	groups := [][]channelids.Member{
		{{YouTubeChannelID: "youtube-1", DiscordChannelID: "discord-1"}},
		{{YouTubeChannelID: "youtube-2", DiscordChannelID: "discord-2"}},
		{{YouTubeChannelID: "youtube-3", DiscordChannelID: "discord-3"}},
	}
	searchErr := errors.New("search failed")
	searcher := &fakeUtamitaSearcher{
		results: [][]domain.Result{
			{{ChannelId: "youtube-1"}},
			{{ChannelId: "youtube-2"}},
		},
		quotas: []domain.QuotaUsage{{SearchListRequests: 2, Units: 2}, {SearchListRequests: 1, Units: 1}},
		errors: []error{nil, searchErr},
	}

	results, quota, err := SearchVideos(searcher, groups)
	if !errors.Is(err, searchErr) {
		t.Errorf("error = %v, want %v", err, searchErr)
	}
	if results != nil {
		t.Errorf("results = %+v, want nil", results)
	}
	if want := (domain.QuotaUsage{SearchListRequests: 3, Units: 3}); quota != want {
		t.Errorf("quota = %+v, want %+v", quota, want)
	}
	if !reflect.DeepEqual(searcher.groups, groups[:2]) {
		t.Errorf("searched groups = %+v, want %+v", searcher.groups, groups[:2])
	}
}
