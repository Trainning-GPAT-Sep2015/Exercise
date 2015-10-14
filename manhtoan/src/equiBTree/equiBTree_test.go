package equiBTree

import (
	"golang.org/x/tour/tree"
	"testing"
)

func TestWalk(t *testing.T) {
	testcasesWalk := []struct {
		in   *tree.Tree
		want []int
	}{
		{tree.New(1), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{tree.New(2), []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}},
	}
	for _, c := range testcasesWalk {
		ch1 := make(chan int)
		go Walk(c.in, ch1)
		for _, w := range c.want {
			got := <-ch1
			if got != w {
				t.Errorf("Expect %d, but return %d", w, got)
			}
		}
	}
}

func TestSame(t *testing.T) {
	testcasesSame := []struct {
		in1  *tree.Tree
		in2  *tree.Tree
		want bool
	}{
		{tree.New(1), tree.New(1), true},
		{tree.New(2), tree.New(1), false},
		{tree.New(5), tree.New(7), false},
		{tree.New(10), tree.New(10), true},
	}

	for _, c := range testcasesSame {
		got := Same(c.in1, c.in2)
		if got != c.want {
			t.Errorf("Expect %v, but return %v", c.want, got)
		}
	}
}

func BenchmarkWalk(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ch := make(chan int)
		go Walk(tree.New(2), ch)
		<-ch
	}
}

func BenchmarkSame(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Same(tree.New(10), tree.New(10))
	}
}
