package main

import "errors"

type Dialector interface {
	Schema() string
	Parse() *Plugin
}

var (
	// ErrInvalidDsn 无效的 dsn 地址
	ErrInvalidDsn = errors.New("invalid dsn")
)

// Options .
type Options struct {
	Dsn string
}

type Option func(o *Options)

// Dsn .
func Dsn(dsn string) Option {
	return func(o *Options) {
		o.Dsn = dsn
	}
}
