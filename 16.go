package main

import "fmt"

func main() {
	arr := []int{0, 13, 54, 2, 1, 543, 3, 6, 8, 223, 148, 432, 65, 999, 1211, 333, 444, 525}
	arr = quickSort(arr)

	fmt.Println(arr)
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	first := arr[0]
	var less, greater []int

	for i := 1; i < len(arr); i++ {
		val := arr[i]

		if val < first {
			less = append(less, val)
		} else {
			greater = append(greater, val)
		}
	}

	result := append(quickSort(less), first)
	result = append(result, quickSort(greater)...)

	return result
}
