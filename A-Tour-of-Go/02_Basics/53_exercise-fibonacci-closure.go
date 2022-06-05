package main

import "fmt"

func fibonacci() func() int {
	fib_n0 := 0
	fib_n1 := 1
	return func() int {
		fib_n1, fib_n0 = fib_n0, fib_n0+fib_n1
		return fib_n1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
