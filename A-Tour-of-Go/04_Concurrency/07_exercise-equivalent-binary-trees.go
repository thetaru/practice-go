package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {

	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	v1, ok1 := <-c1
	v2, ok2 := <-c2

	if !ok1 || !ok2 {
		return false
	}

	if v1 != v2 {
		return false
	}

	return true
}

func main() {
	ch := make(chan int)

	// 再帰関数でチャネルをクローズする方法(クロージャ)
	go func() {
		defer close(ch)
		Walk(tree.New(1), ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
