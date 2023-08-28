package kdb

import (
	"os"
	"sync"
	"time"

	"github.com/xuender/kit/types"
)

const (
	// _min 2023-08-28.
	_min            = 1693191157559
	_tsPosition     = 12
	_serialPosition = 4
	_serialMax      = 256
	_machineMax     = 16
	MachineKey      = "MACHINE_NUM"
)

type IDer interface {
	ID() uint64
}

// idWorker 序号生成.
type idWorker struct {
	machine uint64
	mutex   sync.Mutex
	serial  uint8
	stamp   int64
}

// NewID ID生成器.
// ID 结构
// 时间戳.
// 12 位序号.
// 4 位机器码.
func NewID() IDer {
	return &idWorker{
		machine: uint64(GetMachine() % _machineMax),
		stamp:   time.Now().UnixMilli() - _min,
	}
}

func GetMachine() uint8 {
	if machinKey := os.Getenv(MachineKey); machinKey != "" {
		if key, err := types.ParseInteger[uint8](machinKey); err == nil {
			return key
		}
	}

	bytes := GetIP()
	if len(bytes) == 0 {
		return 0
	}

	return bytes[len(bytes)-1]
}

func (p *idWorker) ID() uint64 {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	now := time.Now().UnixMilli() - _min
	for ; p.serial == 0 && now == p.stamp; now = time.Now().UnixMilli() - _min {
		time.Sleep(time.Microsecond)
	}

	ret := uint64(now<<_tsPosition) + (uint64(p.serial) << _serialPosition)

	p.serial++
	p.stamp = now

	return ret | p.machine
}

// Decode ID解码.
func Decode(id uint64) (uint64, uint64, uint64) {
	machine := id % _machineMax
	serial := (id >> _serialPosition) % _serialMax
	ts := id >> _tsPosition

	return ts + _min, serial, machine
}
