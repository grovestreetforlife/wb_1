package main

import "fmt"

// Human - родительская структура
type Human struct {
	Name string
	Age  int
}

// Buy - метод родительской структуры
func (h *Human) Buy() {
	fmt.Println("Покупаем акции")
}

// Sell - методы родительской структуры
func (h *Human) Sell() {
	fmt.Println("Фиксируем прибыль")
}

// Action - структура, со встроенными методами от родительской Human
type Action struct {
	Human
}

func (h *Action) Bustle() {
	fmt.Println("Произошла суета")
}

func main() {
	a := Action{
		Human{
			Name: "John",
			Age:  40,
		}}
	a.Buy()
	a.Sell()
	a.Bustle()
}
