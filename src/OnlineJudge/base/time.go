package base

import (
	"time"
)

func GetDefaultTime() time.Time {
	const shortForm = "2006-Jan-02"
	t, _ := time.Parse(shortForm, "1970-Jan-02")
	return t
}
