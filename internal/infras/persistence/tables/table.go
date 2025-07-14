package tables

import (
	"context"
	"strings"

	"xorm.io/xorm"
	"xorm.io/xorm/schemas"

	entity "github.com/charlesbases/reverse/internal/domain/entity/tables"
	repo "github.com/charlesbases/reverse/internal/domain/repo/tables"
)

var _ repo.TableRepository = (*tableImpl)(nil)

type tableImpl struct {
	db *xorm.Engine
}

// NewTableRepository returns repo.TableRepository implement
func NewTableRepository(db *xorm.Engine) repo.TableRepository {
	return &tableImpl{db: db}
}

// Tables returns all matching tables
func (t *tableImpl) Tables(ctx context.Context, opts ...repo.Option) ([]*entity.TableEntity, error) {
	metas, err := t.db.DBMetas()
	if err != nil {
		return nil, err
	}
	if len(metas) == 0 {
		return nil, xorm.ErrNotExist
	}

	options := repo.NewOptions(opts...)
	tables := make([]*entity.TableEntity, 0, len(metas))
	for _, m := range metas {
		if options.Matching(m.Name) {
			tables = append(tables, t.convert(m))
		}
	}

	return tables, nil
}

func (t *tableImpl) convert(source *schemas.Table) *entity.TableEntity {
	columns := source.Columns()

	table := &entity.TableEntity{
		Name:    source.Name,
		Comment: source.Comment,
		Columns: make([]*entity.ColumnsEntity, 0, len(columns)),
	}

	for _, v := range columns {
		table.Columns = append(table.Columns, &entity.ColumnsEntity{
			Name: v.Name,
			ColumnType: entity.ColumnTypeEntity{
				Name:    strings.ToUpper(v.SQLType.Name),
				Length:  v.SQLType.DefaultLength,
				Length2: v.SQLType.DefaultLength2,
			},
			DataType:        t.db.SQLType(v),
			IsPrimaryKey:    v.IsPrimaryKey,
			IsAutoIncrement: v.IsAutoIncrement,
			IsNullable:      v.Nullable,
			Default:         v.Default,
			Comment:         v.Comment,
			Collation:       v.Collation,
		})
	}

	return table
}
