package parseint

type error interface {
	Error() string
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func Parse(s string) (int, error) {
	var result int
	if len(s) == 1 {
		if isNumbericDigit(s[0]) {
			return int(s[0] - '0'), nil
		} else {
			return result, &errorString{"Is not a number"}
		}
	} else {
		for i := range s {
			if isNumbericDigit(s[i]) {
				temp := int(s[i] - '0')
				result = (result * 10) + temp
			} else if s[i] == '-' {
				continue
			} else {
				return 0, &errorString{"Is not a number"}
			}
		}
	}
	if s[0] == '-' {
		return result * (-1), nil
	} else {
		return result, nil
	}
}

func isNumbericDigit(s byte) bool {
	return s >= '0' && s <= '9'
}

// TODO: TESTING
// TODO: BENCHMARK
// TODO: Cover
