package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToLower(word)

	for i, c := range word {
		if c == ' ' || c == '-' {
			continue
		}

		for j := i + 1; j < len(word); j++ {
			if word[j] == word[i] {
				return false
			}
		}
	}

	return true
}
