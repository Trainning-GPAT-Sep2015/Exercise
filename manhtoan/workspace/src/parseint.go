package main

import (
	"fmt"
	"reflect"
)

func Parse(s string) (int, bool) {
	var result int
	if len(s) == 1 {
		if isNumbericDigit(s) {
			return convToInt(s), true
		} else {
			return result, false
		}
	} else {
        for 
		for i := 0; i < len(s); i++ {
			if i == 0 && string(s[i]) == "-" { // rune 2 byte
				i++
				if isNumbericDigit(string(s[i])) {
					temp := convToInt(string(s[i]))
					result = (result * 10) + temp
                    // s[i] - int('0') -> int
				} else {
					return 0, false
				}
			} else {
				if isNumbericDigit(string(s[i])) {
					temp := convToInt(string(s[i]))
					result = (result * 10) + temp
				} else {
					return 0, false
				}
			}
		}
	}
	if string(s[0]) == "-" {
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

func convToInt(c string) int {
	switch c {
	case "0":
		return 0
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	}
	return -1
}

func main() {
	s := "-2323"
	a, _ := Parse(s)
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
}


// TODO: TESTING
// TODO: BENCHMARK
// TODO: Cover  