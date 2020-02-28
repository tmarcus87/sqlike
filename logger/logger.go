package logger

import (
	"fmt"
	"time"
)

type LogLevel int

const (
	DebugLevel LogLevel = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
)

type Logger interface {
	Level(level LogLevel)
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type defaultLogger struct {
	level LogLevel
}

func (l *defaultLogger) Level(level LogLevel) {
	l.level = level
}

func (l *defaultLogger) Debug(msg string) {
	if l.level >= DebugLevel {
		l.print("debug", msg)
	}
}

func (l *defaultLogger) Info(msg string) {
	if l.level >= InfoLevel {
		l.print("info", msg)
	}
}

func (l *defaultLogger) Warn(msg string) {
	if l.level >= WarnLevel {
		l.print("warn", msg)
	}
}

func (l *defaultLogger) Error(msg string) {
	if l.level >= ErrorLevel {
		l.print("error", msg)
	}
}

func (l *defaultLogger) print(level string, msg string) {
	fmt.Printf("%v [%s] %s\n", time.Now().Format(time.RFC3339), level, msg)
}

var holder Logger

func init() {
	holder = &defaultLogger{}
}

func SetLogger(logger Logger) {
	holder = logger
}

func Debug(msg string) {
	holder.Debug(msg)
}

func Info(msg string) {
	holder.Info(msg)
}

func Warn(msg string) {
	holder.Warn(msg)
}

func Error(msg string) {
	holder.Error(msg)
}
