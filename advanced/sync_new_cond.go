package main

import (
	"sync"
	"time"
)

const bufferSize = 5

type buffer struct {
	cond *sync.Cond
	items  []int
	mu   sync.Mutex
}

func newBuffer(size int) *buffer {
	b := &buffer{
		items: make([]int, 0, size),
	}
	b.cond = sync.NewCond(&b.mu)
	return b
}

func (b *buffer) produce(item int){
	b.mu.Lock()
	defer b.mu.Unlock()
	
	for len(b.items) == bufferSize {
		b.cond.Wait()
	}
	
	b.items = append(b.items, item)
	b.cond.Signal()
}

func (b *buffer) consume() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	for len(b.items) == 0 {
		b.cond.Wait()
	}
	
	item := b.items[0]
	b.items = b.items[1:]
	b.cond.Signal()
	return item
}

func producer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 20 {
		b.produce(i + 100)
		time.Sleep(time.Millisecond * 100)
	}
}

func consumer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 20 {
		b.consume()
		time.Sleep(time.Millisecond * 200)
	}
}

func syncNewCond() {
	b := newBuffer(bufferSize)
	var wg sync.WaitGroup
	
	wg.Add(2)
	go producer(b, &wg)
	go consumer(b, &wg)
	wg.Wait()
}