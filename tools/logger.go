package tools

import (
	"github.com/astaxie/beego/logs"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"github.com/uniplaces/carbon"
)

func NewFileLogrus(path string) *logrus.Logger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(&lumberjack.Logger{
		Filename:   path,
		MaxSize:    500, // megabytes
		MaxBackups: 30,
		MaxAge:     365,   //days
		Compress:   false, // disabled by default
	})
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: carbon.DefaultFormat,
	})
	return logger
}

func UseDefaultLogrus() {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: carbon.DefaultFormat,
	})
}

func UseDefaultLogs() {
	logs.SetLogFuncCall(true)
	logs.SetLogFuncCallDepth(3)
	logs.SetLevel(logs.LevelDebug)
}
