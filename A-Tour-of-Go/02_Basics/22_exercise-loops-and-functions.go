package main

import (
	"fmt"
)

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func Sqrt(x float64) float64 {
	var count int = 1
	var z float64 = 1.0
	var accuracy float64 = 0.0001
	var diff float64

	for count <= 10 {
		diff = (z*z - x) / (2 * z)
		if abs(diff) < accuracy {
			count++
		} else {
			count = 1
			z -= diff
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
