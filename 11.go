package main

import "fmt"

func main() {
	safari := map[string]struct{}{
		"slow":     {},
		"a lot ad": {},
		"lag":      {},
		"bug":      {},
	}

	mozila := map[string]struct{}{
		"fast":  {},
		"no ad": {},
		"lag":   {},
		"bug":   {},
	}
	intersection := getIntersectionOfSets(safari, mozila)

	fmt.Printf("%v\n", intersection)
}

func getIntersectionOfSets(first, second map[string]struct{}) map[string]struct{} {
	intersection := make(map[string]struct{})
	// Провверяем наличие одинаковых ключей в мапах
	for k := range first {
		if _, ok := second[k]; ok {
			intersection[k] = struct{}{}
		}
	}
	return intersection
}
