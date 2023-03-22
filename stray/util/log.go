package util

import (
	"fmt"
	"io"
	"os"
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

type Options struct {
	enable       bool
	enableStdout bool
	level        Level

	logFilename     string
	logFileMaxSize  int
	rotatMaxSize    int
	rotatMaxBackups int
	rotateMaxDays   int
	rotateCompress  bool
}

type logger struct {
	level       Level
	fileWriter  io.Writer
	printWriter io.Writer
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
		return fmt.Sprintf("LEVEL(%d)", l)
	}
}

func (log *logger) Debugf(format string, args ...interface{}) {
	log.write(DebugLevel, format, args...)
}

func (log *logger) write(level Level, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	if log.printWriter != nil {
		log.printWriter.Write([]byte(msg))
		log.printWriter.Write([]byte("\n"))
	}
}

func newLogger(opts *Options) Logger {
	log := &logger{
		level:      opts.level,
		fileWriter: nil,
		printWriter: func() io.Writer {
			if opts.enableStdout {
				return os.Stdout
			} else {
				return nil
			}
		}(),
	}
	return log
}

func Test() {
	opts := &Options{}
	fmt.Printf("%v\n", opts.enable)

	log := newLogger(&Options{
		enable:       true,
		level:        DebugLevel,
		enableStdout: true,
	})
	log.Debugf("Test log %s %s %s", "1", "2", "3")
}
