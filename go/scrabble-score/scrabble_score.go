package scrabble

import "strings"

var points = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10,
}

func Score(word string) int {
	score := 0
	for _, char := range word {
		for match, point := range points {
			if strings.Contains(match, strings.ToUpper(string(char))) {
				score += point
			}
		}
	}
	return score
}
