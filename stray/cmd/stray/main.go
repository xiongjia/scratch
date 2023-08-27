package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	fmt.Println("init")
}

func testString(a *[3]string) (r string) {
	r = "test"
	fmt.Printf("a1: %s\n", a[1])
	a[1] = "change1"
	for _, v := range *a {
		fmt.Printf("data: %s\n", v)
	}
	return
}

func testString2(a [3]string) (r string) {
	r = "test"
	fmt.Printf("a1: %s\n", a[1])
	a[1] = "change3"
	for _, v := range a {
		fmt.Printf("data: %s\n", v)
	}
	return
}

type identity struct {
	FirstName string
	Age       int
}

func main() {
	flag.Int("flagname", 1234, "help message for flagname")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	viper.SetConfigName("test.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	i := viper.GetInt("flagname") // retrieve value from viper
	fmt.Printf("test %d\n", i)

	fmt.Println("chan tests")

	errc := make(chan string, 1)
	go func() {
		fmt.Printf("Start sleep\n")
		time.Sleep(time.Second * 100)
		errc <- "hello1"
	}()

	fmt.Printf("Waiting...")
	select {
	case str := <-errc:
		fmt.Println(str)
	}

	fmt.Printf("exit \n")
}
