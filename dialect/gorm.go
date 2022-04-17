package dialect

import (
	"context"
	"errors"
	"time"

	"github.com/charlesbases/reverse/logger"
	"github.com/charlesbases/reverse/types"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

type (
	Table struct {
		TableName    string         // 表名
		TableComment string         // 表注释
		Fields       []*TableColumn // 表字段
	}

	TableColumn struct {
		ColumnName    string // 列名
		ColumnKey     string // 键类别
		Extra         string // 自增
		IsNull        string // NOT NULL
		DataType      string // 类型
		ColumnType    string // 类型+长度
		ColumnComment string // 注释
	}
)

type Dialect interface {
	Options() *types.Options
	Tables() []*Table
	Schema() string
	ParseColumnTag(tf *TableColumn) types.Tag
	ParseColumnType(tf *TableColumn) string
}

// open gorm.Open
func open(d gorm.Dialector) *gorm.DB {
	db, err := gorm.Open(d, &gorm.Config{Logger: new(gorml)})
	if err != nil {
		logger.Fatalf(" - db connect failed. %v", err)
	}

	logger.Infor("connection succeeded")
	return db
}

// gorml .
type gorml struct{}

func (gl *gorml) LogMode(level glogger.LogLevel) glogger.Interface {
	return gl
}

func (gl *gorml) Info(ctx context.Context, s string, i ...interface{}) {
	logger.Inforf(s, i...)
}

func (gl *gorml) Warn(ctx context.Context, s string, i ...interface{}) {
	logger.Warnf(s, i...)
}

func (gl *gorml) Error(ctx context.Context, s string, i ...interface{}) {
	logger.Errorf(s, i...)
}

func (gl *gorml) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()

	switch {
	case err != nil && !errors.Is(err, gorm.ErrRecordNotFound):
		logger.Errorf("%s | %v", sql, err)
	}
}
