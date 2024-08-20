package main

import (
	"errors"
	"fmt"
)

func binarySearch(arr []int, target int) (int, error) {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid, nil // Элемент найден, возвращаем его индекс и отсутствие ошибки
		}

		if arr[mid] < target {
			low = mid + 1 // Ищем в правой половине
		} else {
			high = mid - 1 // Ищем в левой половине
		}
	}

	return -1, errors.New("not found")
}

func main() {
	arr := []int{-10, -5, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := -1

	index, err := binarySearch(arr, target)

	if err != nil {
		fmt.Printf("Err: %v\n", err)
		return
	}

	fmt.Printf("Элемент %d под индексом %d.\n", target, index)
}
