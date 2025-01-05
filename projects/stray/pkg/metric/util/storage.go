package util

import (
	"github.com/prometheus/prometheus/storage"
)

func NewNotReadyAppender() storage.Appender {
	return notReadyAppender{}
}
