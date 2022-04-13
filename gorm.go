package main

import (
	"context"
	"errors"
	"time"

	"github.com/charlesbases/reverse/logger"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

// gormOpen open gorm
func gormOpen(d gorm.Dialector) *gorm.DB {
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
