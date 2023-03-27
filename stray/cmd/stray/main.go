package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Context struct {
	serverPort int16
}

func welcome(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello %s\n", ps.ByName("name"))
}

func main() {
	ctx := &Context{}
	ctx.serverPort = 8871

	router := httprouter.New()
	router.GET("/hello/:name", welcome)
	fmt.Printf("Server starting at port %d ...\n", ctx.serverPort)
	http.ListenAndServe(fmt.Sprintf(":%d", ctx.serverPort), router)
}
