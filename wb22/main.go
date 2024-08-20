package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(987765345432123)
	b := big.NewInt(987654321123456)

	sum := new(big.Int)
	sub := new(big.Int)
	mul := new(big.Int)
	div := new(big.Int)

	sum.Add(a, b)
	sub.Sub(a, b)
	mul.Mul(a, b)
	div.Quo(a, b)

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("a + b: ", sum)
	fmt.Println("a - b: ", sub)
	fmt.Println("a * b: ", mul)
	fmt.Println("a / b: ", div)
}
