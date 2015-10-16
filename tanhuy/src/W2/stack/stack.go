package stack

// import (
// 	"fmt"
// )

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

// func main() {
// 	stack := new(Stack)

// 	stack.Push("Things")
// 	stack.Push("and")
// 	stack.Push("Stuff")

// 	for stack.Size() > 0 {
// 		// We have to do a type assertion because we get back a variable of type
// 		// interface{} while the underlying type is a string.
// 		fmt.Printf("%s ", stack.Pop())
// 	}
// 	fmt.Println()
// }
