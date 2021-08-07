package dateutils

import "time"

const (
	LAYOUT = "2006-01-02T15:04:05Z"
)

func GetNow() time.Time {
	return time.Now()
}

func GetNowString() string {
	return GetNow().Format(LAYOUT)
}
