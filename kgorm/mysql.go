package kgorm

import (
	"time"

	"github.com/samber/lo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDB(dsn string) *gorm.DB {
	var (
		mysqlDB = lo.Must1(gorm.Open(mysql.Open(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			Logger:                                   NewGormLogger(),
		}))
		sqlDB   = lo.Must1(mysqlDB.DB())
		maxOpen = 200
		maxIdle = 100
	)

	sqlDB.SetMaxOpenConns(maxOpen)
	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetConnMaxLifetime(time.Minute)

	return mysqlDB
}
