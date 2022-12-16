package stringset

import (
	"fmt"
	"strings"
)

type Set map[string]bool

func (s Set) String() string {
	keys := make([]string, 0, len(s))

	for key := range s {
		keys = append(keys, fmt.Sprintf("\"%s\"", key))
	}

	return fmt.Sprintf("{%s}", strings.Join(keys, ", "))
}

func (s Set) Add(str string) {
	s[str] = true
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(str string) bool {
	return s[str]
}

func Subset(needle, haystack Set) bool {

	for str := range needle {
		if !haystack.Has(str) {
			return false
		}
	}

	return true
}

func Disjoint(needle, haystack Set) bool {

	for str := range needle {
		if haystack.Has(str) {
			return false
		}
	}

	return true
}

func Equal(needle, haystack Set) bool {

	if len(needle) != len(haystack) {
		return false
	}

	return Subset(needle, haystack)
}

func Intersection(needle, haystack Set) Set {

	set := Set{}

	for str := range needle {
		if haystack.Has(str) {
			set.Add(str)
		}
	}

	return set
}

func Difference(needle, haystack Set) Set {

	set := Set{}

	for str := range needle {
		if !haystack.Has(str) {
			set.Add(str)
		}
	}

	return set
}

func Union(a, b Set) Set {
	union := Set{}

	for key := range a {
		union.Add(key)
	}

	for key := range b {
		union.Add(key)
	}

	return union
}

func New() Set {
	return Set{}
}

func NewFromSlice(slice []string) Set {
	set := Set{}

	for _, str := range slice {
		set.Add(str)
	}

	return set
}
