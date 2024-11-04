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

// Tuple4
type Tuple4[V1, V2, V3, V4 any] struct {
	V1 V1
	V2 V2
	V3 V3
	V4 V4
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

// Tuple5
type Tuple5[V1, V2, V3, V4, V5 any] struct {
	V1 V1
	V2 V2
	V3 V3
	V4 V4
	V5 V5
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

// Tuple6
type Tuple6[V1, V2, V3, V4, V5, V6 any] struct {
	V1 V1
	V2 V2
	V3 V3
	V4 V4
	V5 V5
	V6 V6
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
