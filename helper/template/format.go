package template

import (
	"time"
)

func FomatDate(date *time.Time) string {
	if date == nil {
		return ""
	}
	return date.Format("02/01/2006 15:04:05")
}
