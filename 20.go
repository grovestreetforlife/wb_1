package main

import (
	"fmt"
	"strings"
)

func reverse(str string) string {
	res := strings.Split(str, " ")

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return strings.Join(res, " ")
}

func main() {
	fmt.Println(reverse("snow dog sun"))
	fmt.Println(reverse("позавчера вчера сегодня завтра послезавтра"))
}
