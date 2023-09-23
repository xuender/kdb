package kv

import (
	"github.com/cockroachdb/pebble"
	"github.com/samber/lo"
)

type KV[K KeyType, V any] struct {
	*pebble.DB
}

func New[K KeyType, V any](pdb *pebble.DB) *KV[K, V] {
	return &KV[K, V]{pdb}
}

func NewPath[K KeyType, V any](path string) *KV[K, V] {
	return &KV[K, V]{lo.Must1(pebble.Open(path, &pebble.Options{}))}
}

func (p *KV[K, V]) GetItem(key K) (V, bool) {
	data, closer, err := p.Get(Key(key))
	if err == nil && closer != nil {
		closer.Close()
	}

	var ret V

	if err == nil {
		lo.Must0(Unmarshal(data, &ret))

		return ret, true
	}

	return ret, false
}

func (p *KV[K, V]) SetItem(key K, value V) error {
	data, err := Marshal(value)
	if err != nil {
		return err
	}

	return p.Set(Key(key), data, pebble.Sync)
}
