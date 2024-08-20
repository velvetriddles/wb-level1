package main

import (
	"fmt"
)

// Реализация быстрой сортировки (quicksort)
func quicksort(arr []int) {
	if len(arr) < 2 {
		return
	}

	// Выбираем опорный элемент
	pivot := arr[len(arr)/2]

	// Делим массив на три части: меньше, равно и больше опорного элемента
	low, high := 0, len(arr)-1

	for low <= high {
		// Находим элементы, которые нужно поменять местами
		for arr[low] < pivot {
			low++
		}
		for arr[high] > pivot {
			high--
		}

		if low <= high {
			// Меняем элементы местами
			arr[low], arr[high] = arr[high], arr[low]
			low++
			high--
		}
	}

	quicksort(arr[:high+1])
	quicksort(arr[low:])
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("Исходный массив:", arr)
	quicksort(arr)
	fmt.Println("Отсортированный массив:", arr)
}
