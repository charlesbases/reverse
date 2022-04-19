package dialer

import (
	"os"
	"sort"
	"strings"

	"github.com/charlesbases/reverse/logger"
	"github.com/charlesbases/reverse/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	postgresSchema        = "public"
	postgresSchemaColumns = "information_schema.columns"
)

// postgresDialect .
type postgresDialect struct {
	opts *types.Options
	db   *gorm.DB

	schema string
}

// Postgres .
func Postgres(opts *types.Options) Dialect {
	d := &postgresDialect{opts: opts}
	d.connect()
	return d
}

// connect .
func (d *postgresDialect) connect() {
	d.db = open(postgres.Open(d.opts.Dsn))

	// schema
	args := strings.Split(d.opts.Dsn, " ")
	for _, arg := range args {
		if strings.HasPrefix(arg, "dbname") {
			var sli = strings.Split(arg, "=")
			if len(sli) == 2 {
				d.schema = sli[1]
			}
		}
	}
}

func (d *postgresDialect) Options() *types.Options {
	return d.opts
}

func (d *postgresDialect) Schema() string {
	return d.schema
}

func (d *postgresDialect) ParseColumnTag(tf *TableColumn) types.Tag {
	var tag types.Tag = make([]string, 0)

	// json
	tag.Append("json", tf.ColumnName)

	// orm
	var ormtag types.TagType = make([]string, 0)
	ormtag.Append("column", tf.ColumnName)
	if tf.IsNull == "NO" {
		ormtag.Append("not null")
	}

	tag.Append("gorm", ormtag)
	return tag
}

func (d *postgresDialect) ParseColumnType(tf *TableColumn) string {
	var gotype = types.PostgresType2GoType[tf.DataType]
	return gotype
}

func (d *postgresDialect) Tables() []*Table {
	return d.load(d.tables()...)
}

// tables .
func (d *postgresDialect) tables() []string {
	var tables = make([]string, 0)
	err := d.db.Table("pg_tables").
		Where("schemaname = ?", postgresSchema).
		Pluck("tablename", &tables).
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

// load .
func (d *postgresDialect) load(v ...string) []*Table {
	// tables comment
	var tablesComment = make(map[string]string, len(v))
	{
		var expands = make([]*TableExpand, 0, len(v))
		err := d.db.Table("information_schema.tables ist").
			Select([]string{"ist.table_name AS table_name", "d.description AS table_desc"}).
			Joins("JOIN pg_class c ON c.relname = ist.table_name").
			Joins("LEFT JOIN pg_description d ON d.objoid = c.oid AND d.objsubid = '0'").
			Where("ist.table_schema = ? AND ist.table_name IN ?", postgresSchema, v).
			Find(&expands).
			Error
		if err != nil {
			logger.Fatal("load tables comment failed. %v", err)
		}

		for _, item := range expands {
			tablesComment[item.TableName] = item.TableDesc
		}
	}

	// tables columns
	var tables = make([]*Table, 0, len(v))
	{
		for _, item := range v {
			logger.Debugf("find table: %s", item)

			var table = new(Table)
			table.TableName = item

			// table columns
			err := d.db.Table("information_schema.columns isc").
				Select([]string{
					"isc.table_name",
					"isc.column_name",
					// "? AS column_key",
					// "? AS extra",
					"isc.udt_name AS data_type",
					// "? AS column_type",
					"isc.is_nullable AS is_null",
					"pgd.description AS column_desc",
				}).
				Joins("JOIN pg_class pgc ON pgc.relname = isc.table_name").
				Joins("LEFT JOIN pg_description pgd ON pgd.objoid = pgc.oid AND pgd.objsubid = isc.ordinal_position").
				Where("isc.table_schema = ? AND isc.table_name = ?", postgresSchema, item).
				Order("isc.ordinal_position").
				Find(&table.Columns).
				Error
			if err != nil {
				logger.Fatalf("load %s.columns failed. %v", item, err)
			}

			// table columns comment
			for _, column := range table.Columns {
				column.ColumnDesc = strings.ReplaceAll(column.ColumnDesc, "\n", "  ")
				column.ColumnDesc = strings.TrimSpace(column.ColumnDesc)

				if len(column.ColumnDesc) == 0 {
					column.ColumnDesc = column.ColumnName
				}
			}

			// table comment
			if tableComment, found := tablesComment[item]; found {
				table.TableDesc = strings.TrimSpace(tableComment)
			}
			if len(table.TableDesc) == 0 {
				table.TableDesc = "."
			}

			tables = append(tables, table)
		}
	}

	return tables
}
