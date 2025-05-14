package main

import "fmt"


func producer_channel(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}
func filter_channel(in <-chan int, out chan<- int) {
	for v := range in {
		if v%2 == 0 {
			fmt.Println("Even number:", v)
			out <- v
		}
	}
	close(out)
}

func closing_channels() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go producer_channel(ch1)
	go filter_channel(ch1, ch2)

	for v := range ch2 {
		fmt.Println("Filtered value:", v)
	}
}


// CLoSING CHANNELS TWICE
// func main() {
// 	// Create a channel
// 	ch := make(chan int)

// 	// If we accidentally close a channel twice, it will panic
// 	go func() {
// 		close(ch)
// 		close(ch)
// 	}()

// 	time.Sleep(time.Second)
// }

// Range over a closed channel
// func main() {
// 	ch := make(chan int)
	
// 	go func() {
// 		for i := range 5 {
// 			ch <- i
// 			println("Sent:", i)
// 		}
// 		close(ch) // Close the channel after sending
// 	}()

// 	for val := range ch {
// 		println("Received:", val)
// 	}
// }

// Receiving from a closed channel
// func main() {
// 	ch := make(chan int)
// 	close(ch) // Close the channel before sending
	
// 	val, ok := <-ch
// 	if !ok {
// 		println("Channel is closed, no value received")
// 	} else {
// 		println("Received:", val)
// 	}
// }

// Simple closing channel example
// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		for i := range 5 {
// 			ch <- i
// 			println("Sent:", i)
// 		}
// 		close(ch)
// 	}()

// 	for val := range ch {
// 		println("Received:", val)
// 	}
// }