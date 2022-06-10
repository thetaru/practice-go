package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	// バッファの要素に対して、多くても少なくてもだめ
	ch <- 1
	ch <- 2
	//ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
