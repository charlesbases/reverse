package generators

// GeneratorRepository generator repossitory
type GeneratorRepository interface {
	// Generate go files from matching tables according to the template
	Generate(filepath string, text string, data interface{}, opts ...Option) error
}
