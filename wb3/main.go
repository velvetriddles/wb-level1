package main

import (
	"fmt"
	"math"
	"sync"
)

/*
	=== Задача №3 ===

	Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….)
	с использованием конкурентных вычислений.

*/

// Первый способ
func first() {
	numbers := [5]int{2, 4, 6, 8, 10}
	//Канал для передачи квадратов
	squareChan := make(chan int, len(numbers))

	wg := new(sync.WaitGroup)

	// Цикл по элементам последовательности
	for _, num := range numbers {
		wg.Add(1)
		// Запуск горутины
		go func(n int, wg *sync.WaitGroup, squareChan chan<- int) {
			// Уменьшение счетчика горутин перед выходом из функции
			defer wg.Done()
			// Вычисление квадрата числа и передача результата в канал
			squareChan <- n * n
		}(num, wg, squareChan)
	}

	wg.Wait()
	close(squareChan)

	sum := 0
	// Цикл по значениям, полученным из канала squareChan
	for square := range squareChan {
		sum += square
	}
	fmt.Println("First method result:", sum)
}

// Второй способ - Мой любимый. PIPELINE THE BEST
func second() {
	array := [...]int{2, 4, 6, 8, 10}

	in := make(chan int)     // начальный канал для входящих данных
	result := make(chan int) // канал для результата
	wg := new(sync.WaitGroup)

	wg.Add(len(array))

	// Запускаем пайплайн
	go func() {
		defer close(result) // закрываем канал результата после завершения всех операций

		for i, value := range array { // проход по исходному массиву
			out := make(chan int) // каждый раз создаем свой выходной канал для новой горутины

			go func(in, out chan int, wg *sync.WaitGroup, value int) {
				defer wg.Done()
				defer close(out)                               // обязательно закрываем там, где пишем
				sum := <-in + int(math.Pow(float64(value), 2)) // подсчет квадрата и добавление к текущей сумме
				out <- sum                                     // записываем результат в выходной канал
			}(in, out, wg, value)

			in = out // передаем указатель на дескриптор выходного в входной канал, то есть входной канал для принимающей горутины является выходным отдающей горутины

			if i == len(array)-1 { // когда достигли последнего элемента, то есть остановка пайплайна
				result <- <-out // передаем результат в канал
			}
		}
	}()

	in <- 0   // запускаем цепочку передачи данных, 0 потому что изначально сумма равна нулю, сумма динамически растет в горутинах
	wg.Wait() // ждем завершения всех горутин

	fmt.Println("Second method result:", <-result)
}

func main() {
	// Вызов первого способа
	first()

	// Вызов второго способа
	second()
}
