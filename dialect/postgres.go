package dialect

import (
	"strings"

	"github.com/charlesbases/reverse/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// postgresDialect .
type postgresDialect struct {
	opts   *types.Options
	db     *gorm.DB
	schema string
}

// Postgres .
func Postgres(opts *types.Options) Dialect {
	d := &postgresDialect{opts: opts}
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
	panic("implement me")
}

func (d *postgresDialect) ParseColumnType(tf *TableColumn) string {
	var gotype = types.PostgresType2GoType[tf.DataType]
	return gotype
}

func (d *postgresDialect) Tables() []*Table {
	panic("implement me")
}
