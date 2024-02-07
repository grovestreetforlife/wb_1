package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	firstChan := make(chan int)
	secondChan := make(chan int)

	for _, v := range arr {
		go func(v int, firstChan chan<- int) {
			firstChan <- v

		}(v, firstChan)
	}

	go func(firstChan <-chan int, secondChan chan<- int) {
		for v := range firstChan {
			secondChan <- v * 2
		}
	}(firstChan, secondChan)

	go func(secondChan <-chan int) {
		for v := range secondChan {
			println(v)
		}
	}(secondChan)
	<-interrupt
	os.Exit(0)
}
