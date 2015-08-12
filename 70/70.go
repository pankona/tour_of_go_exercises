package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
)

func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	if t.Left != nil {
		walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		walk(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for v := range ch1 {
		if v != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}

	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("ok")
	} else {
		fmt.Println("ng")
	}

	if !Same(tree.New(1), tree.New(2)) {
		fmt.Println("ok")
	} else {
		fmt.Println("ng")
	}
}
