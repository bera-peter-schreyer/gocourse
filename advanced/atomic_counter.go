package main

import (
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	value int64
}

func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func (c *AtomicCounter) GetValue() int64 {
	return atomic.LoadInt64(&c.value)
}

func atomicCounter() {
	// Access one value from multiple goroutines without using mutex
	var wg sync.WaitGroup
	counter := &AtomicCounter{}
	numGoRoutines := 10
	for range numGoRoutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.Increment()
			}
		}()
	}
	wg.Wait()
	println("Final counter value:", counter.GetValue())
}