package converters

import (
	"strings"
	"unicode"

	"github.com/charlesbases/reverse/internal/domain/entity/tables"
	"github.com/charlesbases/reverse/internal/domain/repo/converters"
)

var _ converters.ConverterRepository = (*converterImpl)(nil)

type converterImpl struct {
	options *converters.Options
}

// NewConverterRepository returns converter.ConverterRepository implement
func NewConverterRepository(opts ...converters.Option) converters.ConverterRepository {
	return &converterImpl{options: converters.NewOptions(opts...)}
}

// ConvertColumnType default sql type change to go types
func (c *converterImpl) ConvertColumnType(t *tables.ColumnTypeEntity) string {
	switch t.Name {
	case TinyInt, UnsignedTinyInt:
		return "int8"
	case Bit, SmallInt, MediumInt, Int, Integer, Serial:
		return "int"
	case BigInt, BigSerial:
		return "int64"
	case UnsignedBit, UnsignedSmallInt, UnsignedMediumInt, UnsignedInt:
		return "int"
	case UnsignedBigInt:
		return "int64"
	case Float, Real:
		return "float32"
	case Double:
		return "float64"
	case Char, NChar, Varchar, NVarchar, TinyText, Text, NText, MediumText, LongText, Enum, Set, UUID, Clob, SysName:
		return "string"
	case TinyBlob, Blob, LongBlob, Bytea, Binary, MediumBlob, VarBinary, UniqueIdentifier:
		return "byte"
	case Bool:
		return "bool"
	case DateTime, Date, Time, TimeStamp, TimeStampz, SmallDateTime, Year:
		return "time.Time"
	case Decimal, Numeric, Money, SmallMoney:
		return "string"
	default:
		return "string"
	}
}

// ConvertColumnName column name change to go struct field name
func (c *converterImpl) ConvertColumnName(name string) string {
	parts := strings.Split(name, "_")
	for i, part := range parts {
		if len(part) == 0 {
			continue
		}

		// acronyms
		if c.options.Matching(part) {
			parts[i] = strings.ToUpper(part)
			continue
		}

		// camel case
		runes := []rune(part)
		runes[0] = unicode.ToUpper(runes[0])
		parts[i] = string(runes)
	}
	return strings.Join(parts, "")
}
