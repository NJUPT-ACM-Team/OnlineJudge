package base

import (
	"time"
)

func GetDefaultTime() time.Time {
	const shortForm = "2006-Jan-02"
	t, _ := time.Parse(shortForm, "1970-Jan-02")
	return t
}

func InTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func UnmarshalTime(str string) (time.Time, error) {
	return time.Parse(time.RFC1123, str)
}

func MarshalTime(t time.Time) string {
	return t.UTC().Format(time.RFC1123)
}
