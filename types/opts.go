package types

// Options .
type Options struct {
	Type string
	Dir  string
	Dsn  string
}

// DefaultOption .
func DefaultOption() *Options {
	return &Options{
		Dir: "models",
	}
}
