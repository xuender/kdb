package kgorm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kdb/kgorm"
)

func TestIDModel_BeforeCreate(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	model := &kgorm.IDModel{}
	mem := kgorm.NewMemDB()

	_ = mem.AutoMigrate(model)
	ass.Nil(mem.Create(model).Error)
	ass.Greater(model.ID, uint64(0))
}

func TestIDModel(t *testing.T) {
	t.Parallel()

	type Model struct {
		kgorm.IDModel
		Name string
	}

	ass := assert.New(t)
	model := &Model{Name: "title"}
	mem := kgorm.NewDB("mem")

	_ = mem.AutoMigrate(model)
	ass.Nil(mem.Create(model).Error)
	ass.Greater(model.ID, uint64(0))
	ass.Equal(model.Name, "title")
}
