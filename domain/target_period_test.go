package domain

import (
	"testing"
	"time"
)

func TestPreviousDayPeriodUsesJSTCalendarDay(t *testing.T) {
	tests := []struct {
		name  string
		now   time.Time
		start string
		end   string
	}{
		{"before JST midnight", time.Date(2026, time.August, 9, 14, 59, 0, 0, time.UTC), "2026-08-08T00:00:00+09:00", "2026-08-09T00:00:00+09:00"},
		{"after JST midnight", time.Date(2026, time.August, 9, 15, 0, 0, 0, time.UTC), "2026-08-09T00:00:00+09:00", "2026-08-10T00:00:00+09:00"},
		{"year boundary", time.Date(2026, time.December, 31, 16, 0, 0, 0, time.UTC), "2026-12-31T00:00:00+09:00", "2027-01-01T00:00:00+09:00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start, end := PreviousDayPeriod(tt.now)
			if got := start.Format(time.RFC3339); got != tt.start {
				t.Errorf("start = %q, want %q", got, tt.start)
			}
			if got := end.Format(time.RFC3339); got != tt.end {
				t.Errorf("end = %q, want %q", got, tt.end)
			}
			if got := PreviousDayDate(tt.now); got != start.Format("2006-01-02") {
				t.Errorf("target date = %q, want %q", got, start.Format("2006-01-02"))
			}
		})
	}
}
