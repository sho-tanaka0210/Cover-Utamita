package consts

type gamers string

// ゲーマーズ
const (
	// 大神ミオ
	OokamiMio gamers = "UCp-5t9SrOQwXMU7iIjQfARg"

	// 猫又おかゆ
	NekomataOkayu gamers = "UCvaTdHTWBGv3MKj3KVqJVCw"

	// 戌神ころね
	InugamiKorone gamers = "UChAnqc_AY5_I3Px5dig3X1Q"
)

// ゲーマーズのチャンネルID
var Gamers = []gamers{OokamiMio, NekomataOkayu, InugamiKorone}
