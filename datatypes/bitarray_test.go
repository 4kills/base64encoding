package datatypes

import (
	"testing"
)

func BenchmarkGet(b *testing.B) {
	bits := NewBitArray(1024)
	size := bits.Len()
	for i := 0; i < b.N; i++ {
		bits.Get(i % size)
	}
}

func BenchmarkSet(b *testing.B) {
	bits := NewBitArray(1024)
	size := bits.Len()
	for i := 0; i < b.N; i++ {
		bits.Set(i%size, true)
	}
}
