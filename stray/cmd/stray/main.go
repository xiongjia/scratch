package main

import (
	"os"

	"stray/util"
)

func main() {
	conf := util.ParseConf(os.Args[1:])
	logger := util.NewLogger("", conf.LogLevel)

	logger.Debugf("Log level = %s", util.LogLevelToStr(conf.LogLevel))
}
