package main

import (
	"bytes"
	"flag"
	"fmt"
	"stray/util"

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

	data := identity{
		FirstName: "test1",
		Age:       11,
	}
	var buf bytes.Buffer
	err := util.Marshal(&buf, data)
	if err != nil {
		fmt.Printf("err %s\n", err)
	}
}
