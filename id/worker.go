package id

import (
	"sync"
	"time"
)

type Worker interface {
	Next() uint64
}

// worker 序号生成.
type worker struct {
	machine uint64
	mutex   sync.Mutex
	serial  uint8
	stamp   int64
}

// NewWorker ID生成器.
func NewWorker(machine uint8) Worker {
	return &worker{
		machine: uint64(machine % _machineMax),
		stamp:   time.Now().UnixMilli() - _min,
	}
}

func (p *worker) Next() uint64 {
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
