package config

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

// DefaultDBConfig default db config
const DefaultDBConfig = `source:
  # mysql: username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local&timeout=5s&readTimeout=5s&writeTimeout=5s
  # postgresql: postgres://username:password@127.0.0.1:5432/postgres?sslmode=disable&connect_timeout=5
  dsn: ""
  max_lifetime: 5s
  max_idle_conns: 10
  max_open_conns: 20
`

// DBConfig database config
type DBConfig struct {
	DSN          string        `json:"dsn" mapstructure:"dsn"`
	MaxLifetime  time.Duration `json:"max_lifetime" mapstructure:"max_lifetime"`
	MaxIdleConns int           `json:"max_idle_conns" mapstructure:"max_idle_conns"`
	MaxOpenConns int           `json:"max_open_conns" mapstructure:"max_open_conns"`
}

// NewDB returns xorm.Engine
func NewDB(driver string) (*xorm.Engine, error) {
	c := &DBConfig{}
	if err := config.ReadSection("source", c); err != nil {
		return nil, err
	}

	db, err := xorm.NewEngine(driver, c.DSN)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(c.MaxLifetime)

	// conn pool
	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetMaxOpenConns(c.MaxOpenConns)

	return db, nil
}
