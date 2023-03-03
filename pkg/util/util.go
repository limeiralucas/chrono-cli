package util

import (
	"time"
)

const dayDuration = time.Duration(time.Hour) * 24
const weekDuration = dayDuration * 7

func GetWeekStartAndEnd(tm time.Time) (time.Time, time.Time) {
	weekday := tm.Weekday()

	year, month, day := tm.Date()
	dayStart := time.Date(year, month, day, 0, 0, 0, 0, tm.Location())

	weekStart := dayStart.Add(-time.Duration(weekday) * dayDuration)
	weekEnd := weekStart.Add(weekDuration - time.Duration(time.Second))

	return weekStart, weekEnd
}

func AddWeek(tm time.Time, week int8) time.Time {
	return tm.Add(time.Duration(week) * weekDuration)
}
