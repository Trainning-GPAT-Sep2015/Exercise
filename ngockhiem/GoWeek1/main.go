package main

import (
	"fmt"
	"regexp"
	"strings"
	"strconv"
	"errors"
	"reflect"
)

type Token interface {
	String() string
}
type Expr []Token
type IntToken struct {
	value int
}
func (this IntToken) String() string {
	return strconv.Itoa(this.value)
}

type BracketToken struct {
	bracket string
}
func (this BracketToken) String() string {
	return string(this.bracket)
}

type OperatorToken struct {
	op string
}
func (this OperatorToken) String() string {
	return string(this.op)
}

type VariableToken struct {
	name string
}
func (this VariableToken) String() string {
	return string(this.name)
}

var (
	op = regexp.MustCompile(`^(\+|\-|\*|\/)$`)
	intval = regexp.MustCompile(`^([0-9]+)$`)
	variable = regexp.MustCompile(`^([a-z])$`)
	bracket = regexp.MustCompile(`^(\(|\))$`)
)

func GenerateToken(expr string) (tokenList []Token) {
	expr = expr + "/"
	current := 0
	for i, _ := range expr {
		_, e := CreateToken(expr[current:i + 1])
		if e != nil {
			lastToken, _ := CreateToken(expr[current:i])
			tokenList = append(tokenList, lastToken)
			current = i
		}

	}
	return
}
func CreateToken(s string) (t Token, err error) {
	switch {
	case intval.MatchString(s):
		value, _ := strconv.Atoi(s)
		t = IntToken{value: value}
	case bracket.MatchString(s):
		t = BracketToken{bracket: s}
	case variable.MatchString(s):
		t = VariableToken{name: s }
	case op.MatchString(s):
		t = OperatorToken{op: s }
	default:
		err = errors.New("Not valid")
	}
	return
}
func IsLower(op1, op2 OperatorToken) bool {
	if op1.op == "+" || op1.op == "-" {
		switch op2.op {
		case "+", "-":
			return false
		case "*", "/":
			return true
		}
	}
	return false
}

func Postfix(tokenList []Token) (output []Token) {
	stack := []Token{}
	_ = stack
	for _, t := range tokenList {
		switch t.(type) {
		case IntToken:
			output = append(output, t)
//			fmt.Println("IntToken ",stack, output)
		case OperatorToken:
			if len(stack) == 0 {
				stack = append(stack, t)
			} else if _, ok := stack[len(stack) - 1].(BracketToken); ok {
				// Neu dinh stack la ( thi push vao`
				stack = append(stack, t)
			} else {
				// pop ra toi khi gap op nho hon hoac gap (
				lower := IsLower(stack[len(stack) - 1].(OperatorToken), reflect.ValueOf(t).Interface().(OperatorToken))
				for ; len(stack) > 0 && !lower; {
					output = append(output, stack[len(stack) - 1])
					stack = stack[:len(stack) - 1]
					if len(stack) == 0 || reflect.TypeOf(stack[len(stack) - 1]).Name() == "BracketToken" {
						break
					}
					lower = IsLower(reflect.ValueOf(stack[len(stack) - 1]).Interface().(OperatorToken), reflect.ValueOf(t).Interface().(OperatorToken))
				}
				stack = append(stack, t)
			}
//			fmt.Println("OPToken ",stack, output)
		case VariableToken:
			output = append(output, t)
//			fmt.Println("VarToken ",stack, output)
		case BracketToken:
			bracket := reflect.ValueOf(t).FieldByName("bracket").String()
			if bracket == "(" {
				stack = append(stack, t)
			} else {
				for ; reflect.TypeOf(stack[len(stack)-1]).Name() != "BracketToken";  {
					output = append(output, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
				stack = stack[:len(stack) - 1]
			}
//			fmt.Println("BracketToken ",stack, output)
		default:
			panic("Postfix Error")
		}
	}
	for ; len(stack) > 0; {
		output = append(output, stack[len(stack) - 1])
		stack = stack[:len(stack) - 1]
	}
//	fmt.Println(stack, output)
	return
}

/**
	calculate value of a simple expression
	return IntToken
 */
func CalcValue(a IntToken, b OperatorToken, c IntToken)  (result IntToken) {
	switch b.op {
	case "+":
		result = IntToken{value: a.value + c.value}
	case "-":
		result = IntToken{value: a.value - c.value}
	case "*":
		result = IntToken{value: a.value * c.value}
	case "/":
		if ( c.value ==0) {
			panic("Divide by 0")
		}
		result = IntToken{value: a.value / c.value}
	}
	return
}

/**
	Calculate the expression result
 */
func CalcExpr(postfix []Token , vars map[string]int) Token {
	result := []Token{}
	_ = result
	for _,v := range postfix {
		switch reflect.TypeOf(v).Name() {
		case "IntToken":
			result = append(result, v)
		case "VariableToken":
			value, ok := vars[reflect.ValueOf(v).FieldByName("name").String()]
			if ok {
				result = append(result, IntToken{value: value})
			} else {
				panic("variable not in vars")
			}
		case "OperatorToken":
			if len(result) < 2 {
				panic("Not enough parameter in result stack")
			}
			newValue := CalcValue(reflect.ValueOf(result[len(result)-2]).Interface().(IntToken),v.(OperatorToken),reflect.ValueOf(result[len(result)-1]).Interface().(IntToken))
			result = result[:len(result)-2]
			result = append(result, newValue)
		}
	}
	if len(result) != 1 {
		panic("Something Wrong With The Calculation")
	}
	return result[0]
}

func main() {
	var expr string = "321-987"
	expr = strings.Replace(expr, " ", "", -1)
	fmt.Println("Expr : ",expr)
	a := GenerateToken(expr)
	fmt.Println("Postfix : ",Postfix(a))
	var vars map[string]int = map[string]int{"a": 25, "b": 99}
	fmt.Println("Expr Value : ",CalcExpr(Postfix(a), vars))
}
