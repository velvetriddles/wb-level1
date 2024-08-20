package main

import (
	"fmt"
	"os"
	"sort"
)

// Мапа как промежуточное хранилище
func mmap(set1, set2 []int) []int {
	result := []int{}
	m := make(map[int]struct{})

	// Заполняем мапку элементами первого множества
	for _, val := range set1 {
		m[val] = struct{}{}
	}

	// Проверяем элементы второго множества на наличие в мапе
	for _, val := range set2 {
		if _, ok := m[val]; ok {
			result = append(result, val)
			delete(m, val) // Удаляем, потому что могут попасться импостеры
		}
	}

	return result
}

// Two pointers - мощная тема, сам в шоке
func pointers(set1, set2 []int) []int {
	sort.Ints(set1)
	sort.Ints(set2)

	i, j := 0, 0
	result := []int{}

	for i < len(set1) && j < len(set2) {
		if set1[i] == set2[j] {
			// Проверка на дубликаты
			if len(result) == 0 || result[len(result)-1] != set1[i] {
				result = append(result, set1[i])
			}
			i++
			j++
		} else if set1[i] < set2[j] { // тут мы двигаем в зависимости от того в какой мапе элемент больше другого, потому что если
			// меньше в set1 то нет нам смысла проверять его с дальнейшими set2
			i++
		} else {
			j++
		}
	}

	return result
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run [path] [map or pointers]")
		return
	}

	method := os.Args[1]

	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}
	var result []int

	switch method {
	case "map":
		result = mmap(set1, set2)
	case "pointers":
		result = pointers(set1, set2)
	default:
		fmt.Println("Invalid method")
		return
	}

	fmt.Println("Result: ", result)
}
