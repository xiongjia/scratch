package collection

import (
	"log/slog"

	"github.com/creasty/defaults"
)

func First[T any](collection []T) (T, bool) {
	length := len(collection)
	if length == 0 {
		var t T
		err := defaults.Set(&t)
		if err != nil {
			slog.Warn("set default value error", slog.Any("err", err))
		}
		return t, false
	}
	return collection[0], true
}

func FirstOrEmpty[T any](collection []T) T {
	i, _ := First(collection)
	return i
}
