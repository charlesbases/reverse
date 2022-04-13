package main

import (
	"strings"

	"github.com/charlesbases/reverse/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// postgresDialector .
type postgresDialector struct {
	opts   *Options
	gormDB *gorm.DB

	schema string
}

// PostgresDialector .
func PostgresDialector(dsn string) Dialector {
	var opts = &Options{Dsn: dsn}

	dialer := &postgresDialector{opts: opts, gormDB: gormOpen(postgres.Open(opts.Dsn))}
	dialer.schema = dialer.Schema()

	return dialer
}

// Schema 数据库
func (d *postgresDialector) Schema() string {
	if d.schema != "" {
		return d.schema
	}
	args := strings.Split(d.opts.Dsn, " ")
	for _, arg := range args {
		if strings.HasPrefix(arg, "dbname") {
			var sli = strings.Split(arg, "=")
			if len(sli) == 2 {
				return sli[1]
			}
		}
	}

	logger.Fatalf(`invalid schema. "%s"`, d.opts.Dsn)
	return ""
}

// Parse .
func (d *postgresDialector) Parse() *Plugin {
	panic("implement me")
}
