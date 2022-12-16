package secret

type CodeMap map[uint]string

var codemap = CodeMap{
	1 << 0: "wink",
	1 << 1: "double blink",
	1 << 2: "close your eyes",
	1 << 3: "jump",
}

func Handshake(code uint) []string {
	result := make([]string, 0, 4)
	reverse := code&(1<<4) != 0

	for bit, pharse := range codemap {
		switch {
		case code&bit == 0:
			continue
		case reverse:
			result = append([]string{pharse}, result...)
		default:
			result = append(result, pharse)
		}
	}

	return result
}
