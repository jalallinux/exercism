package scrabble

import "strings"

/*
	Letter	Value
	A, E, I, O, U, L, N, R, S, T	1
	D, G	2
	B, C, M, P	3
	F, H, V, W, Y	4
	K	5
	J, X	8
	Q, Z	10
*/

func Score(word string) int {
	score := 0
	for _, letter := range strings.ToUpper(word) {
		score += scoreForLetter(letter)
	}
	return score
}

func scoreForLetter(letter rune) int {
	switch letter {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	default:
		return 0
	}
}
