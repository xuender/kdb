package kv

import (
	"bytes"
	"encoding/binary"
	"reflect"

	"golang.org/x/exp/constraints"
)

type KeyType interface {
	constraints.Integer | constraints.Float | bool | string | []byte
}

func Key[K KeyType](key K, prefix ...byte) []byte {
	return append(prefix, keys(key)...)
}

func keys[K KeyType](key K) []byte {
	switch value := any(key).(type) {
	case []byte:
		return value
	case string:
		return []byte(value)
	case int:
		return Key(int64(value))
	case uint:
		return Key(uint64(value))
	}

	buf := &bytes.Buffer{}
	_ = binary.Write(buf, binary.BigEndian, key)

	return buf.Bytes()
}

func ToKey[K KeyType](data []byte, prefix ...byte) K {
	if bytes.HasPrefix(data, prefix) {
		return toKey[K](data[len(prefix):])
	}

	return toKey[K](data)
}

func toKey[K KeyType](data []byte) K {
	ret := new(K)

	switch value := any(*ret).(type) {
	case string:
		reflect.ValueOf(ret).Elem().SetString(string(data))

		return *ret
	case int:
		reflect.ValueOf(ret).Elem().SetInt(ToKey[int64](data))

		return *ret
	case uint:
		reflect.ValueOf(ret).Elem().SetUint(ToKey[uint64](data))

		return *ret
	default:
		pass(value)
	}

	buf := bytes.NewBuffer(data)
	_ = binary.Read(buf, binary.BigEndian, ret)

	return *ret
}

func pass(_ any) {}
