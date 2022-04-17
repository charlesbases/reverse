package dialect

import (
	"os"
	"path"
	"sort"
	"strings"

	"github.com/charlesbases/reverse/logger"
	"github.com/charlesbases/reverse/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	mysqlSchemaTables  = "information_schema.TABLES"
	mysqlSchemaColumns = "information_schema.COLUMNS"
)

// mysqlDialect .
type mysqlDialect struct {
	opts   *types.Options
	db     *gorm.DB
	schema string
}

// Mysql .
func Mysql(opts *types.Options) Dialect {
	d := &mysqlDialect{opts: opts}
	d.connect()
	return d
}

// connect .
func (d *mysqlDialect) connect() {
	d.db = open(mysql.Open(d.opts.Dsn))
	d.schema = path.Base(d.opts.Dsn)
}

func (d *mysqlDialect) Options() *types.Options {
	return d.opts
}

func (d *mysqlDialect) Schema() string {
	return d.schema
}

func (d *mysqlDialect) ParseColumnTag(tf *TableColumn) types.Tag {
	var tag types.Tag = make([]string, 0)

	// json
	tag.Append("json", tf.ColumnName)

	// orm
	var ormtag types.TagType = make([]string, 0)
	ormtag.Append("column", tf.ColumnName)
	ormtag.Append("type", tf.ColumnType)
	if tf.IsNull == "NO" {
		ormtag.Append("not null")
	}
	if tf.ColumnKey == "PRI" {
		ormtag.Append("primary_key")
	}
	if tf.Extra == "auto_increment" {
		ormtag.Append("auto_increment")
	}

	return tag
}

func (d *mysqlDialect) ParseColumnType(tf *TableColumn) string {
	var gotype = types.MysqlType2GoType[tf.DataType]
	if strings.HasSuffix(tf.ColumnType, "unsigned") {
		return "u" + gotype
	} else {
		return gotype
	}
}

func (d *mysqlDialect) Tables() []*Table {
	return d.fileds(d.tables()...)
}

// tables parse tables
func (d *mysqlDialect) tables() []string {
	var tables = make([]string, 0)
	err := d.db.Table(mysqlSchemaColumns).
		Where("TABLE_SCHEMA = ?", d.schema).
		Group("TABLE_NAME").
		Pluck("TABLE_NAME", &tables).
		Error
	if err != nil {
		logger.Fatal("load tables error: %v", err)
	}
	if len(tables) == 0 {
		logger.Warnf("load tables error: no table in %s", d.schema)
		os.Exit(1)
	}
	sort.Strings(tables)
	return tables
}

// fileds parse table columns
func (d *mysqlDialect) fileds(v ...string) []*Table {
	var tables = make([]*Table, 0, len(v))

	for _, item := range v {
		logger.Debugf("find table: %s", item)

		var table = &Table{TableName: item}

		{
			// 获取表字段信息
			err := d.db.Table(mysqlSchemaColumns).
				Select([]string{
					"TABLE_NAME AS table_name",
					"COLUMN_NAME AS column_name",
					"COLUMN_KEY AS column_key",
					"EXTRA AS extra",
					"DATA_TYPE AS data_type",
					"COLUMN_TYPE AS column_type",
					"IS_NULLABLE AS is_null",
					"COLUMN_COMMENT AS column_comment",
				}).
				Where("TABLE_SCHEMA = ? AND TABLE_NAME = ?", d.schema, table.TableName).
				Order("ORDINAL_POSITION").
				Find(&table.Fields).
				Error
			if err != nil {
				logger.Fatal("information_columns error: %v", err)
			}
		}

		{
			// 获取表注释
			err := d.db.Table(mysqlSchemaTables).
				Where("TABLE_SCHEMA = ? AND TABLE_NAME = ?", d.schema, table.TableName).
				Pluck("TABLE_COMMENT", &table.TableComment).
				Error
			if err != nil {
				logger.Fatal("information_tables error: %v", err)
			}

			if len(strings.TrimSpace(table.TableComment)) == 0 {
				table.TableComment = "."
			}
		}

		tables = append(tables, table)
	}

	return tables
}
