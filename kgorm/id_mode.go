package kgorm

import (
	"time"

	"github.com/xuender/kdb/id"
	"gorm.io/gorm"
)

type IDModel struct {
	ID        id.ID `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (p *IDModel) BeforeCreate(_ *gorm.DB) error {
	if p.ID == 0 {
		p.ID = id.New()
	}

	return nil
}
