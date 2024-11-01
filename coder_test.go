package godata

import (
	"testing"
)

func BenchmarkEncodeT1Int(b *testing.B) {
	t := Tuple1[int]{V1: 42}
	enc := CreateEncoderFunc(&t)
	buf := make([]byte, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc(&t, buf)
	}
}

func BenchmarkEncodeT2Int(b *testing.B) {
	t := Tuple2[int, int]{V1: 42, V2: 42}
	enc := CreateEncoderFunc(&t)
	buf := make([]byte, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc(&t, buf)
	}
}

func BenchmarkEncodeT3Int(b *testing.B) {
	t := Tuple3[int, int, int]{V1: 42, V2: 42, V3: 42}
	enc := CreateEncoderFunc(&t)
	buf := make([]byte, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc(&t, buf)
	}
}

func BenchmarkEncodeT1String(b *testing.B) {
	t := Tuple1[string]{V1: "hello"}
	enc := CreateEncoderFunc(&t)
	buf := make([]byte, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc(&t, buf)
	}
}

func BenchmarkEncodeT2String(b *testing.B) {
	t := Tuple2[string, string]{V1: "hello", V2: "hello"}
	enc := CreateEncoderFunc(&t)
	buf := make([]byte, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc(&t, buf)
	}
}

func BenchmarkEncodeT3String(b *testing.B) {
	t := Tuple3[string, string, string]{V1: "hello", V2: "hello", V3: "hello"}
	enc := CreateEncoderFunc(&t)
	buf := make([]byte, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc(&t, buf)
	}
}
