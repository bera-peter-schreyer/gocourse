package main

import (
	"strconv"
	"time"
)

// func main() {
// 	done := make(chan struct{})

// 	go func() {
// 		fmt.Println("Working...")
// 		time.Sleep(2 * time.Second)
// 		fmt.Println("Done!")
// 		done <- struct{}{}
// 	}()

// 	<- done
// 	fmt.Println("Main function finished")
// }

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		// Blocking until value is sent
// 		ch <- 9
// 		time.Sleep(1 * time.Second)
// 		fmt.Println("Sent value")
// 	}()

// 	// Blocking until value is received
// 	receiver := <-ch
// 	fmt.Println("Received value:", receiver)
// }

// SYCHRONIZATION EXAMPLE WITH SET NUMBER OF GOROUTINES
// func main() {
// 	numGoRoutines := 3
// 	done := make(chan int, numGoRoutines)

// 	for i := 0; i < numGoRoutines; i++ {
// 		go func(id int) {
// 			fmt.Printf("Working on %d...\n", id)
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Done:", id)
// 			done <- id
// 		}(i)
// 	}

// 	for i := 0; i < numGoRoutines; i++ {
// 		id := <-done
// 		fmt.Printf("Received done signal from %d\n", id)
// 	}
// 	fmt.Println("All goroutines finished")
// }

// SYCHRONIZING DATA EXCHANGE BETWEEN GOROUTINES
func channel_synchronization() {
	
	data := make(chan string)

	go func() {	
		for i := 0; i < 5; i++ {
			data <- "Hello " + strconv.Itoa(i)
			time.Sleep(100 * time.Millisecond)
		}
		// Close the channel after sending all data to avoid deadlock
		close(data)
	}()

	for value := range data {
		// This implicitly creates a receiver
		println("received value: ", value)
	}
	println("Main function finished")

}