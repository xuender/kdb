package kv

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"reflect"

	"google.golang.org/protobuf/proto"
)

func Marshal(data any) ([]byte, error) {
	switch value := data.(type) {
	case []byte:
		return value, nil
	case string:
		return []byte(value), nil
	case int:
		return Key(int64(value)), nil
	case uint:
		return Key(uint64(value)), nil
	case proto.Message:
		return proto.Marshal(value)
	case json.Marshaler:
		return value.MarshalJSON()
	default:
		buf := &bytes.Buffer{}
		if binary.Write(buf, binary.BigEndian, data) == nil {
			return buf.Bytes(), nil
		}

		err := gob.NewEncoder(buf).Encode(value)

		return buf.Bytes(), err
	}
}

var (
	_messageType = reflect.TypeOf((*proto.Message)(nil)).Elem()
	_jsonType    = reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
)

func Unmarshal[V any](data []byte, value V) error {
	valValue := reflect.ValueOf(value)
	switch {
	case valValue.Type().Implements(_messageType):
		val := valValue.Interface().(proto.Message)

		return proto.Unmarshal(data, val)
	case valValue.Type().Implements(_jsonType):
		val := valValue.Interface().(json.Unmarshaler)

		return val.UnmarshalJSON(data)
	}

	var ret V

	switch val := any(ret).(type) {
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
	default:
		pass(val)

		buf := bytes.NewBuffer(data)
		if binary.Read(buf, binary.BigEndian, value) == nil {
			return nil
		}

		return gob.NewDecoder(buf).Decode(value)
	}
}
