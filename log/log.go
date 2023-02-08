package glog

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"time"
	gtime "yxProject/time"
)

var pLog zerolog.Logger
var pLogFile *os.File
var pTime int64

/*
	日志输出，同时打印日志到文件与控制台
	日志格式：{"l":"info","t":"02-06 21:24:22","msg":"测试"}
	调用方式：glog.Log().Info().Msg("测试")
	项目地址：https://github.com/rs/zerolog
	参考源码：
	http://www.zengyuzhao.com/archives/211
	https://blog.csdn.net/geekqian/article/details/125942407
*/
//初始化日志
func init() {
	iniLog()
}

// 初始化日志
func iniLog() {
	creadLogFile()
	//设置zerolog全局设置
	zerolog.TimestampFieldName = "t"
	zerolog.LevelFieldName = "l"
	zerolog.MessageFieldName = "msg"
	zerolog.TimeFieldFormat = "01-02 15:04:05"
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	multi := zerolog.MultiLevelWriter(os.Stdout, pLogFile)
	pLog = zerolog.New(multi).With().Timestamp().Logger()
	pTime = gtime.GetHourUnix(0, 0)
}

// 创建日志文件
func creadLogFile() {
	logDir := "./run_log/"
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		fmt.Println("创建日志目录失败:", err)
		return
	}
	fileName := logDir + time.Now().Format("2006-01-02") + ".log"
	pLogFile, _ = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
}

// Log 获得实例对象调用
func Log() *zerolog.Logger {
	if pTime != gtime.GetHourUnix(0, 0) {
		pLogFile.Close()
		iniLog()
	}
	return &pLog
}

// Test 测试
func Test() {
	pLog.Print("hello world")
	pLog.Info().Msg("测试")
	pLog.Log().Msg("hello,world")
}
