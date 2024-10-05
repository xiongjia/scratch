package util

import (
	"log/slog"
	"net/http"
)

func HttpUtilWriteObject[T any](w http.ResponseWriter, statusCode int, obj T) error {
	w.WriteHeader(statusCode)
	return JsonWriterFromObject(w, obj)
}

func HttpUtilWriteObjectQuite[T any](w http.ResponseWriter, statusCode int, obj T) {
	err := HttpUtilWriteObject(w, statusCode, obj)
	if err != nil {
		slog.Error("Write HTTP Response error",
			slog.Any("err", err), slog.Int("status", statusCode), slog.Any("response", obj))
	}
}

func HttpUtilWriteError(w http.ResponseWriter, statusCode int, errObj error) {
	http.Error(w, errObj.Error(), statusCode)
}
