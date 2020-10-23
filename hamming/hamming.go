package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	count := 0
	if len(a) != len(b) {
		return -1, errors.New("Invalid")
	}

	target := []rune(b)
	for i, c := range a {
		if c != target[i] {
			count++
		}
	}
	return count, nil
}
