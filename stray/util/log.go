package util

import (
	"fmt"
	"strconv"
)

type Level int8

const (
	TraceLevel Level = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

type Logger interface {
	Debugf(format string, args ...interface{})
}

type logger struct {
	name string
}

func (l Level) String() string {
	switch l {
	case TraceLevel:
		return "TRACE"
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case PanicLevel:
		return "PANIC"
	case FatalLevel:
		return "FATAL"
	default:
		return strconv.Itoa(int(l))
	}
}

func (log *logger) Debugf(format string, args ...interface{}) {
	fmt.Println(log.name)
	log.write(format, args...)
}

func (log *logger) write(format string, args ...interface{}) {
	fmt.Println(format)
	for _, arg := range args {
		switch arg.(type) {
		case string:
			fmt.Printf("string = %v\n", arg)
		}
	}
}

func newLogger() Logger {
	l := &logger{name: "log1"}
	return l
}

func Test() {
	fmt.Println("test1")

	fmt.Printf("test data = %d, %v\n", TraceLevel, DebugLevel.String())

	log := newLogger()
	log.Debugf("abc", "1", "2", "3")
}
