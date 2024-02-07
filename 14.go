package main

import "fmt"

func main() {
	c := make(chan interface{})
	typeIdentifier(c)
	close(c)
	typeIdentifier(true)
	typeIdentifier(1414)
	typeIdentifier("string")
	typeIdentifier(nil)
}

func typeIdentifier(in interface{}) {
	switch in.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan interface{}:
		fmt.Println("channel")
	default:
		fmt.Println("unknown")
	}
}
