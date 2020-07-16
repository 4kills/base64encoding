package base64encoding

import (
	"errors"
)

func (enc Encoder64) decode(s string) ([]byte, error) {
	if len(s) < 1 {
		return nil, errors.New("base64decoding error: string is empty")
	}

	bits, err := base64ToBits([]byte(s), []byte(enc.posMap))
	if err != nil {
		return nil, err
	}

	return bits.bits[1:], nil
}

func base64ToBits(s, posMap []byte) (BitArray, error) {
	bitLen := 6
	overflow := 8 - (len(s) * bitLen) % 8
	bits := NewBitArray(overflow + len(s) * bitLen)

	for i := 0; i < len(s); i++ {
		num, err := findValue(s[i], posMap)
		if err != nil {
			return BitArray{}, err
		}

		curPart := i * bitLen
		for j := 0; j < bitLen; j++ {
			newBit := false
			if (0x20 >> j) & num > 0 {
				newBit = true
			}

			bits.Set(overflow + curPart + j, newBit)
		}
	}

	return bits, nil
}

func findValue(s byte, posMap []byte) (int, error) {
	position := posMap[s]
	idx := int(position - 1)
	if position == 0 {
		return idx, errors.New("base64decoding: semantic: string was invalid, character not found in codeset")
	}
	return idx, nil
}
