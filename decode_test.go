package base64encoding

import (
	"math"
	"testing"
)

func BenchmarkEncoder64_DecodeNum(b *testing.B) {
	enc := New()
	num := enc.EncodeNum(math.MaxUint64)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		enc.DecodeNum(num)
	}
}

func BenchmarkEncoder64_DecodeLong(b *testing.B) {
	strLen := 640
	enc := New()
	str := ""
	for i := 0; i < strLen / 64; i++ {
		str += enc.codeSet
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		enc.Decode(str)
	}
}