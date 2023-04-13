package util

import (
	"io"
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
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

func makeAppender(confLevel LogLevel, level LogLevel, prefix string, w io.Writer) (LogAppenderf, LogAppender) {
	appender := log.New(os.Stdout, prefix, log.Ldate|log.Ltime|log.Lshortfile)
	appender.SetOutput(w)
	return appender.Printf, appender.Print
}

func makeWriter(f string) io.Writer {
	if f == "" {
		return os.Stdout
	} else {
		return &lumberjack.Logger{
			Filename:   f,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		}
	}
}

func NewLogger(filename string, level LogLevel) *Logger {
	writer := makeWriter(filename)
	logger := &Logger{}
	logger.Debugf, logger.Debug = makeAppender(level, DebugLevel, "DEBUG ", writer)
	return logger
}
