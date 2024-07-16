package template

import (
	"fmt"
	"time"
)

func FomatDate(date *time.Time) string {
	if date == nil {
		return ""
	}
	return date.Format("02/01/2006 15:04:05")
}

func MemoryFormat(value int) string {
	var result string
	if value > 1e3 {
		val := float32(value) / 1e3
		result = fmt.Sprintf("%.2f KB", val)
	}
	if value > 1e6 {
		val := float32(value) / 1e6
		result = fmt.Sprintf("%.2f MB", val)
	}
	if value > 1e9 {
		val := float32(value) / 1e9
		result = fmt.Sprintf("%.2f GB", val)
	}
	return result
}
