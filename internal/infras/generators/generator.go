package generators

import (
	"html/template"
	"os"

	"github.com/charlesbases/reverse/internal/domain/repo/generators"
)

var _ generators.GeneratorRepository = (*generatorImpl)(nil)

type generatorImpl struct {
}

// NewGeneratorRepository returns generator.GeneratorRepository implement
func NewGeneratorRepository() generators.GeneratorRepository {
	return &generatorImpl{}
}

// Generate go files from matching tables according to the template
func (g *generatorImpl) Generate(filepath string, text string, data interface{}, opts ...generators.Option) error {
	options := generators.NewOptions(opts...)

	t := template.New(filepath)
	t.Funcs(options.FuncMap)

	if _, err := t.Parse(text); err != nil {
		return err
	}

	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return t.Execute(file, data)
}
