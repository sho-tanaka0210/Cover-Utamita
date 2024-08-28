package infrastructure

type Response struct {
	Kind string `json:"kind"`
	ETag string `json:"etag"`

	Items []*struct {
		Id *struct {
			Kind    string `json:"kind"`
			VideoId string `json:"videoId"`
		} `json:"id"`

		Snippet *struct {
			ChannelId string `json:"channelId"`
		} `json:"snippet"`
	} `json:"items"`

	PageInfo *struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
}
