package base64encoding

import (
	"errors"
	"unicode"
)

// StandardCodeSet is the default, http safe code set of base64encoding
const StandardCodeSet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"

// Base64WebSet is the standard base64 set for encoding data(e.g. images) in html files.
// However, this is not secure for using in URLs due to the '/' character
const Base64WebSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// EasilyReadableCodeSet is a set with no characters that look alike (e.g. 0 & O, l & I)
const EasilyReadableCodeSet = "*)23456789abcdefghi_klmnopqrstuvwxyzABCDEFGH+JKLMNOPQRSTUVWXYZ-$"

// Encoder64 with multiple methods. Contains codeSet
type Encoder64 struct {
	codeSet string
	posMap []byte
}

// Decode decodes a given string and returns an error if the string is not in a correct format
func (enc Encoder64) Decode(encodedStr string) ([]byte, error) {
	return enc.decode(encodedStr)
}

// DecodeNum decodes a given string, converts it to an unsigned int64 and returns an error if
// the string is not in a correct format
func (enc Encoder64) DecodeNum(encodedStr string) (uint64, error) {
	b, err := enc.decode(encodedStr)
	if err != nil {
		return 0, err
	}

	var num uint64
	for i := 0; i < 8; i++ {
		num = num | uint64(b[i])
		if i != 7 {
			num = num << 8
		}
	}
	return num, nil
}

// Encode encodes a given byte array to base64
func (enc Encoder64) Encode(bytes []byte) string {
	return enc.encode(bytes)
}

// EncodeNum encodes a given uint64 to base64 by splitting the num into a byte array
func (enc Encoder64) EncodeNum(num uint64) string {
	var (
		b  byte
		by [8]byte
	)
	for i := 7; i >= 0; i-- {
		b = byte(num)
		num = num >> 8
		by[i] = b
	}
	return enc.encode(by[:])
}

// CodeSet returns the codeSet, which the encoder uses to, well, en/decode
func (enc Encoder64) CodeSet() string {
	return enc.codeSet
}

// New returns a new Encoder64 with the standard, http safe code set
func New() Encoder64 {
	enc, _ :=  newCustom(StandardCodeSet)
	return enc
}

// NewWeb returns a new Encoder64 with the base64web encoding set used for encoding data in html
func NewWeb() Encoder64 {
	enc, _ := newCustom(Base64WebSet)
	return enc
}

// NewCustom returns a new Encoder64 with the provided custom 64 character code set and an error
// if the set is unfit. Please note: the first (left most) character of the string will be the least
// significant character (0 * 64^n), while the last will be the most significant(63 * 64^n)
func NewCustom(codeSet string) (Encoder64, error) {
	return newCustom(codeSet)
}

func newCustom(code string) (Encoder64, error) {
	if len(code) != 64 {
		return Encoder64{}, errors.New("base64encoding: length of code set invalid; 64 characters required")
	}

	for _, v := range code {
		if v > unicode.MaxASCII {
			return Encoder64{}, ErrIllegalRune
		}
	}

	b := NewBitArray(unicode.MaxASCII + 1)
	for i := 0; i < len(code); i++ {
		if b.Get(int(code[i])) != false {
			return Encoder64{}, ErrNotDistinct
		}

		b.Set(int(code[i]), true)
	}

	return Encoder64{codeSet: code}, nil
}
