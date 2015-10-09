package Parse

func Parseint(s string) (int, error) {
	//var res int
	for i := range s {
		if s[i] <= '0' || s[i] >= '9' {
			return 0, newErr("Input should content only number")
		}
	}
	res := 0
	for j := 0; j < len(s); j++ {
		res = Convertleter(rune(s[j])) + res*10
	}
	return res, nil
}

func Convertleter(c rune) int {
	return int(c - '0')
}

type MyError struct {
	err string
}

func (m *MyError) Error() string {
	return m.err
}
func newErr(t string) error {
	return &MyError{t}
}
