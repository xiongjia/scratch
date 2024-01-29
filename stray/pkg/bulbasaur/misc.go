package bulbasaur

import (
	kitlog "github.com/go-kit/log"
	kitlevel "github.com/go-kit/log/level"
)

type loggerHandler struct {
	kitlog.Logger
}

func (h *loggerHandler) Log(keyvals ...interface{}) error {

	return nil
}

func makeLogger() kitlog.Logger {
	return kitlevel.NewFilter(&loggerHandler{}, kitlevel.AllowAll())
}
