package kv_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kdb/kv"
)

func TestKey(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	ass.Equal([]byte{1}, kv.Key(true))
	ass.Equal([]byte{0x40, 0x41, 0, 0, 0, 0, 0, 0}, kv.Key(float64(34)))
	ass.Equal([]byte{0, 0, 0, 0, 0, 0, 0, 0x22}, kv.Key(int(34)))
	ass.Equal([]byte{0, 0, 0, 0, 0, 0, 0, 0x22}, kv.Key(uint(34)))
	ass.Equal([]byte{0, 0, 0, 0, 0, 0, 0, 0x22}, kv.Key(int64(34)))
	ass.Equal([]byte{0x61, 0x61}, kv.Key("aa"))
}

func TestToKey(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ass.Equal(int64(34), kv.ToKey[int64]([]byte{0, 0, 0, 0, 0, 0, 0, 0x22}))
	ass.Equal(int(34), kv.ToKey[int]([]byte{0, 0, 0, 0, 0, 0, 0, 0x22}))
	ass.Equal(uint(34), kv.ToKey[uint]([]byte{0, 0, 0, 0, 0, 0, 0, 0x22}))
	ass.Equal("aa", kv.ToKey[string]([]byte{0x61, 0x61}))
}
