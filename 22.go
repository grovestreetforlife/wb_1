package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, c big.Int
	var b string
	fmt.Println("Ввод чисел: a + b")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	switch b {
	case "+":
		fmt.Println(sum(&a, &c))
	case "-":
		fmt.Println(sub(&a, &c))
	case "*":
		fmt.Println(mul(&a, &c))
	case "/":
		val, ok := div(&a, &c)
		if ok {
			fmt.Println(val)
		} else {
			fmt.Println("Деление на ноль невозможно")
		}
	}

}

func sum(a, c *big.Int) *big.Int {
	return a.Add(a, c)
}

func sub(a, c *big.Int) *big.Int {
	return a.Sub(a, c)
}

func mul(a, c *big.Int) *big.Int {
	return a.Mul(a, c)
}

func div(a, c *big.Int) (*big.Int, bool) {
	if c.Cmp(big.NewInt(0)) == 0 {
		return nil, false
	}
	return a.Div(a, c), true
}
