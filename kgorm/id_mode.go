package kgorm

import (
	"github.com/xuender/kdb/id"
	"gorm.io/gorm"
)

type IDModel struct {
	ID        id.ID          `gorm:"primarykey" json:"id"`
	DeletedAt gorm.DeletedAt `gorm:"index"      json:"-"`
	CreatedAt Msec           `json:"ca"`
	UpdatedAt Msec           `json:"ua"`
}

func (p *IDModel) BeforeCreate(_ *gorm.DB) error {
	if p.ID == 0 {
		p.ID = id.New()
	}

	return nil
}
