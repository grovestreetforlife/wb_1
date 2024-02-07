package main

import "fmt"

func main() {
	a, b := "a", "b"
	a, b = b, a
	fmt.Printf("a = %s, b = %s\n", a, b)
}
