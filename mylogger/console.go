package mylogger

import (
	"fmt"
	"time"
)

// 往终端写日志

// ConsoleLogger 日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

// NewConsoleLogger 构造函数
func NewConsoleLogger(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (c ConsoleLogger) enable(logLevel LogLevel) bool {
	return logLevel >= c.Level
}

func (c ConsoleLogger) log(logLevel LogLevel, format string, a ...interface{}) {
	if c.enable(logLevel) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		timeStr := now.Format("2006-01-02 15:04:05")
		logLevelStr := getLogString(logLevel)
		fileName, funcName, line := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", timeStr, logLevelStr, fileName, funcName, line, msg)
	}
}

// Debug ...
func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

// Info ...
func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}

// Warning ...
func (c ConsoleLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)
}

// Error ...
func (c ConsoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}

// Fatal ...
func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
