package main

import (
	"fmt"
	"strings"
)

func reverse(str string) string {
	var s strings.Builder
	s.Grow(len(str)) // Мини оптимизация 

	for i := len(str) - 1; i >= 0; i-- {
		s.WriteByte(str[i])
	}

	return s.String()
}

func main() {
	str := "главрыба"
	res := reverse(str)
	fmt.Println(res, res == "абырвалг")
}
