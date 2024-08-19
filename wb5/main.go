package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run [path] [N seconds]")
		return
	}

	seconds, err := strconv.Atoi(os.Args[1])
	if err != nil || seconds <= 0 {
		fmt.Println("Invalid time")
		return
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))

	data := make(chan int)

	wg := new(sync.WaitGroup)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(seconds)*time.Second)
	defer cancel()
	wg.Add(2)
	go func(wg *sync.WaitGroup, data <-chan int, ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Timeout, reading goroutine is shutdown")
				return
			case input := <-data:
				fmt.Println("Received: ", input)
			}
		}
	}(wg, data, ctx)

	go func(wg *sync.WaitGroup, data chan<- int, ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Timeout, writing goroutine is exiting")
				return
			case data <- r.Intn(20001) - 10000:
				time.Sleep(time.Second)
			}
		}
	}(wg, data, ctx)

	wg.Wait()

	close(data)
}
