package main

import (
	"fmt"
	"time"
)

// func main() {
// 	// 	BLOCKS ON RECEIVE IF THE BUFFER IS EMPTY
// 	ch := make(chan int, 2)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch <- 1
// 	}()
// 	fmt.Println("Value:", <-ch)
// 	fmt.Println("End of program")
// }

func channelBuffering() {
	// BLOCKS ON SEND IF THE BUFFER IS FULL
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	go func() {
		fmt.Println("Blocking")
		time.Sleep(2 * time.Second)
		fmt.Println("Received: ", <-ch)
	}()
	fmt.Println(("Start blocking"))
	ch <- 3
	fmt.Println(("End blocking"))
	fmt.Println("Received next: ", <-ch)
	fmt.Println("Received final: ", <-ch)
}