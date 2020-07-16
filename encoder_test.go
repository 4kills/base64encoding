package base64encoding

import (
	"math"
	"testing"
)

func TestNewCustom_PairwiseDistinction(t *testing.T) {
	_, err := NewCustom(EasilyReadableCodeSet)
	if err != nil {
		t.Error(err)
	}

	// Easily readable code set with 2 runes alike
	//                             |                                   |
	//                             v                                   v
	faulty := "*)23456789abcdefghi_klmnopqrstuvwxyzABCDEFGH+JKLMNOPQRSTkVWXYZ-$"
	_, err = NewCustom(faulty)
	if err != ErrNotDistinct {
		t.Error(err)
	}
}

func TestEncoder64_Num(t *testing.T) {
	num := uint64(math.MaxUint64)
	enc := New()
	encoded := enc.EncodeNum(num)

	actual, err := enc.DecodeNum(encoded)
	if err != nil {
		t.Error(err)
	}

	if actual != num {
		t.Errorf("unequal: want: %d, got: %d", num, actual)
	}
}

func BenchmarkNewCustom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewCustom(StandardCodeSet)
	}
}
