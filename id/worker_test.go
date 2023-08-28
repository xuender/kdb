package id_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kdb/id"
)

func TestIdWorker_Next(t *testing.T) {
	t.Parallel()

	const max = 10_000

	ass := assert.New(t)
	worker := id.NewWorker(0)
	ids := make([]id.ID, max)

	for i := 0; i < max; i++ {
		ids[i] = worker.Next()
	}

	for i := 1; i < max; i++ {
		ass.Greater(ids[i], ids[i-1])
	}
}
