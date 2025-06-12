package utils

import (
	"fmt"
	"time"
)

func IsWeekend(t time.Time) bool {
	weekday := t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

func GetCurrentTime() time.Time {
	return time.Now()
}

func ConvertStringToDate(date string) time.Time {
	layout := "2006-01-02"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error:", err)
		return time.Time{}
	}

	return t
}
