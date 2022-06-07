package main

import "fmt"

type I interface {
	M()
}

func main() {
	// 型がない
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
