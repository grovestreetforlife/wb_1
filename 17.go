package main

func main() {
	arr := []int{1, -10, 2, -9, -8, 3, -7, 4, -6, 5, -5, 6, -4, 7, -3, 8, -2, 9, -1, 10}
	val, ok := getIdx(8, arr)
	if ok {
		println(val)
	} else {
		println("Not Found")
	}
	val, ok = getIdx(11, arr)
	if ok {
		println(val)
	} else {
		println("Not Found")
	}
}

// Бинарный поиск
func getIdx(val int, arr []int) (int, bool) {
	start, end := 0, len(arr)-1

	for start <= end {
		midIndex := (start + end) / 2
		midElem := arr[midIndex]

		if midElem == val {
			return midIndex, true
		} else if midElem < val {
			start = midIndex + 1
		} else {
			end = midIndex - 1
		}
	}

	return 0, false
}
