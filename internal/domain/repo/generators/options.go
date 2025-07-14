package generators

import "html/template"

// Option apply options
type Option func(o *Options) ()

// Options for generator
type Options struct {
	FuncMap template.FuncMap
}

// NewOptions returns Options with Option
func NewOptions(opts ...Option) *Options {
	options := &Options{FuncMap: template.FuncMap{}}
	for _, opt := range opts {
		opt(options)
	}
	return options
}

// WithFunc add func into FuncMap
func WithFunc(name string, fn interface{}) Option {
	return func(o *Options) {
		o.FuncMap[name] = fn
	}
}
