package main

import (
	"fmt"
	"time"
)

// SCHEDULING DELAYED OPERATIONS
func timers() {
	timer := time.NewTimer(2 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("Timer expired")
	}()

	fmt.Println("Waiting...")
	time.Sleep(3 * time.Second) // blocking timer
	fmt.Println("Main function completed")
}


// EXAMPLE: TIMEOUT
// func longRunningTask() {
// 	for i:= range 20 {
// 		// Simulate a long-running task
// 		fmt.Println(i)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	timeout := time.After(2 * time.Second)
// 	done := make(chan bool)

// 	go func(){
// 		longRunningTask()
// 		done <- true
// 	}()

// 	select {
// 		case <-timeout:
// 			fmt.Println("Task timed out")
// 		case <-done:
// 			fmt.Println("Task completed")
// 	}
// }


// BASIC TIMER EXAMPLE
// func main() {
// 	timer := time.NewTimer(2 * time.Second)
// 	defer timer.Stop()
// 	stopped := timer.Stop()
// 	if stopped {
// 		println("Timer stopped")
// 	} else {
// 		println("Timer was not stopped")
// 	}
// 	timer.Reset(3 * time.Second)
// 	println("Timer reset")
// 	<- timer.C // blocking until the timer expires
// 	println("Timer expired")
// }