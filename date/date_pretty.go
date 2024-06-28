package date

import (
	"fmt"
	"time"
)

func ToPrettyString(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

func CurrentCustomTimeString() string {
	now := time.Now()
	return fmt.Sprintf("%04d%02d%02d_%02d%02d%02d",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second(),
	)
}

func ParseCustomTimeString(timeStr string) (time.Time, error) {
	layout := "20060102_150405"
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
