package xtime

import (
	"github.com/davidscholberg/go-durationfmt"
	"time"
)


// 获取人性化的时间格式
func GetDurationHuman(duration time.Duration) (string,error) {
	durationFormat, err := durationfmt.Format(duration, "%0h:%0m:%0s")
	if err != nil {
		return "",err
	}

	return durationFormat,nil
}
