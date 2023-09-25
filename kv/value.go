package kv

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"errors"
	"reflect"

	"google.golang.org/protobuf/proto"
)

type Marshaler interface {
	Marshal() ([]byte, error)
}

type Unmarshaler interface {
	Unmarshal([]byte) error
}

// nolint: gochecknoglobals
var (
	_messageType     = reflect.TypeOf((*proto.Message)(nil)).Elem()
	_jsonType        = reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	_unmarshalerType = reflect.TypeOf((*Unmarshaler)(nil)).Elem()
	ErrJSON          = errors.New("json Unmarshaler error")
	ErrUnmarshal     = errors.New("unmarshal error")
	ErrNoInterface   = errors.New("no interface")
)

func Marshal(data any) ([]byte, error) {
	switch value := data.(type) {
	case []byte:
		return value, nil
	case string:
		return []byte(value), nil
	case proto.Message:
		return proto.Marshal(value)
	case json.Marshaler:
		return value.MarshalJSON()
	case Marshaler:
		return value.Marshal()
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
	if err := unInterface(data, value); !errors.Is(err, ErrNoInterface) {
		return err
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

func unInterface[V any](data []byte, value V) error {
	valValue := reflect.ValueOf(value)

	switch {
	case valValue.Type().Implements(_messageType):
		if val, ok := valValue.Interface().(proto.Message); ok {
			return proto.Unmarshal(data, val)
		}

		return proto.Error
	case valValue.Type().Implements(_jsonType):
		if val, ok := valValue.Interface().(json.Unmarshaler); ok {
			return val.UnmarshalJSON(data)
		}

		return ErrJSON
	case valValue.Type().Implements(_unmarshalerType):
		if val, ok := valValue.Interface().(Unmarshaler); ok {
			return val.Unmarshal(data)
		}

		return ErrUnmarshal
	}

	return ErrNoInterface
}
