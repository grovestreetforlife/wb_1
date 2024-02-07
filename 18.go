package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Add(val int) {
	c.mu.Lock()
	c.value += val
	c.mu.Unlock()
}

func main() {
	wg := sync.WaitGroup{}

	counter := Counter{
		mu:    sync.Mutex{},
		value: 0,
	}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Add(1)
		}()
	}

	wg.Wait()

	fmt.Printf("%d\n", counter.value)
}
