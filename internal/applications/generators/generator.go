package generators

import "github.com/charlesbases/reverse/internal/domain/repo/generators"

// GeneratorService generator service
type GeneratorService struct {
	generatorRepo generators.GeneratorRepository
}

// NewGeneratorService returns GeneratorService
func NewGeneratorService(generatorRepo generators.GeneratorRepository) *GeneratorService {
	return &GeneratorService{generatorRepo: generatorRepo}
}

// Generate go files from matching tables according to the template
func (s *GeneratorService) Generate(filepath string, text string, data interface{}, funcMap map[string]interface{}) error {
	var options []generators.Option
	if funcMap != nil && len(funcMap) != 0 {
		options = make([]generators.Option, 0, len(funcMap))
		for k, v := range funcMap {
			options = append(options, generators.WithFunc(k, v))
		}
	}
	return s.generatorRepo.Generate(filepath, text, data, options...)
}
