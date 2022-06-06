package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	// 基底となる値に応じたメソッド(ここではMメソッド)が実行される
	i = &T{"Hello"}
	descibe(i)
	i.M()

	i = F(math.Pi)
	descibe(i)
	i.M()
}

func descibe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
