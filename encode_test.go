package base64encoding

import (
	"math"
	"testing"
)

func BenchmarkEncoder64_EncodeNum(b *testing.B) {
	enc := New()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		enc.EncodeNum(math.MaxUint64)
	}
}

func BenchmarkEncoder64_EncodeLong(b *testing.B) {
	strLen := 640
	enc := New()
	str := ""
	for i := 0; i < strLen / 64; i++ {
		str += enc.codeSet
	}
	bytes := []byte(str)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		enc.Encode(bytes)
	}
}
