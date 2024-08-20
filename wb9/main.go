package main

import (
	"fmt"
)

/*
	=== Задача №9 ===

	Разработать конвейер чисел. Даны два канала: в первый пишутся
	числа (x) из массива, во второй — результат операции x*2,
	после чего данные из второго канала должны выводиться в stdout.

*/

// И все без wg Wait. Отлично просто

func main() {
	nums := make(chan int)
	results := make(chan int)

	arr := []int{1, 2, 3, 4, 5}

	// Запись чисел в первый канал
	go func() {
		for _, x := range arr {
			nums <- x
		}
		close(nums)
	}()

	// Чтение из первого канала, умножения и записи во второй канал
	go func() {
		for x := range nums {
			results <- x * 2
		}
		close(results) // Закрываем канал после отправки всех данных
	}()

	// Чтение из второго и вывод
	for result := range results {
		fmt.Println(result)
	}
}
