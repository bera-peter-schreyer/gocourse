package main

import "time"

func non_blocking_channel_operation() {

	// ========================== NON BLOCKING RECEIVE OPERATION ==========================
	// ch := make(chan int)

	// select {
	// 	case msg := <-ch:
	// 		println("Received:", msg)
	// 	default:
	// 		println("No data received from channel")
	// }

	// ========================== NON BLOCKING SEND OPERATION ==========================
	// select {
	// 	case ch <- 1:
	// 		println("Sent: 1")
	// 	default:
	// 		println("No data sent to channel")
	// }

	// ========================== NON BLOCKING OPERATION IN REAL TIME SYSTEM ==========================
	data := make(chan int)
	done := make(chan bool)
	go func() {
		for {
			select {
				case msg := <-data:
					println("Received:", msg)
				case <-done:
					println("Done receiving data")
					return
				default:
					println("Waiting for data...")
					time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for i := range 5 {
		data <- i
		println("Sent:", i)
		time.Sleep(time.Second)
	}

	done <- true
	close(data)
	close(done)
}