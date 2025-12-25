package hamming

import "errors"

func Distance(a, b string) (int, error) {
	// Handle edge case: different lengths
	if len(a) != len(b) {
		return 0, errors.New("different lengths")
	}

	var distance int = 0

	// Compare each character at the same position
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
