package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var Walker func(t *tree.Tree)
	Walker = func(t *tree.Tree) {
		if t != nil {
			Walker(t.Left)
			ch <- t.Value
			Walker(t.Right)
		} else {
			return
		}
	}
	Walker(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for v1 := range ch1 {
		v2, ok := <-ch2
		if v1 != v2 || !ok {
			return false
		}
	}
	return true
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	fmt.Println("t1:", t1)
	fmt.Println("t2:", t2)
	fmt.Println(Same(t1, t2))
}
