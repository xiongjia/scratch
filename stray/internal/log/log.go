package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func init() {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	var opts []zap.Option
	opts = append(opts, zap.AddCaller())
	opts = append(opts, zap.AddCallerSkip(1))
	opts = append(opts, zap.AddStacktrace(zap.DPanicLevel))
	atmLevel := zap.NewAtomicLevel()
	core := zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), atmLevel)
	logger := zap.New(core, opts...)
	Logger = logger.Sugar()
}
