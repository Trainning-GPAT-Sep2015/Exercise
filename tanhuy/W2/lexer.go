package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type Token int

const (
	ILLEGAL_TOKEN = iota
	//EOF
	WS

	INTLIT
	OPERATOR
	VARIABLE
	PARENTHESIS
)

type Item struct {
	val string
	typ Token
}

var (
	IntergerLiteral = regexp.MustCompile("[0-9]+")
	Operator        = regexp.MustCompile(`\+|\*|-|/`)
	Space           = regexp.MustCompile(` |\t`)
	Variable        = regexp.MustCompile("A|B")
	Parenthesis     = regexp.MustCompile(`\(|\)`)
)

func getType(str string) int {
	if IntergerLiteral.MatchString(str) {
		return INTLIT
	} else if Operator.MatchString(str) {
		return OPERATOR
	} else if Space.MatchString(str) {
		return WS
	} else if Variable.MatchString(str) {
		return VARIABLE
	} else if Parenthesis.MatchString(str) {
		return PARENTHESIS
	} else {
		return ILLEGAL_TOKEN
	}

}

func getOperatorWeight(str string) int {
	weight := -1
	switch str {
	case "+":
		weight = 1
	case "-":
		weight = 1
	case "*":
		weight = 2
	case "/":
		weight = 2
	case "(":
		weight = 3
	case ")":
		weight = 3
	}
	return weight
}

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value string
	next  *Element
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Push(v string) {
	s.top = &Element{v, s.top}
	s.size++
}

func (s *Stack) Pop() string {
	var first string
	if s.size > 0 {
		first, s.top = s.top.value, s.top.next
		s.size--
		//return
	}
	return first
}

func infix2postfix(infix []string) []string {
	var postfix []string
	stack := Stack{}
	for i := range infix {
		if Operator.MatchString(infix[i]) {
			for stack.Size() > 0 && stack.top.value != "(" && getOperatorWeight(stack.top.value) >= getOperatorWeight(infix[i]) {
				postfix = append(postfix, stack.top.value)
				_ = stack.Pop()
			}
			stack.Push(infix[i])
		} else if IntergerLiteral.MatchString(infix[i]) || Variable.MatchString(infix[i]) {
			postfix = append(postfix, infix[i])
		} else if infix[i] == "(" {
			stack.Push(infix[i])
		} else if infix[i] == ")" {
			for stack.Size() > 0 && stack.top.value != "(" {
				postfix = append(postfix, stack.top.value)
				_ = stack.Pop()
			}
			_ = stack.Pop()
		}
	}
	for stack.Size() > 0 {
		postfix = append(postfix, stack.top.value)
		_ = stack.Pop()
	}
	return postfix
}

func evaluate(postfix []string) int {
	stack := Stack{}
	for i := range postfix {
		if IntergerLiteral.MatchString(postfix[i]) {
			stack.Push(postfix[i])
		} else if Operator.MatchString(postfix[i]) {
			operand2, _ := strconv.ParseInt(stack.Pop(), 10, 64)
			operand1, _ := strconv.ParseInt(stack.Pop(), 10, 64)

			switch postfix[i] {
			case "+":
				res := operand1 + operand2
				stack.Push(strconv.Itoa(int(res)))
			case "-":
				res := operand1 - operand2
				stack.Push(strconv.Itoa(int(res)))
			case "*":
				res := operand1 * operand2
				stack.Push(strconv.Itoa(int(res)))
			case "/":
				if operand2 != 0 {
					res := operand1 / operand2
					stack.Push(strconv.Itoa(int(res)))
				} else {
					panic("Divided by zero")
				}
			}
		}
	}
	result, _ := strconv.ParseInt(stack.Pop(), 10, 64)
	return int(result)
}

func main() {
	input := "11+2"
	fmt.Println("Input is: ", input)
	var after []string
	i := 0

	for i < len(input) {
		var temp string
		count := 0
		for j := i; j < len(input); j++ {
			if getType(string(input[j])) != getType(string(input[i])) {
				break
			}
			temp += string(input[j])
			count++
		}
		if getType(temp) != WS {
			after = append(after, temp)
		}
		switch getType(string(temp)) {
		case INTLIT:
			fmt.Printf("%s is %s\n", temp, "IntLit")
		case VARIABLE:
			fmt.Printf("%s is %s\n", temp, "Variable")
		case OPERATOR:
			fmt.Printf("%s is %s\n", temp, "Operator")
		case PARENTHESIS:
			fmt.Printf("%s is %s\n", temp, "Parenthesis")
		case WS:
		default:
			fmt.Printf("%s is %s\n", temp, "Illiteral token")
		}
		if count > 0 {
			i += count
		} else {
			i++
		}
	}
	fmt.Println("Infix: ", after)
	fmt.Println("Postfix: ", infix2postfix(after))
	fmt.Println("Evaluate: ", evaluate(infix2postfix(after)))
}
