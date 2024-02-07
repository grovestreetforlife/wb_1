package main

import (
	"fmt"
	"sync"
)

func main() {
	// arr - исходный массив
	arr := [5]int{2, 4, 6, 8, 10}

	fmt.Println("1 вар:")
	// Первый вариант
	wg := new(sync.WaitGroup)
	for i := 0; i < len(arr); i++ { //Можно и for _, v := range arr { передавая в рутину v
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Println(x * x)
		}(arr[i])
	}
	// Ожидание завершения горутин через WaitGroup
	wg.Wait()

	fmt.Println("2 вар:")

	// Второй вариант
	channel := make(chan int, 5)
	for _, v := range arr {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			channel <- x * x
		}(v)
	}
	wg.Wait()
	close(channel)
	for val := range channel {
		fmt.Println(val)
	}

	fmt.Println("3 вар:")

	// Третий вариант, если важен порядок
	out := [5]int{}
	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go func(x int, i int) {
			defer wg.Done()
			out[i] = x * x
		}(arr[i], i)
	}
	wg.Wait()
	for _, v := range out {
		fmt.Println(v)
	}
}
