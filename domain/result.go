package domain

type Result struct {
	DiscordID string `json:"discordId"`
	ChannelID string `json:"channelId"`
	URL       string `json:"url"`
}

type QuotaUsage struct {
	SearchListRequests int64
	Units              int64
}
