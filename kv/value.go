package kv

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"reflect"

	"github.com/gogo/protobuf/proto"
)

func Marshal(data any) ([]byte, error) {
	switch value := data.(type) {
	case []byte:
		return value, nil
	case string:
		return []byte(value), nil
	case json.Marshaler:
		return value.MarshalJSON()
	case proto.Message:
		return proto.Marshal(value)
	case int:
		return Key(int64(value)), nil
	case uint:
		return Key(uint64(value)), nil
	default:
		buf := &bytes.Buffer{}
		if binary.Write(buf, binary.BigEndian, data) == nil {
			return buf.Bytes(), nil
		}

		err := gob.NewEncoder(buf).Encode(value)

		return buf.Bytes(), err
	}
}

func Unmarshal[V any](data []byte, value V) error {
	ret := new(V)

	switch val := any(*ret).(type) {
	case *string:
		reflect.ValueOf(value).Elem().SetString(string(data))

		return nil
	case *[]byte:
		reflect.ValueOf(value).Elem().Set(reflect.ValueOf(data))

		return nil
	case *int:
		reflect.ValueOf(value).Elem().SetInt(ToKey[int64](data))

		return nil
	case *uint:
		reflect.ValueOf(value).Elem().SetUint(ToKey[uint64](data))

		return nil
	case json.Unmarshaler:
		err := val.UnmarshalJSON(data)
		reflect.ValueOf(value).Elem().Set(reflect.ValueOf(val))

		return err
	case proto.Message:
		err := proto.Unmarshal(data, val)
		reflect.ValueOf(value).Elem().Set(reflect.ValueOf(val))

		return err
	default:
		pass(val)

		buf := bytes.NewBuffer(data)
		if binary.Read(buf, binary.BigEndian, value) == nil {
			return nil
		}

		return gob.NewDecoder(buf).Decode(value)
	}
}
