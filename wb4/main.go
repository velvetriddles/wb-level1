package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Получаем количество воркеров из аргументов командной строки
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run [path] [numWorkers]")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("Invalid number of workers")
		return
	}

	// Создаем контекст с отменой по сигналу ctrl + c - SIGINT
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Настраиваем канал для получения системных сигналов
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	data := make(chan int)
	wg := new(sync.WaitGroup)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Запуск воркеров
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, data, wg)
	}

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	// Отдельная горутинка для отправки данных
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				data <- r.Intn(20001) - 10000
			}
		}
	}()

	// Ожидание сигнала завершения
	<-sigChan
	fmt.Println("\nGracefully shutdown...")
	cancel()    // Оповещаем всех воркеров о завершении
	wg.Wait()   // Ждем завершения всех воркеров
	close(data) // Закрываем канал данных
	fmt.Println("All workers have been shut down.")
}

// worker - функция для запуска воркера, который обрабатывает данные из канала
func worker(ctx context.Context, workerID int, data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d is shutting down...\n", workerID)
			return
		case input := <-data:
			fmt.Printf("Worker %d received: %d\n", workerID, input)
		}
	}
}
