package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
  say("direct")

	go say("routines")
  say("direct2")

  go func(msg string) {
    fmt.Println(msg)
  }("going")

  go func() {
    say("test2")
  }()

  fmt.Scanln()
	fmt.Println("done")
}

