package infrastructure

import (
	"testing"
	"time"
)

func TestPreviousDayPeriodUsesJSTCalendarDay(t *testing.T) {
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatal(err)
	}
	now := time.Date(2026, time.August, 10, 15, 30, 0, 0, location)

	publishedAfter, publishedBefore := previousDayPeriod(now, location)
	wantAfter := time.Date(2026, time.August, 9, 0, 0, 0, 0, location)
	wantBefore := time.Date(2026, time.August, 10, 0, 0, 0, 0, location)
	if !publishedAfter.Equal(wantAfter) {
		t.Fatalf("publishedAfter = %s, want %s", publishedAfter, wantAfter)
	}
	if !publishedBefore.Equal(wantBefore) {
		t.Fatalf("publishedBefore = %s, want %s", publishedBefore, wantBefore)
	}
}
