package main

import (
	"flag"
	"fmt"

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

	const errno1 = 1
	const errno2 = 2
	const errno3 = 3
	a := [3]string{"no error", "Eio", "invalid argument"}
	b := a
	fmt.Printf("test1 %v; %v\n", a, b)
	fmt.Printf("r %v\n", testString(&a))
	fmt.Printf("test1 again %v; %v\n", a, b)

	testString2(a)
	fmt.Printf("test2 again %v; %v\n", a, b)

	a2 := []string{errno1: "no error", errno2: "Eio", errno3: "invalid argument"}
	fmt.Printf("test2 %v\n", a2)
	m1 := map[int]string{errno1: "no error", errno2: "Eio", errno3: "invalid argument"}
	fmt.Printf("test3 %v\n", m1)
}
