package main

import "time"

func multiplexing_select() {
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
		close(ch)
	}()

	for {
		select {
			case msg, ok := <-ch:
				if !ok {
					println("Channel closed")
					// Do cleanup here
					return
				}
				println("Received:", msg)
			case <-time.After(1 * time.Second):
				println("Timeout: No message received within 1 second")
			// default:
				// 	println("No data received from channel")
		}
	}
}

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch <- 1
// 		close(ch)
// 	}()

// 	select {
// 		case msg, ok := <-ch:
// 			if !ok {
// 				println("Channel closed")
// 				return
// 			}
// 			println("Received:", msg)
// 		case <-time.After(3 * time.Second):
// 			println("Timeout: No message received within 3 seconds")
// 		case <-time.After(2 * time.Second):
// 			println("Timeout: No message received within 2 seconds")
// 		case <-time.After(1 * time.Second):
// 			println("Timeout: No message received within 1 second")
// 		// default:
// 		// 	println("No data received from channel")
// 	}
// }


// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)
  
// 	go func() {
// 		ch1 <- 42
// 	}()
		
// 	go func() {
// 		ch2 <- 84
// 	}()

// 	time.Sleep(time.Second)

// 	for range 2 {
// 		select {
// 		case msg := <-ch1:
// 			println("Received from ch1:", msg)
// 		case msg := <-ch2:
// 			println("Received from ch2:", msg)
// 		default:
// 			println("No data received from either channel")
// 		}
// 	}
// }