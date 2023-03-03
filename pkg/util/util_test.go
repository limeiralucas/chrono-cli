package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GetWeekStartAndEnd(t *testing.T) {
	date := time.Date(2023, 3, 3, 12, 12, 0, 0, time.UTC)
	expectedStart := time.Date(2023, 2, 26, 0, 0, 0, 0, time.UTC)
	expectedEnd := time.Date(2023, 3, 4, 23, 59, 59, 0, time.UTC)

	weekStart, weekEnd := GetWeekStartAndEnd(date)

	assert.Equal(t, expectedStart, weekStart)
	assert.Equal(t, expectedEnd, weekEnd)
}

func Test_AddWeekPlus(t *testing.T) {
	date := time.Date(2023, 3, 3, 12, 14, 0, 0, time.UTC)
	expected := time.Date(2023, 3, 10, 12, 14, 0, 0, time.UTC)

	newDate := AddWeek(date, 1)

	assert.Equal(t, expected, newDate)
}

func Test_AddWeekSubtract(t *testing.T) {
	date := time.Date(2023, 3, 3, 12, 14, 0, 0, time.UTC)
	expected := time.Date(2023, 2, 24, 12, 14, 0, 0, time.UTC)

	newDate := AddWeek(date, -1)

	assert.Equal(t, expected, newDate)
}
