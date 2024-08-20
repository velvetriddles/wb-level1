package main

import (
	"fmt"
	"math"
)

func determineType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("int")
	case chan string:
		fmt.Println("string")
	default:
		fmt.Println("Hz 4e za tip")
	}
}

func main() {
	var a int = 14343
	var b string = "WB"
	var c bool = true
	var d chan int = make(chan int)
	var e chan string = make(chan string)
	var f float64 = math.Pi // попадет в дефолт

	determineType(a)
	determineType(b)
	determineType(c)
	determineType(d)
	determineType(e)
	determineType(f)
}
