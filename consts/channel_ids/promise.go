package consts

type promise string

// Promise
const (
	// IRyS
	IRYS promise = "UC8rcEBzJSleTkf_-agPM20g"

	// Ceres Fauna
	CeresFauna promise = "UCO_aKKYxn4tvrqPjcTzZ6EQ"

	// Ouro Kronii
	OuroKronii promise = "UCmbs8T6MWqUHP1tIQvSgKrg"

	// Nanashi Mumei
	NanashiMumei promise = "UC3n5uGu18FoCy23ggWWp8tA"

	// Hakos Baelz
	HakosBaelz promise = "UCgmPnx-EEeOrZSg5Tiw7ZRQ"
)

// PromiseのチャンネルID
var Promises = []promise{IRYS, CeresFauna, OuroKronii, NanashiMumei, HakosBaelz}
