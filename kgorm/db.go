package kgorm

import (
	"strings"

	"gorm.io/gorm"
)

func NewDB(dsn string) *gorm.DB {
	switch {
	case strings.HasPrefix(dsn, "file"), strings.HasPrefix(dsn, "mem"):
		return NewMemDB()
	default:
		return NewMysqlDB(dsn)
	}
}
