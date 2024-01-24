package util

import "io"

type dummyWriter struct{}

func (w *dummyWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func makeDummyWriter() io.Writer {
	return &dummyWriter{}
}
