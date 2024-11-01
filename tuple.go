package main

import "errors"

// Error
var (
	ErrExtraField     = errors.New("field number more than avaliable in tuple")
	ErrWrongFieldType = errors.New("wrong field type")
)

// Tuple is the interface all tuples need to implement.
type Tuple interface {
	SetField(int, any) error
	GetField(int) any
}

// Tuple1
type Tuple1[V1 any] struct {
	V1 V1
}

func (t *Tuple1[V1]) SetField(field int, val any) error {
	if field != 0 {
		return ErrExtraField
	}
	nval, ok := val.(V1)
	if !ok {
		return ErrWrongFieldType
	}
	t.V1 = nval
	return nil
}

func (t *Tuple1[V1]) GetField(field int) any {
	if field != 0 {
		panic("wrong field")
	}
	return t.V1
}

// Tuple2
type Tuple2[V1, V2 any] struct {
	V1 V1
	V2 V2
}

func (t *Tuple2[V1, V2]) SetField(field int, val any) error {
	switch field {
	case 0:
		nval, ok := val.(V1)
		if !ok {
			return ErrWrongFieldType
		}
		t.V1 = nval
		return nil
	case 1:
		nval, ok := val.(V2)
		if !ok {
			return ErrWrongFieldType
		}
		t.V2 = nval
		return nil
	default:
		return ErrExtraField
	}
}

func (t *Tuple2[V1, V2]) GetField(field int) any {
	switch field {
	case 0:
		return t.V1
	case 1:
		return t.V2
	default:
		panic("wrong field")
	}
}

// Tuple3
type Tuple3[V1, V2, V3 any] struct {
	V1 V1
	V2 V2
	V3 V3
}

func (t *Tuple3[V1, V2, V3]) SetField(field int, val any) error {
	switch field {
	case 0:
		nval, ok := val.(V1)
		if !ok {
			return ErrWrongFieldType
		}
		t.V1 = nval
		return nil
	case 1:
		nval, ok := val.(V2)
		if !ok {
			return ErrWrongFieldType
		}
		t.V2 = nval
		return nil
	case 2:
		nval, ok := val.(V3)
		if !ok {
			return ErrWrongFieldType
		}
		t.V3 = nval
		return nil
	default:
		return ErrExtraField
	}
}

func (t *Tuple3[V1, V2, V3]) GetField(field int) any {
	switch field {
	case 0:
		return t.V1
	case 1:
		return t.V2
	case 2:
		return t.V3
	default:
		panic("wrong field")
	}
}
