package godata

import (
	"testing"
)

func BenchmarkEncodeT1Int(b *testing.B) {
	t := T1(42)
	enc := CreateEncoderFunc(&t)
	buf := make([]byte, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc(&t, buf)
	}
}

func BenchmarkDecodeT1Int(b *testing.B) {
	t := T1(42)
	enc := CreateEncoderFunc(&t)
	buf := make([]byte, 100)
	enc(&t, buf)

	nt := Tuple1[int]{}
	dec := CreateDecoderFunc(&nt)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dec(&nt, buf)
	}
}

func BenchmarkEncodeT2Int(b *testing.B) {
	t := T2(42, 42)
	enc := CreateEncoderFunc(&t)
	buf := make([]byte, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc(&t, buf)
	}
}

func BenchmarkDecodeT2Int(b *testing.B) {
	t := T2(42, 42)
	enc := CreateEncoderFunc(&t)
	buf := make([]byte, 100)
	enc(&t, buf)

	nt := Tuple2[int, int]{}
	dec := CreateDecoderFunc(&nt)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dec(&nt, buf)
	}
}

func BenchmarkEncodeT3Int(b *testing.B) {
	t := T3(42, 42, 42)
	enc := CreateEncoderFunc(&t)
	buf := make([]byte, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc(&t, buf)
	}
}

func BenchmarkDecodeT3Int(b *testing.B) {
	t := T3(42, 42, 42)
	enc := CreateEncoderFunc(&t)
	buf := make([]byte, 100)
	enc(&t, buf)

	nt := Tuple3[int, int, int]{}
	dec := CreateDecoderFunc(&nt)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dec(&nt, buf)
	}
}

func BenchmarkEncodeT1String(b *testing.B) {
	t := T1("hello")
	enc := CreateEncoderFunc(&t)
	buf := make([]byte, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc(&t, buf)
	}
}

func BenchmarkDecodeT1String(b *testing.B) {
	t := T1("hello")
	enc := CreateEncoderFunc(&t)
	buf := make([]byte, 100)
	enc(&t, buf)

	nt := Tuple1[string]{}
	dec := CreateDecoderFunc(&nt)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dec(&nt, buf)
	}
}

func BenchmarkEncodeT2String(b *testing.B) {
	t := T2("hello", "hello")
	enc := CreateEncoderFunc(&t)
	buf := make([]byte, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc(&t, buf)
	}
}

func BenchmarkDecodeT2String(b *testing.B) {
	t := T2("hello", "hello")
	enc := CreateEncoderFunc(&t)
	buf := make([]byte, 100)
	enc(&t, buf)

	nt := Tuple2[string, string]{}
	dec := CreateDecoderFunc(&nt)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dec(&nt, buf)
	}
}

func BenchmarkEncodeT3String(b *testing.B) {
	t := T3("hello", "hello", "hello")
	enc := CreateEncoderFunc(&t)
	buf := make([]byte, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc(&t, buf)
	}
}

func BenchmarkDecodeT3String(b *testing.B) {
	t := T3("hello", "hello", "hello")
	enc := CreateEncoderFunc(&t)
	buf := make([]byte, 100)
	enc(&t, buf)

	nt := Tuple3[string, string, string]{}
	dec := CreateDecoderFunc(&nt)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dec(&nt, buf)
	}
}
