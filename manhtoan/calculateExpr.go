package main

import (
	//	"bufio"
	"fmt"
	"regexp"
	"strings"
)

const (
	ILLEGAL_TOKEN = iota
	EOF
	INLIT
	//FLOATLIT
	_IDENT
	_OPERATOR
	PARENTHESIS
)

var IN_LIT = regexp.MustCompile(`[0-9]+`)

var OPERATOR = regexp.MustCompile(`\+|-|/|\*`)

var IDENT = regexp.MustCompile(`[a-zA-Z]+`)

var SEPARATOR = regexp.MustCompile(`(|)`)

func isTOKEN(str string) int {
	if IN_LIT.MatchString(str) {
		return INLIT
	} else if OPERATOR.MatchString(str) {
		return _OPERATOR
	} else if IDENT.MatchString(str) {
		return _IDENT
	} else if SEPARATOR.MatchString(str) {
		return PARENTHESIS
	} else {
		return ILLEGAL_TOKEN
	}
}

func isSingleNumber(str string) bool {
	if str >= "0" && str <= "9" {
		return true
	}
	return false
}

func isIdent(str string) bool {
	if IDENT.MatchString(str) {
		return true
	}
	return false
}

func printToken(str string) {

	for i := 0; i < len(str); i++ {
		var tmp string
		count := 0

		// if space character -> skip
		flagSpace := strings.Compare(string(str[i]), " ")
		if flagSpace == 0 {
			continue
		}

		// Implement for number have length > 1
		if isSingleNumber(string(str[i])) {
			for j := i; j < len(str); j++ {
				if !isSingleNumber(string(str[j])) {
					break
				}
				tmp += string(str[j])
				count++
			}
		} else if isIdent(string(str[i])) { // Implement for Ident have length > 1
			for j := i; j < len(str); j++ {
				if !isIdent(string(str[j])) {
					break
				}
				tmp += string(str[j])
				count++
			}
		} else {
			tmp = string(str[i])
		}

		// Print token
		switch isTOKEN(tmp) {
		case INLIT:
			fmt.Println("INLIT", tmp)
		case _OPERATOR:
			fmt.Println("OPERATOR", tmp)
		case PARENTHESIS:
			fmt.Println("PARENTHESIS", tmp)
		case _IDENT:
			fmt.Println("IDENT", tmp)
		default:
			fmt.Println("ERROR TOKEN", tmp)
		}

		if count > 0 {
			i += count - 1
		}
	}

}

func main() {
	var str string = "1 + 2 * 3 * ( 4 - 2 ) + adadd / B + C"
	// fmt.Print("Please input expression: ")
	// scanner := bufio.NewScanner(strings.NewReader(str))
	// scanner.Scan()

	// Print token lexer
	printToken(str)

}
