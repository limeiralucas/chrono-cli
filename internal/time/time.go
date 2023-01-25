package time

import "fmt"

func MilliToHour(milliseconds int) (float32, error) {
	if milliseconds < 0 {
		return 0, fmt.Errorf("negative time provided")
	}
	convertedTime := float32(milliseconds) / 1000 / 60 / 60

	return convertedTime, nil
}
