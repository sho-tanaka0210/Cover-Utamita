package domain

import "net/url"

func VideoURL(videoID string) string {
	videoURL := url.URL{Scheme: "https", Host: "www.youtube.com", Path: "/watch"}
	query := videoURL.Query()
	query.Set("v", videoID)
	videoURL.RawQuery = query.Encode()
	return videoURL.String()
}
