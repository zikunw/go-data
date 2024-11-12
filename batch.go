package godata

// A batch contains multiple tuples
// T must be a tuple type
type Batch[T Tuple] struct {
	encode TupleEncoderFunc
	decode TupleDecoderFunc

	Data []T
}

// NewBatch creates a new batch of a specific type of tuple
func NewBatchOf[T Tuple](t T, size uint) Batch[T] {
	encode := CreateEncoderFunc(t)
	decode := CreateDecoderFunc(t)
	data := make([]T, size)
	for i := range size {
		data[i] = data[i].New().(T)
	}

	return Batch[T]{
		encode: encode,
		decode: decode,
		Data:   data,
	}
}

// Encode a batch into a byte buffer
func (b *Batch[T]) Encode(buf []byte) int {
	offset := 0
	for _, tuple := range b.Data {
		offset += b.encode(tuple, buf[offset:])
	}
	return offset
}

// Decode a batch from a byte buffer
// This will read until either
// 1. The byte buffer is empty OR
// 2. The buffer is full OR
// 3. An error has occured during decoding
//
// This function will return the number of bytes read
// from the buffer.
func (b *Batch[T]) Decode(buf []byte) (int, error) {
	offset := 0
	for _, tuple := range b.Data {
		n, err := b.decode(tuple, buf[offset:])
		offset += n
		if err != nil {
			return offset, err
		}
	}
	return offset, nil
}
