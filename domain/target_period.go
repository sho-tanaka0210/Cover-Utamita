package domain

import "time"

const targetDateLayout = "2006-01-02"

func PreviousDayPeriod(now time.Time) (time.Time, time.Time) {
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		location = time.FixedZone("JST", 9*60*60)
	}
	today := now.In(location)
	end := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, location)
	return end.AddDate(0, 0, -1), end
}

func PreviousDayDate(now time.Time) string {
	start, _ := PreviousDayPeriod(now)
	return start.Format(targetDateLayout)
}
