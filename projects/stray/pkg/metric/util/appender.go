package util

import "github.com/prometheus/prometheus/storage"

func NewNopAppend() storage.Appendable {
	return &nopAppendable{}
}
