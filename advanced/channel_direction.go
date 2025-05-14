package main

func channel_direction() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}

// Receive-only channel
func consumer(ch <-chan int) {
	for v := range ch {
		println("Received:", v)
	}
}	

// Send-only channel
func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}