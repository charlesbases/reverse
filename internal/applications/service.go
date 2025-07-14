package applications

import (
	"github.com/charlesbases/reverse/internal/applications/converters"
	"github.com/charlesbases/reverse/internal/applications/generators"
	"github.com/charlesbases/reverse/internal/applications/tables"
	"github.com/charlesbases/reverse/internal/applications/watchers"
	cr "github.com/charlesbases/reverse/internal/domain/repo/converters"
	gr "github.com/charlesbases/reverse/internal/domain/repo/generators"
	tr "github.com/charlesbases/reverse/internal/domain/repo/tables"
	wr "github.com/charlesbases/reverse/internal/domain/repo/watchers"
)

// Services all services
type Services struct {
	TableService     *tables.TableService
	ConverterService *converters.ConverterService
	GeneratorService *generators.GeneratorService
	WatcherService   *watchers.WatcherService
}

// NewServices returns Services
func NewServices(t tr.TableRepository, c cr.ConverterRepository, g gr.GeneratorRepository, w wr.WatcherRepository) *Services {
	tableService := tables.NewTableService(t)
	converterService := converters.NewConverterService(c)
	generatorService := generators.NewGeneratorService(g)
	watcherService := watchers.NewWatcherService(w)

	return &Services{
		TableService:     tableService,
		ConverterService: converterService,
		GeneratorService: generatorService,
		WatcherService:   watcherService,
	}
}
