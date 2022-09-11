package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("sequences not the same length")
	}

	var hamDist int

	for i, _ := range a {
		if a[i] != b[i] {
			hamDist++
		}
	}

	return hamDist, nil
}
