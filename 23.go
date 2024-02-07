package main

import "fmt"

func main() {
	s := []int{2, 23, 312, 768, 22, 14, 98, 777, 345, 12345, 222, 555, 307}
	idx := 4

	if idx < len(s) && idx >= 0 {
		s = append(s[:idx], s[idx+1:]...)
	}
	fmt.Printf("%v\n", s)
}
