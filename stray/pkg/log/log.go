package log

import (
	"go.uber.org/zap"
)

func GetLogger() zap.Logger {
	return *zap.L()
}

func New() {
	logger := zap.NewExample()
	defer logger.Sync()

}
