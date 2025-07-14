package watchers

import (
	"context"
	"os"
	"os/exec"

	"golang.org/x/tools/imports"

	"github.com/charlesbases/reverse/internal/domain/repo/watchers"
)

var _ watchers.WatcherRepository = (*watcherImpl)(nil)

type watcherImpl struct {
}

// NewWatcherRepository returns watchers.WatcherRepository implement
func NewWatcherRepository() watchers.WatcherRepository {
	return &watcherImpl{}
}

// Format format code
func (w *watcherImpl) Format(ctx context.Context, filepath string) error {
	// go mot tidy
	if err := exec.Command("go", "mod", "tidy").Run(); err != nil {
		return err
	}

	// goimports
	formatted, err := imports.Process(filepath, nil, &imports.Options{
		Comments:  true,
		TabIndent: true,
		TabWidth:  4,
	})
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, formatted, 0644)
}
