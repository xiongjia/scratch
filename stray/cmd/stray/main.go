package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	appName = "stray"
)

type Context struct {
	logFilename string
}

func main() {
	fmt.Printf("App: %s\n", appName)

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("working dir: %s\n", cwd)

	addcmd := flag.NewFlagSet("add", flag.ExitOnError)
	a := addcmd.Int("a", 1, "value 1")

	fmt.Printf("args = %v\n", os.Args)
	addcmd.Parse(os.Args[1:])
	fmt.Printf("test a = %d\n", *a)
}
