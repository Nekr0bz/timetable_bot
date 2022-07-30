package utils

import "time"

func GetMonday(dayOfWeek time.Time) time.Time {
	tld := dayOfWeek.Weekday() - 1
	if tld < 0 {
		tld = 6
	}
	return dayOfWeek.AddDate(0, 0, int(-tld))
}
