package main

import "fmt"

const Pi = 3.14

func main() {
	const world = "World"
	fmt.Println("Hello", world)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
