package main

import (
	"fmt"
	"reflect"
)

func Parse(s string) (int, bool) {
	var result int
	if len(s) == 1 {
		if isNumbericDigit(s) {
			return int(s[0] - '0'), true
		} else {
			return result, false
		}
	} else {
		// for  i := range s
		for i := 0; i < len(s); i++ {
			if i == 0 && s[i] == '-' { // rune 2 byte
				i++
				if isNumbericDigit(string(s[i])) {
					temp := int(s[i] - '0')
					result = (result * 10) + temp
				} else {
					return 0, false
				}
			} else {
				if isNumbericDigit(string(s[i])) {
					temp := int(s[i] - '0')
					result = (result * 10) + temp
				} else {
					return 0, false
				}
			}
		}
	}
	if s[0] == '-' {
		return result * (-1), true
	} else {
		return result, true
	}
}

func isNumbericDigit(s string) bool {
	if s >= "0" && s <= "9" {
		return true
	}
	return false
}

func main() {
	s := "494"
	a, _ := Parse(s)
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
}

// TODO: TESTING
// TODO: BENCHMARK
// TODO: Cover
