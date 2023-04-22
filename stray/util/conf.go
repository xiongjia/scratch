package util

import (
	"flag"
	"fmt"
)

type Conf struct {
	LogLevel LogLevel
}

func ParseConf(arguments []string) *Conf {
	args := flag.NewFlagSet("stray", flag.ExitOnError)

	var logLevel string
	args.StringVar(&logLevel, "logLevel", "DEBUG", "Log Level (Default is DEBUG)")
	args.Parse(arguments)

	fmt.Printf("Log level %s\n", logLevel)
	return &Conf{
		LogLevel: ParseLogLevel(logLevel),
	}
}
