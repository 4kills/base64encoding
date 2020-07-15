package base64encoding

import "testing"

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

func BenchmarkNewCustom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewCustom(StandardCodeSet)
	}
}
