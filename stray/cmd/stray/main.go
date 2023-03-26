package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	appName = "stray"
)

type Context struct {
	logFilename string
}

func welcome(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome!")
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

	mux := http.NewServeMux()
	mux.HandleFunc("/welcome", welcome)
	http.ListenAndServe(":8771", mux)
}
