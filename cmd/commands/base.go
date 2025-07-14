package commands

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/charlesbases/reverse/internal/applications"
	"github.com/charlesbases/reverse/internal/domain/repo/tables"
	"github.com/charlesbases/reverse/internal/infras/config"
	"github.com/charlesbases/reverse/internal/providers"
)

func reverse(ctx context.Context, file string, driver string) error {
	// init config
	if err := config.Init(file); err != nil {
		return err
	}

	// init base config
	conf, err := config.InitTargetConfig()
	if err != nil {
		return err
	}

	// init db
	db, err := config.NewDB(driver)
	if err != nil {
		return err
	}

	// mkdir output dir
	if err := os.MkdirAll(conf.OutputDir, os.ModePerm); err != nil {
		return err
	}

	// init repositories
	provider := providers.NewRepositories(conf, db)

	// init services
	services := applications.NewServices(provider.TableRepo, provider.ConverterRepo, provider.GeneratorRepo, provider.WatcherRepo)

	// get all tables
	source, err := services.TableService.Tables(ctx, tables.WithInclude(conf.IncludeTables...), tables.WithExclude(conf.ExcludeTables...))
	if err != nil {
		return err
	}

	funcMap := map[string]interface{}{
		"Date":        func() string { return time.Now().UTC().Format(time.DateTime) },
		"Version":     func() string { return version },
		"PackageName": func() string { return filepath.Base(conf.OutputDir) },
		"CamelCase":   services.ConverterService.CamelCase,
		"GoType":      services.ConverterService.ConvertColumnType,
	}

	// multiple_files
	for _, table := range source {
		fpath := filepath.Join(conf.OutputDir, fmt.Sprintf("%s.go", table.Name))

		// generate code
		if err := services.GeneratorService.Generate(fpath, conf.Template, table, funcMap); err != nil {
			return err
		}

		// format code
		if err := services.WatcherService.Format(ctx, fpath); err != nil {
			return err
		}
	}

	return nil
}
