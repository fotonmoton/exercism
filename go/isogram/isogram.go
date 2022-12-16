package isogram

import "unicode"

func IsIsogram(word string) bool {
	found := map[rune]bool{}

	for _, char := range word {
		lower := unicode.ToLower(char)

		if unicode.IsLetter(lower) && found[lower] {
			return false
		}

		found[lower] = true
	}

	return true
}
