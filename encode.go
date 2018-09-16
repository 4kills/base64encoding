package base64encoding

import (
	"math"
	"strings"
)

func (enc Encoder64) encode(b []byte) string {
	return bitsToBase64(bytesToBits(b), enc.codeSet)
}

func bitsToBase64(bits []bool, code string) string {
	runs := len(bits) / 6
	remainder := len(bits) % 6

	var sS []string
	if remainder != 0 {
		sS = append(sS, dezToBase64(bitsToDez(bits[:remainder]), code))
	}

	for i := 0; i < runs; i++ {
		dez := bitsToDez(bits[remainder+6*i : (remainder+6)+6*i])
		sS = append(sS, dezToBase64(dez, code))
	}
	return strings.Join(sS, "")
}

func dezToBase64(num int, codeSet string) string {
	base := int(len(codeSet)) // 64
	var encoded []string

	if num == 0 {
		return string(codeSet[0])
	}
	for num != 0 {
		remainder := num % base
		encoded = append([]string{codeSet[remainder : remainder+1]}, strings.Join(encoded, ""))
		num = num / base
	}

	return strings.Join(encoded, "")
}

func bitsToDez(bits []bool) int { // 00 0110
	num := 0
	n := float64(len(bits)) - 1
	for _, val := range bits {
		if val {
			num += int(math.Pow(2, n))
		}
		n--
	}
	return num
}

func bytesToBits(u []byte) []bool {
	var bits []bool
	for _, val := range u {
		bitArr := byteToBits(val)
		bits = append(bits, bitArr[:]...)
	}
	return bits
}

func byteToBits(by byte) [8]bool {
	var bits [8]bool
	pos := 7
	for n := 0; n < 8; n++ {
		mask := math.Pow(2, float64(n))
		if (by & byte(mask)) != 0 {
			bits[pos] = true
		}
		pos--
	}
	return bits
}
