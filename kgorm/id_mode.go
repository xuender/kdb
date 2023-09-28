package kgorm

import (
	"github.com/xuender/kdb/id"
	"gorm.io/gorm"
)

type IDModel struct {
	ID        uint64         `gorm:"primarykey" json:"id"`
	DeletedAt gorm.DeletedAt `gorm:"index"      json:"-"`
	CreatedAt Msec           `json:"ca"`
	UpdatedAt Msec           `json:"ua"`
}

// Version 乐观锁使用的修改时间.
func (p *IDModel) Version() Msec {
	return p.UpdatedAt
}

func (p *IDModel) BeforeCreate(_ *gorm.DB) error {
	if p.ID == 0 {
		p.ID = id.New()
	}

	return nil
}
