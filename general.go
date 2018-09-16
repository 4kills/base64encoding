package base64encoding

import "errors"

// StandardCodeSet is the default, http safe code set of base64encoding
const StandardCodeSet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"

// Base64WebSet is the standard base64 set for encoding data(e.g. images) in html files.
// However, this is not secure for using in URLs due to the '/' character
const Base64WebSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwsyz0123456789+/"

// Encoder64 with multiple methods. Contains codeSet
type Encoder64 struct {
	codeSet string
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
		num = num << 8
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
	return Encoder64{codeSet: StandardCodeSet}
}

// NewWeb returns a new Encoder64 with the base64web encoding set used for encoding data in html
func NewWeb() Encoder64 {
	return Encoder64{codeSet: Base64WebSet}
}

// NewCustom returns a new Encoder64 with the provided custom 64 character code set and an error
// if the set is unfit. Please note: the first (left most) character of the string will be the least
// significant character (0 * 64^n), while the last will be the most significant(63 * 64^n)
func NewCustom(codeSet string) (Encoder64, error) {
	return newCustom(codeSet)
}

func newCustom(code string) (Encoder64, error) {
	if len(code) != 64 {
		err := errors.New("base64encoding error: length of code set invalid; 64 characters required")
		return Encoder64{}, err
	}

	for _, val := range code {
		if val > 255 {
			err := errors.New(`base64encoding error: illegal rune: at least one character of the provided code set
			is not an ASCII or extended ASCII character`)
			return Encoder64{}, err
		}
	}

	return Encoder64{codeSet: code}, nil
}
