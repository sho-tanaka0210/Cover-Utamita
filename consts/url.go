package consts

// GET
const (
	Query    string = `歌ってみた|cover|"original song"|original|"covered by"|mv|official|オリジナル曲`
	VideoUrl string = "https://www.youtube.com/watch?v="
	// API側の上限まで取得し、ページ数とクォータ消費を抑えるため最大値を使用する。
	MaxResults int64 = 50

	// Search Queries専用枠はリクエスト回数と同じ単位で集計される。
	SearchListQuotaUnits int64 = 1
)

// Request Parameter
const (
	Utattemita    string = "歌ってみた"
	Cover         string = "cover"
	OriginalSong  string = "original song"
	Original      string = "original"
	CoveredBy     string = "covered by"
	Mv            string = "mv"
	Official      string = "official"
	OriginalKyoku string = "オリジナル曲"
)

// POST
const ()
