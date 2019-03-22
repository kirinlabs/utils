package datetime

import "time"

func Year(t ...time.Time) int {
	tmp := time.Now()
	if len(t) > 0 {
		tmp = t[0]
	}
	return tmp.Year()
}

func Month(t ...time.Time) int {
	tmp := time.Now()
	if len(t) > 0 {
		tmp = t[0]
	}
	return int(tmp.Month())
}

func Day(t ...time.Time) int {
	tmp := time.Now()
	if len(t) > 0 {
		tmp = t[0]
	}
	return tmp.Day()
}

func YearDay(t ...time.Time) int {
	tmp := time.Now()
	if len(t) > 0 {
		tmp = t[0]
	}
	return tmp.YearDay()
}