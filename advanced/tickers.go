package main

import (
	"fmt"
	"time"
)

func tickers() {
	ticker := time.NewTicker(1 * time.Second)
	stop := time.After(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case tick := <-ticker.C:
			fmt.Println("Tick", tick)
		case <-stop:
			fmt.Println("Stopping ticker")
			return
		}
	}
}

// SCHEDULING PERIODIC TASKS
// func periodicTask() {
// 	fmt.Println(time.Now())
// }

// func main() {
// 	ticker := time.NewTicker(1 * time.Second)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ticker.C:
// 			periodicTask()
// 		}
// 	}
// }

// SIMPLE EXAMPLE
// func main() {
// 	ticker := time.NewTicker(1 * time.Second)
// 	defer ticker.Stop()
// 	// for tick := range ticker.C {
// 	// 	// Infinite loop of ticks
// 	// 	println("Tick", tick)
// 	// }
// }