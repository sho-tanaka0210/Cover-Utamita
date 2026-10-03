package infrastructure

type Response struct {
	Kind string `json:"kind"`
	ETag string `json:"etag"`

	Items []*struct {
		ID *struct {
			Kind    string `json:"kind"`
			VideoID string `json:"videoId"`
		} `json:"id"`

		Snippet *struct {
			ChannelID string `json:"channelId"`
		} `json:"snippet"`
	} `json:"items"`

	PageInfo *struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
}
