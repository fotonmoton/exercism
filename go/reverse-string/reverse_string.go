package reverse

import (
	"sort"
	"strings"
)

func Reverse(in string) string {
	splitted := strings.Split(in, "")
	sort.SliceStable(splitted, func(_, _ int) bool { return true })
	return strings.Join(splitted, "")
}
