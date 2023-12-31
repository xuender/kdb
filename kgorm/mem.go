package kgorm

import (
	"github.com/samber/lo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewMemDB() *gorm.DB {
	memDB := lo.Must1(gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   NewGormLogger(),
	}))

	memDB.Debug()

	return memDB
}
