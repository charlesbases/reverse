package tables

import (
	"context"

	"github.com/charlesbases/reverse/internal/domain/entity/tables"
)

// TableRepository table repository
type TableRepository interface {
	// Tables returns all matching tables
	Tables(ctx context.Context, opts ...Option) ([]*tables.TableEntity, error)
}
