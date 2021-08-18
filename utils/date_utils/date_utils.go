package dateutils

import "time"

const (
	LAYOUT = "2006-01-02T15:04:05Z"
	LAYOUTDB = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now()
}

func GetNowString() string {
	return GetNow().Format(LAYOUT)
}

func GetNowDBString() string {
	return GetNow().Format(LAYOUTDB)
}