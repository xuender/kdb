package id

import (
	"strconv"

	"github.com/xuender/kit/types"
)

// nolint: gochecknoglobals
var _default Worker

// nolint: gochecknoinits
func init() {
	SetDefault(NewWorker(GetMachine()))
}

// ID 结构.
// 40 位作为毫秒数(35年).
// 8 位序号(256个).
// 4 位服务器代码(16个).
type ID uint64

// ID 返回不重复的ID.
func New() ID {
	return _default.Next()
}

func SetDefault(worker Worker) {
	_default = worker
}

func Parse(str string) (ID, error) {
	id, err := types.ParseInteger[uint](str)
	if err == nil {
		return ID(id), nil
	}

	return 0, err
}

func (p ID) String() string {
	return strconv.Itoa(int(p))
}

// Decode ID解码.
func (p ID) Decode() (uint64, uint64, uint64) {
	machine := p % _machineMax
	serial := (p >> _serialPosition) % _serialMax
	ts := p >> _tsPosition

	return uint64(ts + _min), uint64(serial), uint64(machine)
}
