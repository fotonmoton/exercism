// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {

	getChars := func(str string) string {
		onlyChars := make([]string, 0)
		for _, chr := range str {
			if unicode.IsLetter(chr) {
				onlyChars = append(onlyChars, string(chr))
			}
		}

		return strings.Join(onlyChars, "")
	}

	removeSpaces := func(str string) string {
		return strings.NewReplacer(" ", "", "\t", "", "\n", "", "\r", "").Replace(str)
	}

	isSilence := func(str string) bool {
		return "" == str
	}

	isUpper := func(str string) bool {
		onlyChars := getChars(str)
		return onlyChars != "" && onlyChars == strings.ToUpper(onlyChars)
	}

	isQuestion := func(str string) bool {
		return strings.HasSuffix(str, "?")
	}

	withoutSpaces := removeSpaces(remark)

	switch {
	case isSilence(withoutSpaces):
		return "Fine. Be that way!"
	case isUpper(withoutSpaces) && isQuestion(withoutSpaces):
		return "Calm down, I know what I'm doing!"
	case isQuestion(withoutSpaces):
		return "Sure."
	case isUpper(withoutSpaces):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}
