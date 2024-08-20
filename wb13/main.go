package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 10

	// a' = a xor b
	// b = a' xor b
	// b = a xor (b xor b) = a xor 0 = a (Ассоциативность и идемпотентность)

	a = a ^ b // a теперь хранит XOR обоих чисел
	b = a ^ b // b теперь равно исходному значению a
	a = a ^ b // a теперь равно исходному значению b

	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
