package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()

	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()

	defer c.mu.Unlock()
	return c.v[key]
}

// 複数の異なるgorutineが変数への書き込み時に競合しないようにするために排他制御が必要となる
// ただし、sync.Mutexには値渡しでなく、ポインタで渡すこと(コピーでなく実体を渡す)
func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
