package utils

import "time"

func IsWeekend(t time.Time) bool {
	weekday := t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

func GetCurrentTime() time.Time {
	return time.Now()
}
