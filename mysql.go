package main

import (
	"path"
	"sort"
	"strings"

	"github.com/charlesbases/generator"
	"github.com/charlesbases/reverse/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	mysqlInformationTables  = "information_schema.TABLES"
	mysqlInformationColumns = "information_schema.COLUMNS"
)

// mysqlDialector .
type mysqlDialector struct {
	opts   *Options
	gormDB *gorm.DB

	schema string
}

// MysqlDialector .
func MysqlDialector(dsn string) Dialector {
	var opts = &Options{Dsn: dsn}

	dialer := &mysqlDialector{opts: opts, gormDB: gormOpen(mysql.Open(opts.Dsn))}
	dialer.schema = dialer.Schema()

	return dialer
}

// Schema 数据库
func (d *mysqlDialector) Schema() string {
	if d.schema != "" {
		return d.schema
	}
	return path.Base(d.opts.Dsn)
}

// Parse .
func (d *mysqlDialector) Parse() *Plugin {
	tables := d.parseTables(d.tables()...)

	var plugin = &Plugin{
		schema:  d.schema,
		structs: make([]*Struct, 0, len(tables)),
		imports: make(map[string]*generator.ExternalPackage),
	}

	for _, table := range tables {
		plugin.structs = append(plugin.structs, &Struct{
			Name:  camelcase(table.TableName),
			Desc:  table.TableComment,
			Table: table,
		})
	}

	return plugin
}

// tables .
func (d *mysqlDialector) tables() []string {
	var tables = make([]string, 0)
	err := d.gormDB.Table(mysqlInformationColumns).
		Where("TABLE_SCHEMA = ?", d.schema).
		Group("TABLE_NAME").
		Pluck("TABLE_NAME", &tables).
		Error
	if err != nil {
		logger.Fatal("load tables error: %v", err)
	}
	if len(tables) == 0 {
		logger.Fatal("load tables error: no table in %s", d.schema)
	}
	sort.Strings(tables)

	return tables
}

// parseTables .
func (d *mysqlDialector) parseTables(v ...string) []*Table {
	var tables = make([]*Table, 0, len(v))

	for _, item := range v {
		logger.Debugf("find table: %s", item)

		var table = &Table{TableName: item}

		{
			// 获取表字段信息
			err := d.gormDB.Table(mysqlInformationColumns).
				Select([]string{
					"TABLE_NAME AS table_name",
					"COLUMN_NAME AS field_name",
					"COLUMN_KEY AS field_key",
					"EXTRA AS extra",
					"IS_NULLABLE AS is_null",
					"DATA_TYPE AS data_type",
					"COLUMN_TYPE AS field_type",
					"COLUMN_COMMENT AS field_comment",
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
			err := d.gormDB.Table(mysqlInformationTables).
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
