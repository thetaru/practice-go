package main

import "fmt"

type I interface {
	M()
}

func main() {
	// ćăăȘă
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
