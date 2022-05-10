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
var PostgresType2GoType = map[string]string{
	"int2":     "int", // 2 字节整数
	"smallint": "int", // 2 字节整数

	"int4":    "int", // 4 字节整数
	"integer": "int", // 4 字节整数

	"int8":   "int64", // 8 字节整数
	"bigint": "int64", // 8 字节整数

	"smallserial": "int",   // int2 自增
	"serial":      "int",   // int4 自增
	"bigserial":   "int64", // int8 自增

	"real":             "float64",
	"float":            "float64",
	"decimal":          "float64",
	"numeric":          "float64",
	"double precision": "float64",

	"varchar": "string",
	"char":    "string",
	"text":    "string",

	"bool":    "bool",
	"boolean": "bool",

	"json":  "string",
	"jsonb": "string",

	// time
	"timestamp": "time.Time",
	"date":      "string",
	"time":      "string",

	// ip
	"cidr":     "string",
	"inet":     "string",
	"macaddr":  "string",
	"macaddr8": "string",
}
