package bulbasaur

import (
	gokitlog "github.com/go-kit/log"
)

type loggerHandler struct {
	gokitlog.Logger
}

func (h *loggerHandler) Log(keyvals ...interface{}) error {
	return nil
}

func makeLogger() *loggerHandler {
	return &loggerHandler{}
}
