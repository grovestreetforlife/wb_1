package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Получаем количество секунд
	var input int
	fmt.Print("Введите кол-во секунд перед завершением работы: ")
	fmt.Scan(&input)

	// Создаем таймер на заданное количество секунд
	timer := time.NewTimer(time.Duration(input) * time.Second)
	// Создаём канал для пересылки данных
	broadcast := make(chan int)
	// Отправляем в эфир
	go func() {
		for {
			time.Sleep(1 * time.Second)
			broadcast <- rand.Int()
		}
	}()
	// Получаем и выводим из канала
	go func() {
		for v := range broadcast {
			fmt.Println(v)
		}
	}()
	<-timer.C
}
