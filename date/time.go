package date

import (
	"fmt"
	"time"
)

func CurTimeStr() string {
	now := time.Now()
	return fmt.Sprintf("%04d%02d%02d_%02d%02d%02d",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second(),
	)
}

func ParseTimeStr(timeStr string) (time.Time, error) {
	layout := "20060102_150405"
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
