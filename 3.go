package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	fmt.Printf("1: %v \n2: %v \n3: %v \n", frst(arr), scnd(arr), thrd(arr))
}

func frst(a [5]int) int {
	sum := 0
	wg := new(sync.WaitGroup)
	channel := make(chan int, 5)
	for i := 0; i < len(a); i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			channel <- x * x
		}(a[i])
	}
	wg.Wait()
	close(channel)
	for v := range channel {
		sum += v
	}
	return sum
}

func scnd(a [5]int) int {
	sum := 0
	wg := new(sync.WaitGroup)
	mux := new(sync.Mutex)
	for _, v := range a {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			res := x * x
			mux.Lock()
			sum += res
			mux.Unlock()
		}(v)
	}
	wg.Wait()
	return sum
}

func thrd(a [5]int) int {
	sum := 0
	wg := new(sync.WaitGroup)
	out := [5]int{}
	for i := 0; i < len(a); i++ {
		wg.Add(1)
		go func(x int, i int) {
			defer wg.Done()
			out[i] = x * x
		}(a[i], i)
	}
	wg.Wait()
	for _, v := range out {
		sum += v
	}
	return sum
}
