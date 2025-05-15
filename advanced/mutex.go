package main

import (
	"sync"
)



func mutex() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup
	numGoRoutines := 5

	wg.Add(numGoRoutines)
	increment := func() {
		defer wg.Done()
		for range 1000 {
			mu.Lock()
			// Code between lock and unlock is critical section that is protected by the mutex
			counter++
			mu.Unlock()
		}
	}

	for range numGoRoutines {
		go increment()
	}
	wg.Wait()
	println("Final counter value:", counter)
}


// ====== SIMPLE MUTEX EXAMPLE
// type counter struct {
// 	mu    sync.Mutex
// 	value int
// }

// func (c *counter) increment() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.value++
// }

// func (c *counter) getValue() int {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.value
// }

// func main() {
// 	var wg sync.WaitGroup
// 	counter := &counter{}
// 	numGoRoutines := 10
	
// 	for range numGoRoutines {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for range 1000 {
// 				counter.increment()
// 			}
// 		}()
// 	}
// 	wg.Wait()
// 	println("Counter value:", counter.getValue())
// }