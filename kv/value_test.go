package kv_test

import (
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/kdb/kv"
)

type St struct {
	Str string
}

type Mt struct {
	Str string
}

func (p *Mt) Marshal() ([]byte, error) {
	return []byte(p.Str), nil
}

func (p *Mt) Unmarshal(data []byte) error {
	p.Str = string(data)

	return nil
}

func TestMarshal(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ass.Equal([]byte{
		23, 127, 3, 1, 1, 2, 83, 116, 1, 255, 128,
		0, 1, 1, 1, 3, 83, 116, 114, 1, 12, 0, 0, 0,
		8, 255, 128, 1, 3, 120, 120, 120, 0,
	}, lo.Must1(kv.Marshal(St{Str: "xxx"})))
	ass.Equal([]byte{0, 0, 0, 0, 0, 0, 0, 0x22}, lo.Must1(kv.Marshal(34)))
	ass.Equal([]byte{1}, lo.Must1(kv.Marshal(true)))
	ass.Equal([]byte{0x40, 0x41, 0, 0, 0, 0, 0, 0}, lo.Must1(kv.Marshal(float64(34))))
	ass.Equal([]byte{0x61, 0x61}, lo.Must1(kv.Marshal("aa")))
	ass.Equal([]byte{0x61, 0x61}, lo.Must1(kv.Marshal(&Mt{Str: "aa"})))
	ass.NotNil(lo.Must1(kv.Marshal([]int{1})))
	ass.NotNil(lo.Must1(kv.Marshal([]uint{1})))
	ass.Equal([]byte{1}, lo.Must1(kv.Marshal(true)))
	ass.Equal([]byte{1}, lo.Must1(kv.Marshal([]byte{1})))
	ass.Equal([]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3}, lo.Must1(kv.Marshal(uint(3))))
	ass.NotNil(lo.Must1(kv.Marshal(time.Now())))
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

	var num3 uint

	ass.Nil(kv.Unmarshal([]byte{0, 0, 0, 0, 0, 0, 0, 0x22}, &num3))
	ass.Equal(uint(34), num3)

	nums := []int{}
	unums := []uint{}

	ass.Nil(kv.Unmarshal([]byte{
		0xb, 0x7f, 0x2, 0x1, 0x2, 0xff, 0x80, 0x0,
		0x1, 0x4, 0x0, 0x0, 0x6, 0xff, 0x80, 0x0, 0x2, 0x2, 0x4,
	},
		&nums))
	ass.Nil(kv.Unmarshal([]byte{11, 127, 2, 1, 2, 255, 128, 0, 1, 6, 0, 0, 5, 255, 128, 0, 1, 1}, &unums))
	ass.Equal([]int{1, 2}, nums)
	ass.Equal([]uint{1}, unums)

	var boo bool

	ass.Nil(kv.Unmarshal([]byte{1}, &boo))
	ass.True(boo)

	ste := &St{}

	ass.Nil(kv.Unmarshal([]byte{
		23, 127, 3, 1, 1, 2, 83, 116, 1, 255, 128, 0, 1, 1, 1,
		3, 83, 116, 114, 1, 12, 0, 0, 0, 8, 255, 128, 1, 3, 120, 120, 120, 0,
	}, ste))
	ass.Equal("xxx", ste.Str)

	var now time.Time

	ass.Nil(kv.Unmarshal([]byte{
		34, 50, 48, 50, 51, 45, 48, 57, 45, 50, 53, 84, 48, 56, 58, 48, 50, 58, 48,
		55, 46, 50, 57, 53, 48, 53, 52, 54, 49, 53, 43, 48, 56, 58, 48, 48, 34,
	}, &now))
	ass.Equal(2023, now.Year())

	bys := []byte{}
	ass.Nil(kv.Unmarshal([]byte{1}, &bys))
	ass.Equal(byte(1), bys[0])
}

func TestUnmarshal_mt(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	mt := &Mt{}

	ass.Nil(kv.Unmarshal([]byte{0x61, 0x61}, mt))
	ass.Equal("aa", mt.Str)
}
