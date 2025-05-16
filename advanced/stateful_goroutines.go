package main

import "time"

type statefulWorker struct {
	count int
	ch chan int
}

func (w *statefulWorker) start() {
	for {
		select {
		case value := <-w.ch:
			w.count += value
			println("Count:", w.count)
		}
	}
}

func (w *statefulWorker) send(value int) {
	w.ch <- value
}

func statefulGoroutines() {
	ch := make(chan int)
	worker := statefulWorker{
		count: 0,
		ch:     ch,
	}
	go worker.start()
	for i := range 5 {
		worker.send(i)
		time.Sleep(100 * time.Millisecond)
	}
}
