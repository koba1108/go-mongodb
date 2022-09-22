package helper

import "time"

func ParseTimeFromStr(str, format string) (time.Time, error) {
	if format == "" {
		format = "2006-01-02"
	}
	return time.Parse(format, str)
}
