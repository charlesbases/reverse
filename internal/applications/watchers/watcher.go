package watchers

import (
	"context"

	"github.com/charlesbases/reverse/internal/domain/repo/watchers"
)

// WatcherService watcher service
type WatcherService struct {
	watcherRepo watchers.WatcherRepository
}

// NewWatcherService returns WatcherService
func NewWatcherService(watcherRepo watchers.WatcherRepository) *WatcherService {
	return &WatcherService{watcherRepo: watcherRepo}
}

// Format format code
func (s *WatcherService) Format(ctx context.Context, filepath string) error {
	return s.watcherRepo.Format(ctx, filepath)
}
