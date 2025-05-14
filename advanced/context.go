package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func doWork(ctx context.Context, work string) {
	for {
		select {
			case <-ctx.Done():
				fmt.Println("Work cancelled:", work)
				return
			default:
				fmt.Println("Doing work:", work)
			}
		time.Sleep(500 * time.Millisecond)
	}
}

func logWithContext(ctx context.Context, message string) {
	requestID := ctx.Value("requestID")
	if requestID != nil {
		log.Printf("Request ID: %s, Message: %s\n", requestID, message)
	} else {
		fmt.Println("No Request ID found, Message:", message)
	}
}

func contextExample() {
	rootCtx := context.Background()
	// Can either let the context be cancelled after a timeout or manually cancel it
	// ctx, cancel := context.WithTimeout(rootCtx, 2*time.Second)
	// defer cancel()

	ctx, cancel := context.WithCancel(rootCtx)
	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	ctx = context.WithValue(ctx, "requestID", "123-456")
	go doWork(ctx, "Task 1")
	time.Sleep(3 * time.Second)

	// Even after the context is cancelled, the values are still accessible
	requestID := ctx.Value("requestID")
	if requestID == nil {
		fmt.Println("Request ID not found")
	} else {
		fmt.Println("Request ID:", requestID)
	}

	logWithContext(ctx, "This is a log message")
}


// EXAMPLE: CONTEXT WITH TIMEOUT
// func checkEvenOdd(ctx context.Context, number int) string {
// 	select {
// 		case <-ctx.Done():
// 			return "Operation cancelled"
// 		default:
// 			if number%2 == 0 {
// 				return "Even"
// 			}
// 			return "Odd"
// 	}
// }

// func main() {
// 	ctx := context.TODO()
// 	fmt.Println(checkEvenOdd(ctx, 4))
// 	fmt.Println(checkEvenOdd(ctx, 5))

// 	ctx = context.Background()
// 	ctx, cancel := context.WithTimeout(ctx, 1 * time.Second)
// 	defer cancel()

// 	result := checkEvenOdd(ctx, 6)
// 	fmt.Println("Result from timeout: ", result)
// 	time.Sleep(2 * time.Second)
// 	result = checkEvenOdd(ctx, 6)
// 	fmt.Println("Result after timeout: ", result)
// }

// CONTEXT.TODO VS CONTEXT.BACKGROUND
// func main() {
// 	todoContext := context.TODO()
// 	bkgContext := context.Background()

// 	ctx := context.WithValue(todoContext, "name", "John Doe")
// 	fmt.Println(ctx)
// 	fmt.Println(ctx.Value("name"))

// 	ctx1 := context.WithValue(bkgContext, "city", "New York")
// 	fmt.Println(ctx1)
// 	fmt.Println(ctx1.Value("city"))
// }