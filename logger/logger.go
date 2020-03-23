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
	Debug(fmt string, args ...interface{})
	Info(fmt string, args ...interface{})
	Warn(fmt string, args ...interface{})
	Error(fmt string, args ...interface{})
}

type defaultLogger struct {
	level LogLevel
}

func (l *defaultLogger) Level(level LogLevel) {
	l.level = level
}

func (l *defaultLogger) Debug(fmt string, args ...interface{}) {
	if l.level <= DebugLevel {
		l.print("debug", fmt, args...)
	}
}

func (l *defaultLogger) Info(fmt string, args ...interface{}) {
	if l.level <= InfoLevel {
		l.print("info", fmt, args...)
	}
}

func (l *defaultLogger) Warn(fmt string, args ...interface{}) {
	if l.level <= WarnLevel {
		l.print("warn", fmt, args...)
	}
}

func (l *defaultLogger) Error(fmt string, args ...interface{}) {
	if l.level <= ErrorLevel {
		l.print("error", fmt, args...)
	}
}

func (l *defaultLogger) print(level string, format string, args ...interface{}) {
	a := make([]interface{}, 0)
	a = append(a, time.Now().Format(time.RFC3339), level)
	a = append(a, args...)
	fmt.Printf("%v [%s] "+format+"\n", a...)
}

var holder Logger

func init() {
	holder = &defaultLogger{}
}

func SetLogger(logger Logger) Logger {
	holder = logger
	return holder
}

func SetLevel(lvl LogLevel) {
	holder.Level(lvl)
}

func Debug(fmt string, args ...interface{}) {
	holder.Debug(fmt, args...)
}

func Info(fmt string, args ...interface{}) {
	holder.Info(fmt, args...)
}

func Warn(fmt string, args ...interface{}) {
	holder.Warn(fmt, args...)
}

func Error(fmt string, args ...interface{}) {
	holder.Error(fmt, args...)
}
