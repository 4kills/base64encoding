package base64encoding

import (
	"errors"
)

func (enc Encoder64) decode(s string) ([]byte, error) {
	if len(s) < 1 {
		return nil, errors.New("base64decoding error: string is empty")
	}

	bits, err := base64ToBits([]byte(s), enc.posMap)
	if err != nil {
		return nil, err
	}

	// cut away the first shifted byte
	return bits.bits[1:], nil
}

func base64ToBits(s, posMap []byte) (BitArray, error) {
	bitLen := 6                        // only need 6 bit for the numbers 0-63
	shift := 8 - (len(s) * bitLen) % 8 // shifting the bit so i can cut them away more easily later
	bits := NewBitArray(shift + len(s) * bitLen)

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

			bits.Set(shift+ curPart + j, newBit)
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
