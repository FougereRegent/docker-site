package helper

import (
	"fmt"
)

func OctalToStringFormat(size int64) string {
	var result string
	dataSize := float64(size)

	if res := dataSize / 10e3; res > 1 && res < 1000 {
		result = fmt.Sprintf("%.2f Ko", res)
	} else if res := dataSize / 10e6; res > 1 && res < 1000 {
		result = fmt.Sprintf("%.2f Mo", res)
	} else if res := dataSize / 10e9; res > 1 {
		result = fmt.Sprintf("%.2f Go", res)
	} else {
		result = fmt.Sprintf("%d O", size)
	}

	return result
}
