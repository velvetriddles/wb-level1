package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	=== Задача №7 ===

	Реализовать конкурентную запись данных в map.

*/

// Я еще думал сильно о том как сделать мапу Lock Free - но сдаюсь, это сложновато. Поэтому что есть, то есть

type SafeMap struct {
	mu sync.RWMutex
	m  map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]int),
	}
}

func (s *SafeMap) Set(key string, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func (s *SafeMap) Get(key string) (int, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.m[key]
	return val, ok
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // Инициализация генератора случайных чисел

	safeMap := NewSafeMap()
	wg := &sync.WaitGroup{}

	// Запуск горутин для записи данных в мапу
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				key := fmt.Sprintf("key%d", j)
				randomVal := r.Intn(15)
				safeMap.Set(key, randomVal)
				fmt.Printf("Goroutine %d: Set %s = %d\n", i, key, randomVal)
				time.Sleep(1000 * time.Millisecond) // Имитация работы
			}
		}(i)
	}

	// Запуск горутин для чтения данных из мапы
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				key := fmt.Sprintf("key%d", j)
				val, ok := safeMap.Get(key)
				if ok {
					fmt.Printf("Goroutine %d: Get %s = %d\n", i+5, key, val)
				} else {
					fmt.Printf("Goroutine %d: %s not found\n", i+5, key)
				}
				time.Sleep(1000 * time.Millisecond) // Имитация работы
			}
		}(i)
	}

	// Ожидание завершения всех горутин
	wg.Wait()

	// Дополнительно: Вывод всех значений в мапе после завершения всех операций
	fmt.Println("\nValues in map: ")
	// for k, v := range safeMap.m {
	// 	fmt.Println()
	// }
	fmt.Println(safeMap.m)
}
