package domain

import "strings"

var searchTerms = []string{
	"歌ってみた",
	"cover",
	"original song",
	"original",
	"covered by",
	"mv",
	"official",
	"オリジナル曲",
}

func searchQuery() string {
	parts := make([]string, len(searchTerms))
	for i, term := range searchTerms {
		if strings.Contains(term, " ") {
			parts[i] = `"` + term + `"`
		} else {
			parts[i] = term
		}
	}
	return strings.Join(parts, "|")
}

func titleRetrieval(title string) bool {
	title = strings.ToLower(title)
	for _, term := range searchTerms {
		if strings.Contains(title, term) {
			return true
		}
	}
	return false
}
