package main

import (
	"fmt"
	"strings"
)

func reverse(input string) {
	// Массив из слов
	words := strings.Fields(input)

	for i := len(words) - 1; i >= 0; i-- {
		fmt.Print(words[i])
		if i > 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func main() {
	input := "snow dog sun"
	reverse(input)
}
