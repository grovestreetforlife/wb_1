package main

import (
	"math/rand"
	"strings"
)

var justString string
var primer = "abcdefghijklmnopqrstuvwxyz1234567890"

func createHugeString(l int) string {
	str := strings.Builder{}
	for i := 0; i < l; i++ {
		str.WriteByte(primer[rand.Intn(len(primer))])
	}
	return str.String()
}

func someFunc() {
	v := createHugeString(1 << 10)
	println(v)
	// justString = v[:100] - присваиваем подстроку по ссылке, из-за чего сборщик не почистит память т.к.
	// ссылаемся на большое значение

	// Здесь делаем копию нужной подстроки. При выходе из функции сборщик почистит память т.к. на v уже не ссылаемся
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
	println(justString)
}
