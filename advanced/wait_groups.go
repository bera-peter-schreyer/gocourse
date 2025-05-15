package main

import (
	"sync"
	"time"
)

// ===============CONSTRUCTION EXAMPLE WITHOUT CHANNELS
type worker struct {
	ID int
	Task string
}


func (w *worker) performTask(wg *sync.WaitGroup) {
	defer wg.Done()
	println("Worker", w.ID, "performing task:", w.Task)
	// Simulate some work
	time.Sleep(time.Second)
	println("Worker", w.ID, "finished task:", w.Task)
}

func waitGroups() {
	var wg sync.WaitGroup
	tasks := []string{"digging", "painting", "bricklaying"}

	for i, task := range tasks {
		wg.Add(1)
		w := worker{ID: i + 1, Task: task}
		go w.performTask(&wg)
	}
	wg.Wait() // Wait for all workers to finish
	println("All workers finished")
}

// =============EXAMPLE WITH CHANNELS
// func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Worker", id, "started")
// 	for task := range tasks {
// 		time.Sleep(1 * time.Second)
// 		results <- task * 2
// 	}
// 	fmt.Println("Worker", id, "finished")
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3
// 	numJobs := 6
// 	tasks := make(chan int, numJobs)
// 	results := make(chan int, numJobs)

// 	for i := range numJobs {
// 		tasks <- i + 1
// 	}
// 	close(tasks)

// 	wg.Add(numWorkers)

// 	for i := range numWorkers {
// 		go worker(i+1, tasks, results, &wg)
// 	}

// 	go func() {
// 		// Move this into goroutine so it prints the result updates below in real time and is not blocked by waiting for all workers to finish
// 		wg.Wait()
// 		close(results)
// 	}()

// 	for result := range results {
// 		fmt.Println("Result:", result)
// 	}
// 	fmt.Println("All workers finished")
// }

// =================BASIC EXAMPLE WITHOUT CHANNELS
// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Worker", id, "started")
// 	time.Sleep(time.Second)
// 	fmt.Println("Worker", id, "finished")
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3

// 	wg.Add(numWorkers)

// 	for i := range numWorkers {
// 		go worker(i+1, &wg)
// 	}

// 	wg.Wait() // blocking mechanism
// 	fmt.Println("All workers finished")
// }