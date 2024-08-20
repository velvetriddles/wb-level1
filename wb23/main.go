package main

import "fmt"

func remove(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}

func main() {
	arr := []int{1, 5, 3, 6, 8, 2, 4}
	fmt.Println(arr)
	arr = remove(arr, 2)
	fmt.Println(arr)
}
