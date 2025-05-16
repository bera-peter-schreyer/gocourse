package main

import (
	"sync"
	"time"
)

type rateLimiter struct {
	mu sync.Mutex
	count int
	limit int
	window time.Duration
	resetTime time.Time
}

func newRateLimiter(limit int, window time.Duration) *rateLimiter {
	return &rateLimiter{
		limit:    limit,
		window:  window,
	}
}

func (rl *rateLimiter) allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	
	now := time.Now()
	if now.After(rl.resetTime) {
		rl.count = 0
		rl.resetTime = now.Add(rl.window)
	}

	if rl.count < rl.limit {
		rl.count++
		return true
	}
	return false
}

func rateLimiterFixedWindowAlgo() {
	rateLimiter := newRateLimiter(3, time.Second)

	for range 10 {
		if rateLimiter.allow() {
			// Perform the action
			println("Action performed")
		} else {
			println("Rate limit exceeded, action skipped")
		}
		time.Sleep(200 * time.Millisecond)
	}
}