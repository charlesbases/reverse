package tables

import "regexp"

// MatchingFunc matching table name
type MatchingFunc func(name string) bool

// Options options
type Options struct {
	excludes []MatchingFunc
	includes []MatchingFunc
}

// Option options apply
type Option func(o *Options) ()

// NewOptions returns new Options with Option
func NewOptions(opts ...Option) *Options {
	options := &Options{}
	for _, opt := range opts {
		opt(options)
	}
	return options
}

// Matching returns whether table name meets the matching conditions
func (o *Options) Matching(name string) bool {
	// matching include
	for _, include := range o.includes {
		if include(name) {
			return true
		}
	}

	// matching exclude
	for _, exclude := range o.excludes {
		if exclude(name) {
			return false
		}
	}

	// if includes is nil, it means that all
	return len(o.includes) == 0
}

// WithExclude excluded tablename
func WithExclude(patterns ...string) Option {
	return func(o *Options) {
		for _, pattern := range patterns {
			o.excludes = append(o.excludes, func(name string) bool {
				ok, _ := regexp.MatchString(pattern, name)
				return ok
			})
		}
	}
}

// WithInclude included tablename
func WithInclude(patterns ...string) Option {
	return func(o *Options) {
		for _, pattern := range patterns {
			o.includes = append(o.includes, func(name string) bool {
				ok, _ := regexp.MatchString(pattern, name)
				return ok
			})
		}
	}
}
