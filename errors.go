package base64encoding

import "errors"

var (
	// ErrNotDistinct indicates that a code set was provided that has at least two times the same rune
	ErrNotDistinct = errors.New("base64encoding: characters in codeSet are not pairwise distinct")

	// ErrIllegalRune inidcates that the code set had runes that were not legal
	ErrIllegalRune = errors.New(`base64encoding error: illegal rune: at least one character of the provided code set
	is not an ASCII or extended ASCII character`)
)
