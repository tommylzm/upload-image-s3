package log

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	// 設定 log 格式
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		DisableSorting:  true,
	})

	level := log.Level(6)
	log.SetLevel(level)
}

// Error 回傳 error level 的 log
func Error(err ...interface{}) {
	log.Error(err...)
}

// Trace 回傳 trace level 的 log
func Trace(err ...interface{}) {
	log.Trace(err...)
}

// Debug 回傳 debug level 的 log
func Debug(err ...interface{}) {
	log.Debug(err...)
}

// Info 回傳 info level 的 log
func Info(err ...interface{}) {
	log.Info(err...)
}

// Warning 回傳 warning level 的 log
func Warning(err ...interface{}) {
	log.Warning(err...)
}

// Fatal 回傳 fatal level 的 log
func Fatal(err ...interface{}) {
	log.Fatal(err...)
}

// Panic 回傳 panic level 的 log
func Panic(err ...interface{}) {
	log.Panic(err...)
}

type Fields = log.Fields

func WithFields(fields Fields) *log.Entry {
	return log.WithFields(fields)
}
