package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var numOfWorkers, stopped int

	// Создаем контекст с таймаутом
	ctx, cancel := context.WithCancel(context.Background())

	// Получаем количество воркеров
	fmt.Print("Введите количество воркеров: ")
	fmt.Scanln(&numOfWorkers)

	// Создаем каналы для пересылки данных
	channel := make(chan int, numOfWorkers)
	stop := make(chan bool, numOfWorkers)
	interruptChannel := make(chan os.Signal, 1)

	// Подписываемся на сигнал
	signal.Notify(interruptChannel, os.Interrupt, syscall.SIGTERM)

	// Запускаем воркеры
	for i := 0; i < numOfWorkers; i++ {
		go runWorker(i, channel, stop, ctx)
	}
	// Запускаем горутины
	go func() {
		select {
		case <-interruptChannel:
			fmt.Println("получен сигнал на остановку")
			cancel()
		}
	}()

	for stopped < numOfWorkers {
		select {
		case <-stop:
			stopped++
		default:
			if stopped == 0 {
				data := rand.Int()
				channel <- data
			}
		}
	}
	fmt.Println("все воркеры остановлены")
	close(channel)
	close(stop)
	os.Exit(1)
}

func runWorker(id int, c <-chan int, stop chan bool, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("воркер", id, "остановлен")
			stop <- true
			return
		case n := <-c:
			fmt.Println("воркер", id, "получил", n)
		}
	}
}
