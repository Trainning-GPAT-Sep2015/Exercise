package main

import (
	//	"bufio"
	"fmt"
	"regexp"
	"strings"
)

/********* DEFINE STACK STRING ************/
type strStack struct {
	top   *strNode
	count int
}

type strNode struct {
	value string
	next  *strNode
}

func (s *strStack) length() int {
	return s.count
}

func (s *strStack) isEmpty() bool {
	return s.count == 0
}

func (s *strStack) Push(dataIn string) {
	//TODO
	if s.count == 0 {
		pNew := new(strNode)
		pNew.value = dataIn
		pNew.next = nil
		s.top = pNew
	} else {
		pNew := new(strNode)
		pNew.value = dataIn
		pNew.next = s.top
		s.top = pNew
	}
	s.count++
}

func (s *strStack) Pop() {
	//TODO
	if s.count > 0 {
		var pDel *strNode
		pDel = s.top
		s.top = pDel.next
		s.count--
	}
}

func (s *strStack) Top() string {
	//TODO
	return s.top.value
}

/********* END STACK ***************/

/********* DEFINE STACK INT ************/
type intStack struct {
	top   *intNode
	count int
}

type intNode struct {
	value int
	next  *intNode
}

func (s *intStack) length() int {
	return s.count
}

func (s *intStack) isEmpty() bool {
	return s.count == 0
}

func (s *intStack) Push(dataIn int) {
	//TODO
	if s.count == 0 {
		pNew := new(intNode)
		pNew.value = dataIn
		pNew.next = nil
		s.top = pNew
	} else {
		pNew := new(intNode)
		pNew.value = dataIn
		pNew.next = s.top
		s.top = pNew
	}
	s.count++
}

func (s *intStack) Pop() {
	//TODO
	if s.count > 0 {
		var pDel *intNode
		pDel = s.top
		s.top = pDel.next
		s.count--
	}
}

func (s *intStack) Top() int {
	//TODO
	return s.top.value
}

/********* END STACK ***************/

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

var SEPARATOR = regexp.MustCompile(`\(|\)`)

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

func lexerChecking(str string) {

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
			return
		}

		if count > 0 {
			i += count - 1
		}
	}
}

func getOperatorWeight(op string) int {
	weight := -1
	switch op {
	case "+", "-":
		weight = 1
	case "/", "*", "%":
		weight = 2
	case "^":
		weight = 3
	}
	return weight
}

func hasHigherPrecedence(op1, op2 string) bool {
	op1Weight := getOperatorWeight(op1)
	op2Weight := getOperatorWeight(op2)
	if op1Weight > op2Weight {
		return true
	} else {
		return false
	}
}

func calculateOperator(op string, operand1, operand2 int) int {
	switch op {
	case "+":
		return operand1 + operand2
	case "-":
		return operand1 - operand2
	case "*":
		return operand1 * operand2
	case "/":
		return operand1 / operand2
	default:
		return 0
	}
}

func convertToPostFix(_infix string) string {
	var infix string

	// Skip space in infix string
	for i := 0; i < len(_infix); i++ {
		if string(_infix[i]) == " " {
			continue
		}
		infix += string(_infix[i])
	}

	var postfix string
	myStack := new(strStack)
	for i := 0; i < len(infix); i++ {
		if i == 0 && string(infix[i]) == "-" {
			postfix += string(infix[i])
		} else {
			if string(infix[i]) == " " {
				continue
			} else if string(infix[i]) == "(" { // Left Parenthesis
				myStack.Push(string(infix[i]))
			} else if string(infix[i]) == ")" { // Right Parenthesis
				for !myStack.isEmpty() && string(infix[i]) != "(" {
					postfix += " "
					postfix += myStack.Top()
					myStack.Pop()
				}
				myStack.Pop()
			}
			switch isTOKEN(string(infix[i])) {
			case _OPERATOR:
				if string(infix[i]) == "-" && OPERATOR.MatchString(string(infix[i-1])) {
					postfix += string(infix[i])
				} else {
					postfix += " "
					if myStack.isEmpty() {
						myStack.Push(string(infix[i]))
					} else {
						if hasHigherPrecedence(string(infix[i]), myStack.Top()) {
							myStack.Push(string(infix[i]))
						} else {
							if postfix[len(postfix)-1] != 32 {
								postfix += " "
							}
							postfix += myStack.Top()
							postfix += " "
							myStack.Pop()
							myStack.Push(string(infix[i]))
						}
					}
				}
			case INLIT:
				postfix += string(infix[i])
			}
		}
	}
	for !myStack.isEmpty() {
		postfix += " "
		postfix += myStack.Top()
		myStack.Pop()
	}
	return postfix
}

func evaluatePostfix(postfix string) int {
	myIntStack := new(intStack)
	for i := 0; i < len(postfix); i++ {
		if string(postfix[i]) == " " {
			continue
		} else if IN_LIT.MatchString(string(postfix[i])) {

		} else if OPERATOR.MatchString(string(postfix[i])) {

		}
	}
}

func main() {
	var infix string = "-1 * -212 + 4 / -4 + 3"
	var postfix string
	// Print token lexer
	lexerChecking(infix)

	// Convert to postfix
	postfix = convertToPostFix(infix)
	fmt.Println(postfix)

}
