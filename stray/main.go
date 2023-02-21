package main

import (
	util "stray/util"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	writer := &lumberjack.Logger{
		Filename:   "/Users/xiongjiale/datum/tmp/foo.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
		Compress:   true,
	}
	w := zapcore.AddSync(writer)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.DebugLevel,
	)
	logger2 := zap.New(core)
	logger2.Debug("test1 writer ")
	writer.Write([]byte("Test1 abc\n"))

	logger := zap.NewExample()
	defer logger.Sync()
	logger.Info("test")

	util.Test()
}
