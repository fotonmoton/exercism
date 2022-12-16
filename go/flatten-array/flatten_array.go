package flatten

func Flatten(input interface{}) []interface{} {

	flatten := make([]interface{}, 0)

	if input == nil {
		return flatten
	}

	asserted, ok := input.([]interface{})

	if !ok {
		return append(flatten, input)
	}

	for _, val := range asserted {
		flatten = append(flatten, Flatten(val)...)
	}

	return flatten
}
