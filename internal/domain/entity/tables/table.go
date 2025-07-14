package tables

// TableEntity table information
type TableEntity struct {
	Name    string
	Comment string
	Columns []*ColumnsEntity
}

// ColumnsEntity column in TableEntity
type ColumnsEntity struct {
	Name            string
	ColumnType      ColumnTypeEntity
	DataType        string
	IsPrimaryKey    bool
	IsAutoIncrement bool
	IsNullable      bool
	Default         string
	Comment         string
	Collation       string
}

// ColumnTypeEntity SQLType for column
type ColumnTypeEntity struct {
	Name    string
	Length  int64
	Length2 int64
}
