package main

import (
	"sync"
	"time"
)

type LeakyBucket struct {
	capacity int
	leakRate time.Duration
	tokenCount int
	lastLeakTime time.Time
	mu sync.Mutex
}

func newLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		tokenCount: capacity,
		lastLeakTime: time.Now(),
	}
}

func (b *LeakyBucket) allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(b.lastLeakTime)
	tokensToAdd := int(elapsed / b.leakRate)
	b.tokenCount += tokensToAdd

	if b.tokenCount > b.capacity {
		b.tokenCount = b.capacity
	}
	
	b.lastLeakTime = b.lastLeakTime.Add(time.Duration(tokensToAdd) * b.leakRate)

	if b.tokenCount > 0 {
		b.tokenCount--
		return true
	}
	return false
}

func rateLimitingLeakyBucketAlgo() {
	leakyBucket := newLeakyBucket(5, 500*time.Millisecond)
	for range 10 {
		if leakyBucket.allow() {
			// Perform the action
			println("Action performed")
		} else {
			println("Rate limit exceeded, action skipped")
		}
		time.Sleep(100 * time.Millisecond)
	}
}