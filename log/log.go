package glog

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"time"
	gtime "yxProject/time"
)

var pLog zerolog.Logger
var pLogFile *os.File
var pTime int64
var pType int

/*
	日志输出，同时打印日志到文件与控制台
	日志格式：{"l":"info","t":"02-06 21:24:22","msg":"测试"}
	调用方式：glog.Log().Info().Msg("测试")
	项目地址：https://github.com/rs/zerolog
	参考源码：
	http://www.zengyuzhao.com/archives/211
	https://blog.csdn.net/geekqian/article/details/125942407
*/
// 初始化日志性能型
func _iniLogStdout() {
	//设置zerolog全局设置
	zerolog.TimestampFieldName = "t"
	zerolog.LevelFieldName = "l"
	zerolog.MessageFieldName = "msg"
	zerolog.TimeFieldFormat = "01-02 15:04:05"
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	multi := zerolog.MultiLevelWriter(os.Stdout, pLogFile)
	pLog = zerolog.New(multi).With().Timestamp().Logger()

}

// 初始化日志美化型
func _iniLogConsoleWriter() {
	//03-01 23:26:53 | INFO  | 1 id:11630865529224441856;
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "01-02 15:04:05"}
	consoleWriter.FormatLevel = func(i interface{}) string { //日志等级| INFO  |
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	consoleWriter.FormatMessage = func(i interface{}) string { //消息内容
		return fmt.Sprintf("%s", i)
	}
	consoleWriter.FormatFieldName = func(i interface{}) string { //key
		return fmt.Sprintf("%s:", i)
	}
	consoleWriter.FormatFieldValue = func(i interface{}) string { //value
		return fmt.Sprintf("%s;", i)
	}
	multi := zerolog.MultiLevelWriter(consoleWriter, pLogFile)
	pLog = zerolog.New(multi).With().Timestamp().Logger()
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
	pTime = gtime.GetHourUnix(0, 0)
}

// Log 获得实例对象调用
func Log() *zerolog.Logger {
	if pTime != gtime.GetHourUnix(0, 0) {
		pLogFile.Close()
		IniLog(pType)
	}
	return &pLog
}
func IniLog(t int) {
	pType = t
	creadLogFile()
	if pType == 0 {
		_iniLogStdout()
	} else {
		_iniLogConsoleWriter()
	}
}

// Test 测试
func Test() {
	pLog.Print("hello world")
	pLog.Info().Msg("测试")
	pLog.Log().Msg("hello,world")
}
