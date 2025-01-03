package godata

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

	New() Tuple
}

// Tuple1
type Tuple1[V1 any] struct {
	V1 V1
}

func T1[V1 any](v1 V1) Tuple1[V1] {
	return Tuple1[V1]{V1: v1}
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

func (t *Tuple1[V1]) New() Tuple {
	return &Tuple1[V1]{}
}

// Tuple2
type Tuple2[V1, V2 any] struct {
	V1 V1
	V2 V2
}

func T2[V1, V2 any](v1 V1, v2 V2) Tuple2[V1, V2] {
	return Tuple2[V1, V2]{V1: v1, V2: v2}
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

func (t *Tuple2[V1, V2]) New() Tuple {
	return &Tuple2[V1, V2]{}
}

// Tuple3
type Tuple3[V1, V2, V3 any] struct {
	V1 V1
	V2 V2
	V3 V3
}

func T3[V1, V2, V3 any](v1 V1, v2 V2, v3 V3) Tuple3[V1, V2, V3] {
	return Tuple3[V1, V2, V3]{V1: v1, V2: v2, V3: v3}
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

func (t *Tuple3[V1, V2, V3]) New() Tuple {
	return &Tuple3[V1, V2, V3]{}
}

// Tuple4
type Tuple4[V1, V2, V3, V4 any] struct {
	V1 V1
	V2 V2
	V3 V3
	V4 V4
}

func T4[V1, V2, V3, V4 any](v1 V1, v2 V2, v3 V3, v4 V4) Tuple4[V1, V2, V3, V4] {
	return Tuple4[V1, V2, V3, V4]{V1: v1, V2: v2, V3: v3, V4: v4}
}

func (t *Tuple4[V1, V2, V3, V4]) SetField(field int, val any) error {
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
	case 3:
		nval, ok := val.(V4)
		if !ok {
			return ErrWrongFieldType
		}
		t.V4 = nval
		return nil
	default:
		return ErrExtraField
	}
}

func (t *Tuple4[V1, V2, V3, V4]) GetField(field int) any {
	switch field {
	case 0:
		return t.V1
	case 1:
		return t.V2
	case 2:
		return t.V3
	case 3:
		return t.V4
	default:
		panic("wrong field")
	}
}

func (t *Tuple4[V1, V2, V3, V4]) New() Tuple {
	return &Tuple4[V1, V2, V3, V4]{}
}

// Tuple5
type Tuple5[V1, V2, V3, V4, V5 any] struct {
	V1 V1
	V2 V2
	V3 V3
	V4 V4
	V5 V5
}

func T5[V1, V2, V3, V4, V5 any](v1 V1, v2 V2, v3 V3, v4 V4, v5 V5) Tuple5[V1, V2, V3, V4, V5] {
	return Tuple5[V1, V2, V3, V4, V5]{V1: v1, V2: v2, V3: v3, V4: v4, V5: v5}
}

func (t *Tuple5[V1, V2, V3, V4, V5]) SetField(field int, val any) error {
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
	case 3:
		nval, ok := val.(V4)
		if !ok {
			return ErrWrongFieldType
		}
		t.V4 = nval
		return nil
	case 4:
		nval, ok := val.(V5)
		if !ok {
			return ErrWrongFieldType
		}
		t.V5 = nval
		return nil
	default:
		return ErrExtraField
	}
}

func (t *Tuple5[V1, V2, V3, V4, V5]) GetField(field int) any {
	switch field {
	case 0:
		return t.V1
	case 1:
		return t.V2
	case 2:
		return t.V3
	case 3:
		return t.V4
	case 4:
		return t.V5
	default:
		panic("wrong field")
	}
}

func (t *Tuple5[V1, V2, V3, V4, V5]) New() Tuple {
	return &Tuple5[V1, V2, V3, V4, V5]{}
}

// Tuple6
type Tuple6[V1, V2, V3, V4, V5, V6 any] struct {
	V1 V1
	V2 V2
	V3 V3
	V4 V4
	V5 V5
	V6 V6
}

func T6[V1, V2, V3, V4, V5, V6 any](v1 V1, v2 V2, v3 V3, v4 V4, v5 V5, v6 V6) Tuple6[V1, V2, V3, V4, V5, V6] {
	return Tuple6[V1, V2, V3, V4, V5, V6]{V1: v1, V2: v2, V3: v3, V4: v4, V5: v5, V6: v6}
}

func (t *Tuple6[V1, V2, V3, V4, V5, V6]) SetField(field int, val any) error {
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
	case 3:
		nval, ok := val.(V4)
		if !ok {
			return ErrWrongFieldType
		}
		t.V4 = nval
		return nil
	case 4:
		nval, ok := val.(V5)
		if !ok {
			return ErrWrongFieldType
		}
		t.V5 = nval
		return nil
	case 5:
		nval, ok := val.(V6)
		if !ok {
			return ErrWrongFieldType
		}
		t.V6 = nval
		return nil
	default:
		return ErrExtraField
	}
}

func (t *Tuple6[V1, V2, V3, V4, V5, V6]) GetField(field int) any {
	switch field {
	case 0:
		return t.V1
	case 1:
		return t.V2
	case 2:
		return t.V3
	case 3:
		return t.V4
	case 4:
		return t.V5
	case 5:
		return t.V6
	default:
		panic("wrong field")
	}
}

func (t *Tuple6[V1, V2, V3, V4, V5, V6]) New() Tuple {
	return &Tuple6[V1, V2, V3, V4, V5, V6]{}
}

// Tuple7
type Tuple7[V1, V2, V3, V4, V5, V6, V7 any] struct {
	V1 V1
	V2 V2
	V3 V3
	V4 V4
	V5 V5
	V6 V6
	V7 V7
}

func T7[V1, V2, V3, V4, V5, V6, V7 any](v1 V1, v2 V2, v3 V3, v4 V4, v5 V5, v6 V6, v7 V7) Tuple7[V1, V2, V3, V4, V5, V6, V7] {
	return Tuple7[V1, V2, V3, V4, V5, V6, V7]{V1: v1, V2: v2, V3: v3, V4: v4, V5: v5, V6: v6, V7: v7}
}

func (t *Tuple7[V1, V2, V3, V4, V5, V6, V7]) SetField(field int, val any) error {
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
	case 3:
		nval, ok := val.(V4)
		if !ok {
			return ErrWrongFieldType
		}
		t.V4 = nval
		return nil
	case 4:
		nval, ok := val.(V5)
		if !ok {
			return ErrWrongFieldType
		}
		t.V5 = nval
		return nil
	case 5:
		nval, ok := val.(V6)
		if !ok {
			return ErrWrongFieldType
		}
		t.V6 = nval
		return nil
	case 6:
		nval, ok := val.(V7)
		if !ok {
			return ErrWrongFieldType
		}
		t.V7 = nval
		return nil
	default:
		return ErrExtraField
	}
}

func (t *Tuple7[V1, V2, V3, V4, V5, V6, V7]) GetField(field int) any {
	switch field {
	case 0:
		return t.V1
	case 1:
		return t.V2
	case 2:
		return t.V3
	case 3:
		return t.V4
	case 4:
		return t.V5
	case 5:
		return t.V6
	case 6:
		return t.V7
	default:
		panic("wrong field")
	}
}

func (t *Tuple7[V1, V2, V3, V4, V5, V6, V7]) New() Tuple {
	return &Tuple7[V1, V2, V3, V4, V5, V6, V7]{}
}

// Tuple8
type Tuple8[V1, V2, V3, V4, V5, V6, V7, V8 any] struct {
	V1 V1
	V2 V2
	V3 V3
	V4 V4
	V5 V5
	V6 V6
	V7 V7
	V8 V8
}

func T8[V1, V2, V3, V4, V5, V6, V7, V8 any](v1 V1, v2 V2, v3 V3, v4 V4, v5 V5, v6 V6, v7 V7, v8 V8) Tuple8[V1, V2, V3, V4, V5, V6, V7, V8] {
	return Tuple8[V1, V2, V3, V4, V5, V6, V7, V8]{V1: v1, V2: v2, V3: v3, V4: v4, V5: v5, V6: v6, V7: v7, V8: v8}
}

func (t *Tuple8[V1, V2, V3, V4, V5, V6, V7, V8]) SetField(field int, val any) error {
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
	case 3:
		nval, ok := val.(V4)
		if !ok {
			return ErrWrongFieldType
		}
		t.V4 = nval
		return nil
	case 4:
		nval, ok := val.(V5)
		if !ok {
			return ErrWrongFieldType
		}
		t.V5 = nval
		return nil
	case 5:
		nval, ok := val.(V6)
		if !ok {
			return ErrWrongFieldType
		}
		t.V6 = nval
		return nil
	case 6:
		nval, ok := val.(V7)
		if !ok {
			return ErrWrongFieldType
		}
		t.V7 = nval
		return nil
	case 7:
		nval, ok := val.(V8)
		if !ok {
			return ErrWrongFieldType
		}
		t.V8 = nval
		return nil
	default:
		return ErrExtraField
	}
}

func (t *Tuple8[V1, V2, V3, V4, V5, V6, V7, V8]) GetField(field int) any {
	switch field {
	case 0:
		return t.V1
	case 1:
		return t.V2
	case 2:
		return t.V3
	case 3:
		return t.V4
	case 4:
		return t.V5
	case 5:
		return t.V6
	case 6:
		return t.V7
	case 7:
		return t.V8
	default:
		panic("wrong field")
	}
}

func (t *Tuple8[V1, V2, V3, V4, V5, V6, V7, V8]) New() Tuple {
	return &Tuple8[V1, V2, V3, V4, V5, V6, V7, V8]{}
}

// Tuple9
type Tuple9[V1, V2, V3, V4, V5, V6, V7, V8, V9 any] struct {
	V1 V1
	V2 V2
	V3 V3
	V4 V4
	V5 V5
	V6 V6
	V7 V7
	V8 V8
	V9 V9
}

func T9[V1, V2, V3, V4, V5, V6, V7, V8, V9 any](v1 V1, v2 V2, v3 V3, v4 V4, v5 V5, v6 V6, v7 V7, v8 V8, v9 V9) Tuple9[V1, V2, V3, V4, V5, V6, V7, V8, V9] {
	return Tuple9[V1, V2, V3, V4, V5, V6, V7, V8, V9]{V1: v1, V2: v2, V3: v3, V4: v4, V5: v5, V6: v6, V7: v7, V8: v8, V9: v9}
}

func (t *Tuple9[V1, V2, V3, V4, V5, V6, V7, V8, V9]) SetField(field int, val any) error {
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
	case 3:
		nval, ok := val.(V4)
		if !ok {
			return ErrWrongFieldType
		}
		t.V4 = nval
		return nil
	case 4:
		nval, ok := val.(V5)
		if !ok {
			return ErrWrongFieldType
		}
		t.V5 = nval
		return nil
	case 5:
		nval, ok := val.(V6)
		if !ok {
			return ErrWrongFieldType
		}
		t.V6 = nval
		return nil
	case 6:
		nval, ok := val.(V7)
		if !ok {
			return ErrWrongFieldType
		}
		t.V7 = nval
		return nil
	case 7:
		nval, ok := val.(V8)
		if !ok {
			return ErrWrongFieldType
		}
		t.V8 = nval
		return nil
	case 8:
		nval, ok := val.(V9)
		if !ok {
			return ErrWrongFieldType
		}
		t.V9 = nval
		return nil
	default:
		return ErrExtraField
	}
}

func (t *Tuple9[V1, V2, V3, V4, V5, V6, V7, V8, V9]) GetField(field int) any {
	switch field {
	case 0:
		return t.V1
	case 1:
		return t.V2
	case 2:
		return t.V3
	case 3:
		return t.V4
	case 4:
		return t.V5
	case 5:
		return t.V6
	case 6:
		return t.V7
	case 7:
		return t.V8
	case 8:
		return t.V9
	default:
		panic("wrong field")
	}
}

func (t *Tuple9[V1, V2, V3, V4, V5, V6, V7, V8, V9]) New() Tuple {
	return &Tuple9[V1, V2, V3, V4, V5, V6, V7, V8, V9]{}
}
