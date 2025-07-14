package converters

import (
	"github.com/charlesbases/reverse/internal/domain/entity/tables"
	"github.com/charlesbases/reverse/internal/domain/repo/converters"
)

// ConverterService converter service
type ConverterService struct {
	converterRepo converters.ConverterRepository
}

// NewConverterService returns ConverterService
func NewConverterService(converter converters.ConverterRepository) *ConverterService {
	return &ConverterService{converterRepo: converter}
}

// ConvertColumnType default sql type change to go types
func (s *ConverterService) ConvertColumnType(t *tables.ColumnTypeEntity) string {
	return s.converterRepo.ConvertColumnType(t)
}

// CamelCase returns name for CamelCase
func (s *ConverterService) CamelCase(name string) string {
	return s.converterRepo.ConvertColumnName(name)
}
