package util

import (
	"fmt"
)

type Logger interface {
	Debugf(format string)
}

type logger struct {
	name string
}

func (log *logger) Debugf(format string) {
	fmt.Println(log.name)
	fmt.Println(format)
}

func newLogger() Logger {
	l := &logger{name: "log1"}
	return l
}

func Test() {
	fmt.Println("test1")

	log := newLogger()
	log.Debugf("abc")
}
