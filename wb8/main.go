package main

import (
	"fmt"
	"os"
	"strconv"
)

func setBit(num64 int64, i int, value int) int64 {
	switch value {
	case 1:
		// Устанавливаем i-й бит в 1
		num64 |= (1 << i)
	case 0:
		// Устанавливаем i-й бит в 0. Для этого нам нужна маска
		//   111011 (маска)
		// &
		//   011101
		// = 011001

		mask := int64(^(1 << i))
		num64 &= mask
	}
	return num64
}

func main() {

	if len(os.Args) < 4 {
		fmt.Println("Usage: go run [path] [num] [number of bit (0 to 63)] [value of bit (1 or 0)]")
		return
	}

	// Преобразование первого аргумента (число) в int64
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid number")
		return
	}

	// Преобразование номера бита в int
	i, err := strconv.Atoi(os.Args[2])
	if err != nil || i < 0 || i >= 64 {
		fmt.Println("Invalid num of bit")
		return
	}

	// Преобразование значение бита в int
	value, err := strconv.Atoi(os.Args[3])
	if err != nil || (value != 0 && value != 1) {
		fmt.Println("Invalid value of bit")
		return
	}
	num64 := int64(num)
	// Установка i-го бита
	num64 = setBit(num64, i, value)
	fmt.Printf("Result: %d (bin: %064b)\n", num64, num64)
}
