package infrastructure

type Response struct {
	Items []*struct {
		Id *struct {
			Kind    string `json:"kind"`
			VideoId string `json:"videoId"`
		} `json:"id"`

		Snippet *struct {
			ChannelId string `json:"channelId"`
		} `json:"snippet"`
	} `json:"items"`
}
