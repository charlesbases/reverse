package main

import (
	"flag"

	"github.com/charlesbases/reverse/logger"
)

var f = flag.String("f", "", "config file")

func main() {
	flag.Parse()

	if *f == "" {
		logger.Fatal("invalid config file")
	}

	decode(*f).run()
}
