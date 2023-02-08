package gtime

import (
	"log"
	"time"
)

// GetHourUnix 获取指定天数后指定小时的时间戳
func GetHourUnix(day, hour int) int64 {
	t := time.Now()
	t = t.AddDate(0, 0, day)
	muteStartTime := time.Date(t.Year(), t.Month(), t.Day(), hour, 0, 0, 0, time.Local).Unix()
	return muteStartTime
}

// GetNowUnix 获取当前时间戳
func GetNowUnix(isMS bool) int64 {
	if isMS {
		return time.Now().UnixNano() / 1e6
	}
	return time.Now().Unix()
}

// GetNowUnixM 获取当前时间戳分钟级
func GetNowUnixM(isMS bool) int64 {
	t := time.Now()
	if isMS {
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, time.Local).UnixNano() / 1e6
	}
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, time.Local).Unix()
}

// Unix2Time 时间戳转文本 2014-01-07 09:32:12
func Unix2Time(t int64, isMS bool) string {
	if isMS {
		t = t / 1000
	}
	return time.Unix(t, 0).Format("2006-01-02 15:04:05")
}

// Test 测试程序
func Test() {
	log.Printf("获取本日零点时间戳 GetHourUnix(0,0) %d", GetHourUnix(0, 0))
	log.Printf("获取第二日零点时间戳 GetHourUnix(1,0) %d", GetHourUnix(1, 0))
	log.Printf("获取当前时间戳 10位 GetNowUnix(false) %d", GetNowUnix(false))
	log.Printf("获取当前时间戳 13为 GetNowUnix(true) %d", GetNowUnix(true))
	log.Printf("时间戳转文本 Unix2Time(1675696844,false) %s", Unix2Time(1675696844, false))
	log.Printf("时间戳转文本 Unix2Time(1675696844256,true) %s", Unix2Time(1675696844256, true))
}
