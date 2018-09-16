package base64encoding

import (
	"errors"
	"math"
	"strings"
)

func (enc Encoder64) decode(s string) ([]byte, error) {
	if len(s) < 1 {
		return nil, errors.New("base64decoding error: string is empty")
	}

	bits, err := base64ToBits(s, enc.codeSet)
	if err != nil {
		return nil, err
	}

	return bitsToBytes(bits), nil
}

func bitsToBytes(bits []bool) []byte {
	bits = bits[len(bits)%8:]

	var b []byte
	for i := 0; i < len(bits); i += 8 {
		b = append(b, byte(bitsToDez(bits[i:i+8])))
	}

	return b
}

func base64ToBits(s, code string) ([]bool, error) {
	var bits []bool
	for _, val := range s {
		num, err := base64Decoding(string(val), code)
		if err != nil {
			return nil, err
		}

		bits = append(bits, numToBits(num)...)
	}
	return bits, nil
}

func numToBits(n int) []bool {
	bits := byteToBits(byte(n))
	return bits[8-6 : 8]
}

func base64Decoding(s, codeSet string) (int, error) {
	num := 0
	n := float64(len(s)) - 1
	for i := 0; i < len(s); i++ {
		index := strings.IndexByte(codeSet, s[i])
		if index == -1 {
			return num, errors.New(`base64decoding error: semantic error: 
			string was invalid, character not found in codeset`)
		}
		num += index * int(math.Pow(64, n))
	}
	return num, nil
}
