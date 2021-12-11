package futil

import (
	"funtester/base"
	"log"
	"time"
)

func Nano() int64 {
	return time.Now().UnixNano()
}

func Milli() int64 {
	return time.Now().UnixMilli()
}

func Second() int {
	return time.Now().Second()
}

func Sleep(second, milli int) {
	time.Sleep(time.Duration(second)*time.Second + time.Duration(milli)*time.Millisecond)
}

func Date() string {
	return time.Now().Format(base.Format)
}

func ToDate(t int64) string {
	return time.UnixMilli(t).Format(base.Format)
}

func ToTime(t string) time.Time {
	location, err := time.Parse(base.Format, t)
	if err != nil {
		log.Println("时间转换出错了!")
	}
	return location
}

func Weekday() int {
	return int(time.Now().Weekday())
}

func Day() int {
	return int(time.Now().Day())
}
