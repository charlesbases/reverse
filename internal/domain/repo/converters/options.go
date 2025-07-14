package converters

// Option optins apply func
type Option func(o *Options) ()

// Options for converter
type Options struct {
	acronyms []string
}

// Matching is name an acronym?
func (o *Options) Matching(name string) bool {
	for _, v := range o.acronyms {
		if name == v {
			return true
		}
	}
	return false
}

// NewOptions returns Options with Option
func NewOptions(opts ...Option) *Options {
	options := &Options{acronyms: make([]string, 0)}
	for _, opt := range opts {
		opt(options)
	}
	return options
}

// WithAcronyms convert the matching abbreviation to uppercase
func WithAcronyms(v ...string) Option {
	return func(o *Options) {
		o.acronyms = append(o.acronyms, v...)
	}
}
