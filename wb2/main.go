package main

import (
	"fmt"
	"math"
	"sync"
)

/*
	=== Задача №2 ===

	Написать программу, которая конкурентно рассчитает значение
	квадратов чисел взятых из массива (2,4,6,8,10) и выведет
	их квадраты в stdout.

	Реализованы два способа:
		1. Использование WaitGroup для синхронизации.
		2. Использование каналов для синхронщины
*/

// Способ 1: Использование sync.WaitGroup
func calculateWithWaitGroup() {
	// Исходный массив
	array := [...]int{2, 4, 6, 8, 10}

	// Создаем `WaitGroup`
	wg := new(sync.WaitGroup)

	// Увеличиваем счетчик на количество элементов в массиве.
	wg.Add(len(array))

	// Цикл по массиву
	for i, value := range array {
		// Для каждого элемента массива запускаем горутину
		go func(i, value int, wg *sync.WaitGroup) {
			// Уменьшаем счетчик после выхода из горутины
			defer wg.Done()

			// Вычисляем квадрат значения и записываем обратно в массив по соответствующему индексу
			array[i] = int(math.Pow(float64(value), 2))

		}(i, value, wg)
	}

	// Ожидаем выполнения всех горутин
	wg.Wait()

	fmt.Println(array)
}

// Способ 2: Использование каналов
func calculateWithChannels() {
	// Исходный массив
	array := []int{2, 4, 6, 8, 10}

	// Создаем канал для передачи результатов
	result := make(chan int)

	// Запускаем горутины для вычисления квадратов
	for _, value := range array {
		go func(value int) {
			result <- int(math.Pow(float64(value), 2))
		}(value)
	}

	// Считываем результаты из канала
	for range len(array) {
		fmt.Println(<-result)
	}
}

func main() {

	fmt.Println("Способ 1: Использование sync.WaitGroup")
	calculateWithWaitGroup()

	fmt.Println("Способ 2: Использование каналов")
	calculateWithChannels()
}
