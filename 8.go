package main

import "fmt"

// setOne устанавливает бит в позиции pos в 1
func setOne(n int64, pos uint) int64 {
	n |= 1 << pos // сдвигаем 1 на количество pos позиций влево, затем выполняем побитовое ИЛИ с n
	return n
}

// setZero устанавливает бит в позиции pos в 0
func setZero(n int64, pos uint) int64 {
	mask := ^(1 << pos) // сдвигаем 1 на количество pos позиций влево, инвертируем все биты, создаем маску
	n &= int64(mask)    // применяем маску к n, сбрасывая бит в позиции pos
	return n
}

func main() {
	var (
		num int64 // исходное число
		pos uint  // позиция
		n   int   // бит
	)
	fmt.Print("Введите число: ")
	fmt.Scanln(&num)

	fmt.Print("Введите позицию: ")
	fmt.Scanln(&pos)
	if pos > 64 {
		fmt.Printf("Позиция должна быть меньше 64\n")
		return
	}

	fmt.Print("Введите какой бит необходимо установить: ")
	fmt.Scanln(&n)
	if n != 1 || n != 0 {
		fmt.Printf("Бит должен быть равен 1 или 0\n")
		return
	}

	switch {
	case n == 1:
		res := setOne(num, pos)
		fmt.Printf("У числа %d, в битовом представлении %b, установлена 1 на позицию № %d. Результат: %d, в битовом представлении %b\n", num, num, pos, res, res)
	case n == 0:
		res := setZero(num, pos)
		fmt.Printf("У числа %d, в битовом представлении %b, установлен 0 на позицию № %d. Результат: %d, в битовом представлении %b\n", num, num, pos, res, res)
	}
}
