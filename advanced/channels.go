package main

import (
	"fmt"
	"math/rand"
	"time"
)

func channels() {
	// Create a channel
	ch := make(chan int)

	// Start a goroutine to send data to the channel
	go func() {
		for i := 0; i < 5; i++ {
			ch <- rand.Intn(100) // Send random number to channel
			time.Sleep(time.Second)
		}
		close(ch) // Close the channel after sending all data
	}()

	// Receive data from the channel
	for value := range ch {
		fmt.Println(value) // Print the received number
	}

}
	