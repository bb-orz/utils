package xtime

import (
	"fmt"
	"time"
)

const (
	timeFormart = "2019-01-01 00:00:00"
)

/**
格式化时间
*/
func FormatTime(time time.Time) (dateTimeFormatted string) {
	dateTimeFormatted = fmt.Sprintf("\"%s\"", time.Local().Format(timeFormart))
	return
}

/**
解析时间字符串
*/
func ParseTime(timeStr string) (*time.Time,error) {
	timeParsed, err := time.ParseInLocation(timeFormart, timeStr, time.Local)
	if err != nil {
		return nil, err
	}

	return &timeParsed,nil
}
