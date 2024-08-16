package consts

type Constant interface {
	ChannelId() string
	DiscordId() string
	GetDiscordId(channelId string) string
}
