package infrastructure

type Response struct {
	Kind          string `json:"kind"`
	ETag          string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	RegionCode    string `json:"regionCode"`

	PageInfo *struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`

	Items []*struct {
		Kind string `json:"kind"`
		ETag string `json:"etag"`

		Id *struct {
			Kind    string `json:"kind"`
			VideoId string `json:"videoId"`
		} `json:"id"`

		Snippet *struct {
			PublishedAt string `json:"publishedAt"`
			ChannelId   string `json:"channelId"`
			Title       string `json:"title"`
			Description string `json:"description"`

			Thumbnails *struct {
				Default *struct {
					Url    string `json:"url"`
					Width  string `json:"width"`
					Height string `json:"height"`
				} `json:"default"`

				Midium *struct {
					Url    string `json:"url"`
					Width  string `json:"width"`
					Height string `json:"height"`
				} `json:"midium"`

				High *struct {
					Url    string `json:"url"`
					Width  string `json:"width"`
					Height string `json:"height"`
				} `json:"high"`
			} `json:"thumbnails"`

			ChannelTitle         string `json:"channelTitle"`
			LiveBroadcastContent string `json:"liveBroadcastContent"`
			PublishTime          string `json:"publishTime"`
		} `json:"snippet"`
	} `json:"items"`
}
