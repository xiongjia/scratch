package util

import (
	"io"
	"log"
	"os"
	"strings"

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

func ParseLogLevel(l string) LogLevel {
	switch {
	case strings.EqualFold(l, "ERROR"):
		return ErrorLevel
	case strings.EqualFold(l, "WARN"):
		return WarnLevel
	case strings.EqualFold(l, "INFO"):
		return InfoLevel
	case strings.EqualFold(l, "DEBUG"):
		return DebugLevel
	case strings.EqualFold(l, "TRACE"):
		return TraceLevel
	default:
		return InfoLevel
	}
}

func LogLevelToStr(l LogLevel) string {
	switch l {
	case ErrorLevel:
		return "ERROR"
	case WarnLevel:
		return "WARN"
	case InfoLevel:
		return "INFO"
	case DebugLevel:
		return "DEBUG"
	case TraceLevel:
		return "TRACE"
	default:
		return "INFO"
	}
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
