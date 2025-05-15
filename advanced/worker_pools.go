package main

import (
	"fmt"
	"time"
)

type ticketRequest struct {
	personID int
	numTickets int
	cost int
}

// Simulate processing a ticket request
func ticketProcessor(requests <-chan ticketRequest, results chan<- int) {
	for req := range requests {
		fmt.Println("Processing request for person", req.personID, "for", req.numTickets, "tickets with total cost", req.cost)
		// Simulate some processing time
		time.Sleep(time.Second)
		results <- req.personID
	}
}

func workerPools() {
	numRequests := 5
	price := 5
	numWorkers := 3

	ticketRequests := make(chan ticketRequest, numRequests)
	ticketResults := make(chan int)

	for range numWorkers {
		go ticketProcessor(ticketRequests, ticketResults)
	}

	for i := range numRequests {
		ticketRequests <- ticketRequest{
			personID: i + 1,
			numTickets: (i + 1) * 2,
			cost: price * (i + 1),
		}
	}
	close(ticketRequests)

	for range numRequests {
		result := <-ticketResults
		fmt.Println("Processed request for person", result)
	}
	close(ticketResults)
}


// BASIC WORKER POOL EXAMPLE
// func worker(id int, tasks <-chan int, results chan<- int) {
// 	for t := range tasks {
// 		println("Worker", id, "processing job", t)
// 		// Simulate some work
// 		time.Sleep(time.Second)
// 		results <- t * 2
// 	}
// }

// func main() {
// 	numWorkers := 3
// 	numTasks := 10

// 	tasks := make(chan int, numTasks)
// 	results := make(chan int, numTasks)

// 	// Create worker pool
// 	for w := range numWorkers {
// 		go worker(w, tasks, results)
// 	}

// 	// Send tasks to workers
// 	for i := range numTasks {
// 		tasks <- i
// 	}

// 	close(tasks)

// 	for range numTasks {
// 		result := <- results
// 		println("Result:", result)
// 	}
// }