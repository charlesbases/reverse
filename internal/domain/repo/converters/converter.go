package converters

import (
	"github.com/charlesbases/reverse/internal/domain/entity/tables"
)

// ConverterRepository converter repository
type ConverterRepository interface {
	// ConvertColumnType default sql type change to go types
	ConvertColumnType(t *tables.ColumnTypeEntity) string
	// ConvertColumnName column name change to go struct field name
	ConvertColumnName(name string) string
}
