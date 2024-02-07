package main

import "fmt"

// Target интерфейс классов адаптера
type Target interface {
	Operation()
}

// Adaptee адаптируемый класс
type Adaptee struct {
}

// AdaptedOperation Метод адаптируемого класса, который нужно вызвать где-то
func (a *Adaptee) AdaptedOperation() {
	fmt.Println("I am AdaptedOperation()")
}

// ConcreteAdapter класс конкретного адаптера
type ConcreteAdapter struct {
	*Adaptee
}

// Operation реализация метода интерфейса, реализующего вызов адаптируемого класса
func (a *ConcreteAdapter) Operation() {
	a.AdaptedOperation()
}

// NewAdapter конструктор адаптера
func NewAdapter(a *Adaptee) Target {
	return &ConcreteAdapter{a}
}

// основной метод для демонстрации
func main() {
	fmt.Println("\nAdapter demo:\n")
	adapter := NewAdapter(&Adaptee{})
	adapter.Operation()
}
