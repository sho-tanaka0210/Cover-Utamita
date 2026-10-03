package domain

import (
	channelids "cover-utamita/consts/hololive/channel_ids"
	"cover-utamita/domain"
)

type UtamitaSearcher interface {
	SearchUtamita(members []channelids.Member) ([]domain.Result, domain.QuotaUsage, error)
}

func SearchVideos(searcher UtamitaSearcher, groups [][]channelids.Member) (results []domain.Result, quotaUsage domain.QuotaUsage, err error) {
	if err := channelids.ValidateMembers(groups); err != nil {
		return nil, quotaUsage, err
	}

	for _, members := range groups {
		r, groupQuotaUsage, err := searcher.SearchUtamita(members)
		quotaUsage.SearchListRequests += groupQuotaUsage.SearchListRequests
		quotaUsage.Units += groupQuotaUsage.Units
		if err != nil {
			return nil, quotaUsage, err
		}
		results = append(results, r...)
	}

	return results, quotaUsage, nil
}
