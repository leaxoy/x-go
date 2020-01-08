package xlog

import (
	"log"
)

type Level int

type logger struct {
	l *log.Logger
}

func (l *logger) Print(args ...interface{}) {}

func (l *logger) Printf(format string, args ...interface{}) {}

type Logger interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})
}

func BaseLogger() Logger {
	return &logger{}
}

func GetLogger(name string) Logger {
	return &logger{}
}
