package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, v := range strings.Split(s, " ") {
		_, ok := m[v]
		if !ok {
			m[v] = 1
		} else {
			m[v]++
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
