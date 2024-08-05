package todaytime

import (
	"strconv"
	"time"
)

func NowTimeFull() (s string) {
	timeLayout := "2006-01-02 15:04:05"
	s = time.Now().Format(timeLayout)
	return s
}

func NowTime() (s string) {
	now := time.Now()
	hour := strconv.Itoa(now.Hour())     //小时
	minute := strconv.Itoa(now.Minute()) //分钟
	second := strconv.Itoa(now.Second()) //秒
	s = " " + hour + ":" + minute + ":" + second
	return s
}
