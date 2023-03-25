package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	addcmd := flag.NewFlagSet("add", flag.ExitOnError)
	a := addcmd.Int("a", 1, "value 1")

	fmt.Printf("args = %v\n", os.Args)
	addcmd.Parse(os.Args[1:])
	fmt.Printf("test a = %d\n", *a)
}
