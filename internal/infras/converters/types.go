package converters

// enumerates all the database column types
var (
	Bit               = "BIT"
	UnsignedBit       = "UNSIGNED BIT"
	TinyInt           = "TINYINT"
	UnsignedTinyInt   = "UNSIGNED TINYINT"
	SmallInt          = "SMALLINT"
	UnsignedSmallInt  = "UNSIGNED SMALLINT"
	MediumInt         = "MEDIUMINT"
	UnsignedMediumInt = "UNSIGNED MEDIUMINT"
	Int               = "INT"
	UnsignedInt       = "UNSIGNED INT"
	Integer           = "INTEGER"
	BigInt            = "BIGINT"
	UnsignedBigInt    = "UNSIGNED BIGINT"
	Number            = "NUMBER"

	Enum = "ENUM"
	Set  = "SET"

	Char             = "CHAR"
	Varchar          = "VARCHAR"
	VARCHAR2         = "VARCHAR2"
	NChar            = "NCHAR"
	NVarchar         = "NVARCHAR"
	TinyText         = "TINYTEXT"
	Text             = "TEXT"
	NText            = "NTEXT"
	Clob             = "CLOB"
	MediumText       = "MEDIUMTEXT"
	LongText         = "LONGTEXT"
	Uuid             = "UUID"
	UniqueIdentifier = "UNIQUEIDENTIFIER"
	SysName          = "SYSNAME"

	Date          = "DATE"
	DateTime      = "DATETIME"
	SmallDateTime = "SMALLDATETIME"
	Time          = "TIME"
	TimeStamp     = "TIMESTAMP"
	TimeStampz    = "TIMESTAMPZ"
	Year          = "YEAR"

	Decimal    = "DECIMAL"
	Numeric    = "NUMERIC"
	Money      = "MONEY"
	SmallMoney = "SMALLMONEY"

	Real   = "REAL"
	Float  = "FLOAT"
	Double = "DOUBLE"

	Binary     = "BINARY"
	VarBinary  = "VARBINARY"
	TinyBlob   = "TINYBLOB"
	Blob       = "BLOB"
	MediumBlob = "MEDIUMBLOB"
	LongBlob   = "LONGBLOB"
	Bytea      = "BYTEA"

	Bool    = "BOOL"
	Boolean = "BOOLEAN"

	Serial    = "SERIAL"
	BigSerial = "BIGSERIAL"

	Json  = "JSON"
	Jsonb = "JSONB"

	XML   = "XML"
	Array = "ARRAY"
)
