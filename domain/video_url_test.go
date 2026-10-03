package domain

import "testing"

func TestVideoURL(t *testing.T) {
	tests := []struct {
		id   string
		want string
	}{
		{"video-id", "https://www.youtube.com/watch?v=video-id"},
		{"a&b c", "https://www.youtube.com/watch?v=a%26b+c"},
	}
	for _, tt := range tests {
		if got := VideoURL(tt.id); got != tt.want {
			t.Errorf("VideoURL(%q) = %q, want %q", tt.id, got, tt.want)
		}
	}
}
