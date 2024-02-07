package main

import "fmt"

func main() {
	// Можно собрать в мапу, срез и *массив
	// *массив если количество известно заранее

	maps("cat", "cat", "dog", "cat", "tree")   // добавляем как ключи => повторов не будет
	slices("cat", "cat", "dog", "cat", "tree") // добавляем как значения, повторы будут
}

func maps(values ...string) {
	res := make(map[string]struct{})

	for _, v := range values {
		res[v] = struct{}{}
	}
	fmt.Printf("%v\n", res)
}

func slices(values ...string) {
	res := make([]string, len(values))
	for i, v := range values {
		res[i] = v
	}
	fmt.Printf("%v\n", res)
}
