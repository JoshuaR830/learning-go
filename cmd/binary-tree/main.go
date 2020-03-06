package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk the tree
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Compare 2 trees for equivalence
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		x, y := <-ch1, <-ch2
		if x != y {
			return false
		}
	}

	return true
}

func printTree(t *tree.Tree) {
	ch := make(chan int)
	go Walk(t, ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(2)

	fmt.Println("First tree")
	printTree(t1)
	fmt.Println("Second tree")
	printTree(t2)

	fmt.Println("Compare trees")
	if Same(t1, t2) {
		fmt.Println("The trees are equal")
	} else {
		fmt.Println("The trees are different")
	}
}
