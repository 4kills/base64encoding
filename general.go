package base64encoding

import "errors"

// StandardCodeSet is the default, http safe code set of base64encoding
const StandardCodeSet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"

// Encoder64 with multiple methods. Contains codeSet
type Encoder64 struct {
	codeSet string
}

// CodeSet returns the codeSet, which the encoder uses to, well, en/decode
func (enc Encoder64) CodeSet() string {
	return enc.codeSet
}

// New returns a new Encoder64 with the standard, http safe code set
func New() Encoder64 {
	return Encoder64{codeSet: StandardCodeSet}
}

// NewCustom returns a new Encoder64 with the provided custom 64 character code set and an error
// if the set is unfit
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
