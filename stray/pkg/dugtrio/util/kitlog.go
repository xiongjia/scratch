package util

import (
	"errors"
	"fmt"

	kitlog "github.com/go-kit/log"
	kitlevel "github.com/go-kit/log/level"
)

var (
	errMissingValue = errors.New("(MISSING)")
)

type loggerHandler struct {
	kitlog.Logger
}

func (h *loggerHandler) Log(keyvals ...interface{}) error {
	for i := 0; i < len(keyvals); i += 2 {
		k := keyvals[i]
		var v interface{} = errMissingValue
		if i+1 < len(keyvals) {
			v = keyvals[i+1]
		}

		if k == kitlevel.Key() {
			fmt.Printf(" log level = %s", v)
		}
		kstr := fmt.Sprint(k)
		fmt.Printf("key = %s, val = %v", kstr, v)
	}

	return nil
}

func NewKitLoggerAdapter() kitlog.Logger {
	return kitlevel.NewFilter(&loggerHandler{}, kitlevel.AllowAll())
}

func NewKitLoggerAdapterNop() kitlog.Logger {
	return kitlog.NewNopLogger()
}
