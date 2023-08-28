package kdb_test

import (
	"os"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/kdb"
)

func TestIdWorder_ID(t *testing.T) {
	t.Parallel()

	const max = 1_000

	ass := assert.New(t)
	worker := kdb.NewID()
	ids := make([]uint64, max)

	for i := 0; i < max; i++ {
		ids[i] = worker.ID()
	}

	for i := 1; i < max; i++ {
		ass.Greater(ids[i], ids[i-1])
	}
}

// nolint: paralleltest
func TestGetMachine(t *testing.T) {
	ass := assert.New(t)

	patch := gomonkey.ApplyFuncReturn(kdb.GetIP, nil)
	defer patch.Reset()

	ass.Equal(uint8(0), kdb.GetMachine())

	patch2 := gomonkey.ApplyFuncReturn(os.Getenv, "3")
	defer patch2.Reset()

	ass.Equal(uint8(3), kdb.GetMachine())
}
