package time

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTime_MilliToHour(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"should return converted half-hour": func(t *testing.T) {
			hours, _ := MilliToHour(1000 * 60 * 30) // 30 min

			assert.Equal(t, float32(0.5), hours)
		},

		"should return converted 5 hours": func(t *testing.T) {
			hours, _ := MilliToHour(1000 * 60 * 60 * 5)

			assert.Equal(t, float32(5), hours)
		},

		"should return error if negative provided": func(t *testing.T) {
			_, err := MilliToHour(-100)

			if assert.Error(t, err) {
				assert.Equal(t, err, fmt.Errorf("negative time provided"))
			}
		},
	}

	for name, run := range testCases {
		t.Run(name, func(t *testing.T) {
			run(t)
		})
	}
}
