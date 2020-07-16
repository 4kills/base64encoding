package datatypes

import "errors"

// BitArray is a fast, memory-efficient implementation of a bit (bool) array.
// Use the NewBitArray function to create one.
type BitArray struct {
	size int
	bits []byte
}

// NewBitArray creates and returns a new BitArray with the specified size
func NewBitArray(size int) BitArray {
	s := size / 8
	if size%8 != 0 {
		s++
	}

	b := make([]byte, s)
	return BitArray{size, b}
}

// FromBytes creates a new BitArray with the provided bytes as initial values.
// Beware that the src slice is *NOT* copied, so mutating it will change the inner state of this Array.
func FromBytes(src []byte) BitArray {
	return BitArray{len(src)*8, src}
}

// Len returns the len/size of the array
func (b BitArray) Len() int {
	return b.size
}

// Get returns the bit at the specified index.
// false = 0. true = 1.
func (b BitArray) Get(index int) bool {
	checkInRange(b, index)

	partition := b.bits[index/8]
	offset := 0x80 >> (index % 8)

	result := partition & byte(offset)
	return result > 0
}

// Set allows you the set the bit at the specified index.
// false = 0. true = 1.
func (b BitArray) Set(index int, value bool) {
	checkInRange(b, index)

	partIdx := index / 8
	partition := b.bits[partIdx]

	val := 0x80 >> (index % 8)

	if !value {
		val = ^val
		b.bits[partIdx] = partition & byte(val)
		return
	}

	b.bits[partIdx] = partition | byte(val)
}

func checkInRange(b BitArray, index int) {
	if index < 0 || b.size <= index {
		panic(errors.New("bitarray: index out of range"))
	}
}
