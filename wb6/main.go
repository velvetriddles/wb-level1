package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

// Способ 1: Остановка горутины с использованием канала для сигнала завершения
func stopWithChannel() {
	fmt.Println("Starting stopWithChannel...")
	stopChan := make(chan struct{})

	go func() {
		for {
			select {
			case <-stopChan:
				fmt.Println("GoroutineCH is stopping with channel...")
				return
			default:
				fmt.Println("GoroutineCH is working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	close(stopChan) // Сигнализируем о завершении горутины
	time.Sleep(1 * time.Second)
	fmt.Println("Finished stopWithChannel.")
}

// Способ 2: Остановка горутины с использованием ctx, cancel
func stopWithContext() {
	fmt.Println("Starting stopWithContext...")
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("GoroutineWC is stopping due to context cancellation...")
				return
			default:
				fmt.Println("GoroutineWC is working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	cancel() // Отмена контекста для завершения горутины
	time.Sleep(1 * time.Second)
	fmt.Println("Finished stopWithContext.")
}

// Способ 3: Остановка горутины с использованием таймера. Так же есть разные вариации с помоьщью пакета time, но здесь конечно же я показал только один
func stopWithTimeAfter() {
	fmt.Println("Starting stopWithTimeAfter...")
	// tt := time.NewTimer()
	// tt2 := time.AfterFunc()

	timeout := time.After(2 * time.Second)
	go func() {
		for {
			select {
			case <-timeout:
				fmt.Println("GoroutineTA is stopping due to timeout...")
				return
			default:
				fmt.Println("GoroutineTA is working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(3 * time.Second) // Даём время на завершение горутины
	fmt.Println("Finished stopWithTimeAfter.")
}

// Способ 4: Остановка горутины с использованием атомарного флага - тоже сомнительная и редчайшее довольствие
func stopWithAtomicFlag() {
	fmt.Println("Starting stopWithAtomicFlag...")
	var stopFlag int32

	go func() {
		for {
			if atomic.LoadInt32(&stopFlag) == 1 {
				fmt.Println("GoroutineAF is stopping with atomic flag...")
				return
			}
			fmt.Println("GoroutineAF is working...")
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second)
	atomic.StoreInt32(&stopFlag, 1) // Устанавливаем флаг для завершения горутины
	time.Sleep(1 * time.Second)
	fmt.Println("Finished stopWithAtomicFlag.")
}

// Способ 5: Остановка всех горутин при задержке горутины. Сомнительное удовольствие. Только для учебных примеров
func stopWithTimeSleep() {
	fmt.Println("Starting stopWithProgramExit...")
	go func() {
		for {
			fmt.Println("GoroutinePE is working...")
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second) // Программа завершится, и все горутины остановятся
	fmt.Println("Finished stopWithProgramExit.")
}

func main() {
	// Способ 1: Остановка горутины с использованием канала
	stopWithChannel()

	// Способ 2: Остановка горутины с использованием context.Context
	stopWithContext()

	// Способ 3: Остановка горутины с использованием таймера time.After
	stopWithTimeAfter()

	// Способ 4: Остановка горутины с использованием атомарного флага
	stopWithAtomicFlag()

	// Способ 5: Остановка всех горутин при завершении программы
	stopWithTimeSleep()
}
