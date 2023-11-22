package log

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level = zapcore.Level

const (
	DebugLevel  Level = zap.DebugLevel  // -1
	InfoLevel   Level = zap.InfoLevel   // 0, default level
	WarnLevel   Level = zap.WarnLevel   // 1
	ErrorLevel  Level = zap.ErrorLevel  // 2
	DPanicLevel Level = zap.DPanicLevel // 3, used in development log
	PanicLevel  Level = zap.PanicLevel  // 4
	FatalLevel  Level = zap.FatalLevel  // 5, FatalLevel logs a message, then calls os.Exit(1).
)

type Logger struct {
	l     *zap.Logger // zap ensure that zap.Logger is safe for concurrent use
	level Level
}

var (
	logger *Logger
)

func NewLogger(writer io.Writer, level Level) *Logger {
	if writer == nil {
		panic("the Logger writer is nil")
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewCore(encoder,
		zapcore.AddSync(writer),
		zapcore.Level(level),
	)
	logger := &Logger{
		l:     zap.New(core),
		level: level,
	}
	return logger
}

func init() {
	logger = NewLogger(os.Stdout, InfoLevel)
}

func Debugf(template string, args ...interface{}) {
	logger.l.Sugar().Debugf(template, args)
}

func Infof(template string, args ...interface{}) {
	logger.l.Sugar().Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	logger.l.Sugar().Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	logger.l.Sugar().Errorf(template, args...)
}

func DPanicf(template string, args ...interface{}) {
	logger.l.Sugar().DPanicf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	logger.l.Sugar().Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	logger.l.Sugar().Fatalf(template, args...)
}
