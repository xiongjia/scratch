package main

import (
	"stray/internal/log"
)

func main() {
	log.Infof("test %d %d abc", 1, 2)
	log.Infof("test %d %s abc", 1, "34567u8i")
}
