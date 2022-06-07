package main

import "fmt"

func main() {
	// 空のインターフェース: すべての型の互換性をもたせる
	var i interface{}
	describe(i)

	// int
	i = 42
	describe(i)

	// string
	i = "hello"
	describe(i)
}

// 引数の型を空のインターフェースにすることで、任意の型で受けられる
// ただし、引数の元の型を忘れることに注意(型assertion 型switchで対応する)
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
