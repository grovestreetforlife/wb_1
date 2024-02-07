package main

import (
	"context"
	"fmt"
	"time"
)

// Горутина завершится автоматически, когда выполнение функции будет завершено
func stopByFunctionEnd() {
	fmt.Println("Так и живём")
}

// Горутина завершится, когда получит сигнал из канала
func stopByChannel(stop <-chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("Канал скомпрометирован, сворачиваемся")
			return
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// Горутина завершится, когда получит сигнал из контекста
func stopByContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Контекст сказал хватит")
			return
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	go stopByFunctionEnd()

	stopChan := make(chan bool)
	go stopByChannel(stopChan)
	time.Sleep(1 * time.Second)
	stopChan <- true

	ctx, cancel := context.WithCancel(context.Background())
	go stopByContext(ctx)
	time.Sleep(1 * time.Second)
	cancel()
}
