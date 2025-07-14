package providers

import (
	"xorm.io/xorm"

	"github.com/charlesbases/reverse/internal/domain/repo/converters"
	"github.com/charlesbases/reverse/internal/domain/repo/generators"
	"github.com/charlesbases/reverse/internal/domain/repo/tables"
	"github.com/charlesbases/reverse/internal/domain/repo/watchers"
	"github.com/charlesbases/reverse/internal/infras/config"
	ci "github.com/charlesbases/reverse/internal/infras/converters"
	gi "github.com/charlesbases/reverse/internal/infras/generators"
	ti "github.com/charlesbases/reverse/internal/infras/persistence/tables"
	wi "github.com/charlesbases/reverse/internal/infras/watchers"
)

// Repositories all repository
type Repositories struct {
	TableRepo     tables.TableRepository
	ConverterRepo converters.ConverterRepository
	GeneratorRepo generators.GeneratorRepository
	WatcherRepo   watchers.WatcherRepository
}

// NewRepositories returns Repositories
func NewRepositories(conf *config.TargetConfig, db *xorm.Engine) *Repositories {
	tableRepo := ti.NewTableRepository(db)
	converterRepo := ci.NewConverterRepository(converters.WithAcronyms(conf.Acronyms...))
	generatorRepo := gi.NewGeneratorRepository()
	watcherRepo := wi.NewWatcherRepository()

	return &Repositories{
		TableRepo:     tableRepo,
		ConverterRepo: converterRepo,
		GeneratorRepo: generatorRepo,
		WatcherRepo:   watcherRepo,
	}
}
