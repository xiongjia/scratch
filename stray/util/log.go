package util

import (
	"log"
	"os"
)

type LogAppenderf func(format string, v ...any)
type LogAppender func(v ...any)

type LogLevel int8

const (
	TraceLevel LogLevel = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
)

type Logger struct {
	Debugf LogAppenderf
	Debug  LogAppender
}

func makeLogWrite(confLevel LogLevel, level LogLevel, prefix string) (LogAppenderf, LogAppender) {
	appender := log.New(os.Stdout, prefix, log.Ldate|log.Ltime|log.Lshortfile)
	return appender.Printf, appender.Print
}

func NewLogger(level LogLevel) *Logger {
	debugAppenderf, debugAppender := makeLogWrite(level, DebugLevel, "DEBUG ")
	logger := &Logger{
		Debugf: debugAppenderf,
		Debug:  debugAppender,
	}
	return logger
}
