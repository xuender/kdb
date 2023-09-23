package kv_test

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/kdb/kv"
)

func TestMarshal(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ass.Equal([]byte{0, 0, 0, 0, 0, 0, 0, 0x22}, lo.Must1(kv.Marshal(34)))
	ass.Equal([]byte{1}, lo.Must1(kv.Marshal(true)))
	ass.Equal([]byte{0x40, 0x41, 0, 0, 0, 0, 0, 0}, lo.Must1(kv.Marshal(float64(34))))
	ass.Equal([]byte{0x61, 0x61}, lo.Must1(kv.Marshal("aa")))
	ass.Equal([]byte{0xb, 0x7f, 0x2, 0x1, 0x2, 0xff, 0x80, 0x0, 0x1, 0x4, 0x0, 0x0, 0x6, 0xff, 0x80, 0x0, 0x2, 0x2, 0x4},
		lo.Must1(kv.Marshal([]int{1, 2})))
	ass.Equal([]byte{1}, lo.Must1(kv.Marshal(true)))
}

func TestUnmarshal(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	var str string

	ass.Nil(kv.Unmarshal([]byte{0x61, 0x61}, &str))
	ass.Equal("aa", str)

	var num int64

	ass.Nil(kv.Unmarshal([]byte{0, 0, 0, 0, 0, 0, 0, 0x22}, &num))
	ass.Equal(int64(34), num)

	var num2 int

	ass.Nil(kv.Unmarshal([]byte{0, 0, 0, 0, 0, 0, 0, 0x22}, &num2))
	ass.Equal(int(34), num2)

	nums := []int{}

	ass.Nil(kv.Unmarshal([]byte{
		0xb, 0x7f, 0x2, 0x1, 0x2, 0xff, 0x80, 0x0,
		0x1, 0x4, 0x0, 0x0, 0x6, 0xff, 0x80, 0x0, 0x2, 0x2, 0x4,
	},
		&nums))
	ass.Equal([]int{1, 2}, nums)

	var boo bool

	ass.Nil(kv.Unmarshal([]byte{1}, &boo))
	ass.True(boo)
}
