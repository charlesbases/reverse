package types

// Abbre 缩写
var Abbre = map[string]string{
	"id":   "ID",
	"ip":   "IP",
	"uid":  "UID",
	"uuid": "UUID",
}

// MysqlType2GoType mysql type to go type
var MysqlType2GoType = map[string]string{
	"int":        "int",
	"integer":    "int",
	"tinyint":    "int",
	"smallint":   "int",
	"mediumint":  "int",
	"bit":        "int",
	"bool":       "bool",
	"bigint":     "int64",
	"enum":       "string",
	"set":        "string",
	"varchar":    "string",
	"char":       "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"text":       "string",
	"longtext":   "string",
	"blob":       "string",
	"tinyblob":   "string",
	"mediumblob": "string",
	"longblob":   "string",
	"binary":     "string",
	"varbinary":  "string",
	"json":       "string",
	"float":      "float64",
	"double":     "float64",
	"decimal":    "float64",
	"time":       "time.Time",
	"date":       "time.Time",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
}

// PostgresType2GoType postgres type to go type
var PostgresType2GoType = map[string]string{}
