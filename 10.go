package main

import "fmt"

func main() {
	temp := [...]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	res := make(map[int][]float64)
	for _, v := range temp {
		// Отбрасываем первый разряд
		key := int(v/10) * 10
		// Проверяем на существование ключа; создаем или добавляем значение
		slice, ok := res[key]
		if !ok {
			slice = []float64{v}
			res[key] = slice
		} else {
			slice = append(slice, v)
			res[key] = slice
		}
	}
	for k, v := range res {
		fmt.Printf("%d: %v\n", k, v)
	}
}
