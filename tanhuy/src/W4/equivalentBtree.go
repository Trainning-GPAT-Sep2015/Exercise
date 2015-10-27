package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	travel(t, ch)
	close(ch)

}
func travel(t *tree.Tree, ch chan int) {
	if t != nil {
		travel(t.Left, ch)
		ch <- t.Value
		travel(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	t := tree.New(2)
	ch := make(chan int)
	go Walk(t, ch)
	//close(ch)
	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println(Same(tree.New(2), tree.New(2)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
