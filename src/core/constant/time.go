package constant

import "time"

func WeekDuration() time.Duration {
	d, _ := time.ParseDuration("168h")
	return d
}

func DayDuration() time.Duration {
	d, _ := time.ParseDuration("24h")
	return d
}
