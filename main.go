package main

import (
	"flag"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/charlesbases/reverse/dialer"
	"github.com/charlesbases/reverse/logger"
	"github.com/charlesbases/reverse/parser"
	"github.com/charlesbases/reverse/types"
)

var f = flag.String("f", "reverse.toml", "config file")

func init() {
	flag.Parse()
}

func main() {
	run(func(opts *types.Options) {
		opts.Type = strings.ToLower(strings.TrimSpace(opts.Type))

		switch opts.Type {
		case "mysql":
			parser.Run(dialer.Mysql(opts))
		case "postgres":
			parser.Run(dialer.Postgres(opts))
		}
	})
}

// run .
func run(fn func(opts *types.Options)) {
	abspath, err := filepath.Abs(*f)
	if err != nil {
		logger.Fatal(err)
	}

	var opts = types.DefaultOption()
	if _, err := toml.DecodeFile(abspath, opts); err != nil {
		logger.Fatal(err)
	}

	fn(opts)
}
