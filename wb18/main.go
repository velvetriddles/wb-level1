package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Еще можно сделать с мьютексами

type Counter struct {
	count int64
}

func NewCounter() *Counter {
	return &Counter{count: 0}
}

func (c *Counter) Add() {
	atomic.AddInt64(&c.count, 1)
}

func (c *Counter) GetValue() int64 {
	return atomic.LoadInt64(&c.count)
}

func main() {
	wg := new(sync.WaitGroup)
	counter := NewCounter()

	// Пошла вода горячая
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Add()
		}()
	}

	wg.Wait()

	fmt.Println("Value:", counter.GetValue())
}
