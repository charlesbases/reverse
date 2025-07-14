package tables

import (
	"context"

	entity "github.com/charlesbases/reverse/internal/domain/entity/tables"
	"github.com/charlesbases/reverse/internal/domain/repo/tables"
)

// TableService table service
type TableService struct {
	tableRepo tables.TableRepository
}

// NewTableService returns new TableService
func NewTableService(tableRepo tables.TableRepository) *TableService {
	return &TableService{tableRepo: tableRepo}
}

// Tables returns all matching tables
func (s *TableService) Tables(ctx context.Context, opts ...tables.Option) ([]*entity.TableEntity, error) {
	list, err := s.tableRepo.Tables(ctx)
	if err != nil {
		return nil, err
	}

	options := tables.NewOptions(opts...)
	result := make([]*entity.TableEntity, 0, len(list))
	for _, table := range list {
		if options.Matching(table.Name) {
			result = append(result, table)
		}
	}
	return result, nil
}
