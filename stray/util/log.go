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
	Tracef LogAppenderf
	Trace  LogAppender
	Debugf LogAppenderf
	Debug  LogAppender
	Infof  LogAppenderf
	Info   LogAppender
	Warnf  LogAppenderf
	Warn   LogAppender
	Errorf LogAppenderf
	Error  LogAppender
}

func makeAppender(confLevel LogLevel, level LogLevel, prefix string, w io.Writer) (LogAppenderf, LogAppender) {
	if level >= confLevel {
		appender := log.New(os.Stdout, prefix, log.Ldate|log.Ltime|log.Lshortfile)
		appender.SetOutput(w)
		return appender.Printf, appender.Print
	} else {
		appenderPrintf := func(format string, v ...any) {
			// NOP
		}
		appenderPrint := func(v ...any) {
			// NOP
		}
		return appenderPrintf, appenderPrint
	}
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
	logger.Tracef, logger.Trace = makeAppender(level, TraceLevel, "TRACE ", writer)
	logger.Debugf, logger.Debug = makeAppender(level, DebugLevel, "DEBUG ", writer)
	logger.Infof, logger.Info = makeAppender(level, InfoLevel, "INFO ", writer)
	logger.Warnf, logger.Warn = makeAppender(level, WarnLevel, "WARN ", writer)
	logger.Errorf, logger.Error = makeAppender(level, ErrorLevel, "ERROR ", writer)
	return logger
}
