package godata

import "testing"

func TestBatch(t *testing.T) {
	batch := NewBatchOf(&Tuple3[int, int, string]{}, 5)
	batch.Data[0] = &Tuple3[int, int, string]{0, 42, "Hello world"}
	batch.Data[1] = &Tuple3[int, int, string]{1, 42, "Hello world"}
	batch.Data[2] = &Tuple3[int, int, string]{2, 42, "Hello world"}
	batch.Data[3] = &Tuple3[int, int, string]{3, 42, "Hello world"}
	batch.Data[4] = &Tuple3[int, int, string]{4, 42, "Hello world"}

	buf := make([]byte, 500)
	n1 := batch.Encode(buf)

	target_batch := NewBatchOf(&Tuple3[int, int, string]{}, 5)
	n2, err := target_batch.Decode(buf)

	if err != nil {
		t.Error(err)
	}
	if n1 != n2 {
		t.Errorf("incorrect deserialized byte count n1=%d, n2=%d", n1, n2)
	}
}
