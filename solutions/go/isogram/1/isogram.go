package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	letters := []string{}

	for i := 0; i < len(word); i++ {
		char := string(word[i])

		// Skip spaces and hyphens
		if char == " " || char == "-" {
			continue
		}

		// Check if letter already exists in slice
		for _, letter := range letters {
			if letter == char {
				return false
			}
		}

		// Add letter to slice
		letters = append(letters, char)
	}

	return true
}
