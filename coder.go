package main

import (
	"reflect"

	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/varint"
)

type TupleEncoderFunc func(Tuple, []byte) int
type TupleDecoderFunc func(Tuple, []byte) (int, error)

// Encoder function for tuple
func CreateEncoderFunc(t Tuple) TupleEncoderFunc {
	typ := reflect.Indirect(reflect.ValueOf(t)).Type()
	numField := typ.NumField()

	encodeFuncs := []EncodeFunc{}
	for i := 0; i < numField; i++ {
		kind := typ.Field(i).Type.Kind()
		if !SupportedKind(kind) {
			panic("Unsupported field type=" + kind.String())
		}
		encodeFuncs = append(encodeFuncs, EncodeFromKind(kind))
	}

	return func(t Tuple, buf []byte) int {
		offset := 0
		for i, f := range encodeFuncs {
			offset += f(t.GetField(i), buf[offset:])
		}
		return offset
	}
}

// Decoder function for tuple
func CreateDecoderFunc(t Tuple) TupleDecoderFunc {
	typ := reflect.Indirect(reflect.ValueOf(t)).Type()
	numField := typ.NumField()

	decodeFuncs := []DecodeFunc{}
	for i := 0; i < numField; i++ {
		kind := typ.Field(i).Type.Kind()
		if !SupportedKind(kind) {
			panic("Unsupported field type=" + kind.String())
		}
		decodeFuncs = append(decodeFuncs, DecodeFromKind(kind))
	}

	return func(t Tuple, buf []byte) (int, error) {
		offset := 0
		for i, f := range decodeFuncs {
			v, n, err := f(buf[offset:])
			offset += n
			if err != nil {
				return offset, err
			}
			t.SetField(i, v)
		}
		return offset, nil
	}
}

// Encoder functions
type EncodeFunc func(any, []byte) int

func EncodeFromKind(kind reflect.Kind) EncodeFunc {
	switch kind {
	case reflect.Bool:
		return EncodeBool
	case reflect.Int:
		return EncodeInt
	case reflect.Int8:
		return EncodeInt8
	case reflect.Int16:
		return EncodeInt16
	case reflect.Int32:
		return EncodeInt32
	case reflect.Int64:
		return EncodeInt64
	case reflect.Uint:
		return EncodeUint
	case reflect.Uint8:
		return EncodeUint8
	case reflect.Uint16:
		return EncodeUint16
	case reflect.Uint32:
		return EncodeUint32
	case reflect.Uint64:
		return EncodeUint64
	case reflect.Float32:
		return EncodeFloat32
	case reflect.Float64:
		return EncodeFloat64
	case reflect.String:
		return EncodeString
	default:
		panic("unknown kind=" + kind.String())
	}
}

func EncodeBool(b any, buf []byte) int {
	var n uint8
	if b.(bool) {
		n = 1
	} else {
		n = 0
	}
	return varint.MarshalUint8(n, buf)
}

func EncodeInt(i any, buf []byte) int {
	return varint.MarshalInt(i.(int), buf)
}

func EncodeInt8(i any, buf []byte) int {
	return varint.MarshalInt8(i.(int8), buf)
}

func EncodeInt16(i any, buf []byte) int {
	return varint.MarshalInt16(i.(int16), buf)
}

func EncodeInt32(i any, buf []byte) int {
	return varint.MarshalInt32(i.(int32), buf)
}

func EncodeInt64(i any, buf []byte) int {
	return varint.MarshalInt64(i.(int64), buf)
}

func EncodeUint(i any, buf []byte) int {
	return varint.MarshalUint(i.(uint), buf)
}

func EncodeUint8(i any, buf []byte) int {
	return varint.MarshalUint8(i.(uint8), buf)
}

func EncodeUint16(i any, buf []byte) int {
	return varint.MarshalUint16(i.(uint16), buf)
}

func EncodeUint32(i any, buf []byte) int {
	return varint.MarshalUint32(i.(uint32), buf)
}

func EncodeUint64(i any, buf []byte) int {
	return varint.MarshalUint64(i.(uint64), buf)
}

func EncodeFloat32(f any, buf []byte) int {
	return varint.MarshalFloat32(f.(float32), buf)
}

func EncodeFloat64(f any, buf []byte) int {
	return varint.MarshalFloat64(f.(float64), buf)
}

func EncodeString(s any, buf []byte) int {
	return ord.MarshalString(s.(string), nil, buf)
}

// Decoder functions

type DecodeFunc func([]byte) (any, int, error)

func DecodeFromKind(kind reflect.Kind) DecodeFunc {
	switch kind {
	case reflect.Bool:
		return DecodeBool
	case reflect.Int:
		return DecodeInt
	case reflect.Int8:
		return DecodeInt8
	case reflect.Int16:
		return DecodeInt16
	case reflect.Int32:
		return DecodeInt32
	case reflect.Int64:
		return DecodeInt64
	case reflect.Uint:
		return DecodeUint
	case reflect.Uint8:
		return DecodeUint8
	case reflect.Uint16:
		return DecodeUint16
	case reflect.Uint32:
		return DecodeUint32
	case reflect.Uint64:
		return DecodeUint64
	case reflect.Float32:
		return DecodeFloat32
	case reflect.Float64:
		return DecodeFloat64
	case reflect.String:
		return DecodeString
	default:
		panic("unknown kind=" + kind.String())
	}
}

func DecodeBool(buf []byte) (any, int, error) {
	v, n, err := varint.UnmarshalUint8(buf)
	if v == uint8(1) {
		return true, n, err
	} else {
		return false, n, err
	}
}

func DecodeInt(buf []byte) (any, int, error) {
	return varint.UnmarshalInt(buf)
}

func DecodeInt8(buf []byte) (any, int, error) {
	return varint.UnmarshalInt8(buf)
}

func DecodeInt16(buf []byte) (any, int, error) {
	return varint.UnmarshalInt16(buf)
}

func DecodeInt32(buf []byte) (any, int, error) {
	return varint.UnmarshalInt32(buf)
}

func DecodeInt64(buf []byte) (any, int, error) {
	return varint.UnmarshalInt64(buf)
}

func DecodeUint(buf []byte) (any, int, error) {
	return varint.UnmarshalUint(buf)
}

func DecodeUint8(buf []byte) (any, int, error) {
	return varint.UnmarshalUint8(buf)
}

func DecodeUint16(buf []byte) (any, int, error) {
	return varint.UnmarshalUint16(buf)
}

func DecodeUint32(buf []byte) (any, int, error) {
	return varint.UnmarshalUint32(buf)
}

func DecodeUint64(buf []byte) (any, int, error) {
	return varint.UnmarshalUint64(buf)
}

func DecodeFloat32(buf []byte) (any, int, error) {
	return varint.UnmarshalFloat32(buf)
}

func DecodeFloat64(buf []byte) (any, int, error) {
	return varint.UnmarshalFloat64(buf)
}

func DecodeString(buf []byte) (any, int, error) {
	return ord.UnmarshalString(nil, buf)
}
