package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(check("abcd"))
	fmt.Println(check("abCdefAaf"))
	fmt.Println(check("aA"))
	fmt.Println(check("aabcd"))

}

func check(s string) bool {
	m := make(map[rune]bool)
	s = strings.ToLower(s)
	for _, v := range s {
		if m[v] {
			return false
		}
		m[v] = true
	}
	return true
}
