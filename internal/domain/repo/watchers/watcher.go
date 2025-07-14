package watchers

import (
	"context"
)

// WatcherRepository file watcher repository
type WatcherRepository interface {
	// Format format code
	Format(ctx context.Context, filepath string) error
}
