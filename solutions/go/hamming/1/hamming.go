package hamming

import (
	"errors"
	"unicode/utf8"
)

// ErrLengthMismatch is returned when the input strings contain a different number of runes.
var ErrLengthMismatch = errors.New("hamming: strings must be of equal length")

// Distance calculates the Hamming distance between two strings.
// It compares Unicode code points (runes) to correctly handle UTF-8 encoded text.
// Returns ErrLengthMismatch if the strings do not contain the same number of runes.
func Distance(a, b string) (int, error) {
	var distance int
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		r1, size1 := utf8.DecodeRuneInString(a[i:])
		r2, size2 := utf8.DecodeRuneInString(b[j:])

		if r1 != r2 {
			distance++
		}

		i += size1
		j += size2
	}

	// If either index has not reached the end of its respective string,
	// the strings have a different number of runes.
	if i < len(a) || j < len(b) {
		return 0, ErrLengthMismatch
	}

	return distance, nil
}