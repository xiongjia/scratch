package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net/http"
	"net/http/httptest"

	"stray/util"

	"github.com/julienschmidt/httprouter"
)

type Context struct {
	serverPort int16
	test       func()
}

func welcome(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Printf("request test\n")
	age := r.URL.Query().Get("age")
	fmt.Fprintf(w, "hello %s - %s\n", ps.ByName("name"), age)
}

func test1() {
	fmt.Println("test func1")
}

func main() {
	logger := util.NewLogger(util.DebugLevel)
	logger.Debug("Debug 1")
	logger.Debugf("Debug %d", 999)

	dbgData := &struct {
		a byte
		b byte
		c int32
	}{1, 1, 3}

	fmt.Printf("Test %#v\n", dbgData)

	dbgBuf := &bytes.Buffer{}
	err := binary.Write(dbgBuf, binary.BigEndian, dbgData)
	if err != nil {
		panic(err)
	}
	fmt.Println(dbgBuf.Bytes())

	fmt.Println(`abc"abc"\n`)

	data1 := [...]int{0, 1, 2}
	s1 := copy(data1[:1], data1[:2])
	data1[0] = 100
	fmt.Printf("data1 %v, s1 %v\n", data1, s1)

	ctx := &Context{
		serverPort: 8871,
		test:       test1,
	}
	ctx.test()

	router := httprouter.New()
	router.GET("/hello/:name", welcome)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello/ut?age=1", nil)
	router.ServeHTTP(w, req)
	fmt.Printf("code = %d, %s", w.Code, w.Body.String())

	fmt.Printf("Server starting at port %d ...\n", ctx.serverPort)
	http.ListenAndServe(fmt.Sprintf(":%d", ctx.serverPort), router)
}
