package main

type tag string

func (t *tag) string() string {
	return string("`" + *t + "`")
}

var abbre = map[string]string{
	"id":   "ID",
	"ip":   "IP",
	"uid":  "UID",
	"uuid": "UUID",
}

// mysqltype mysql type to go type
var mysqltype = map[string]string{
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
