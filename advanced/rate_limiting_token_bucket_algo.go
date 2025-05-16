package main

import "time"

type RateLimiter struct {
	tokens chan struct{}
	refillTime time.Duration
}

func NewRateLimiter(rateLimit int, refillTime time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens:    make(chan struct{}, rateLimit),
		refillTime: refillTime,
	}
	for range rateLimit {
		rl.tokens <- struct{}{}
	}
	go rl.startRefill()
	return rl
}

func (rl *RateLimiter) startRefill() {
	ticker := time.NewTicker(rl.refillTime)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			select {
			// Try to insert a token into the channel
			case rl.tokens <- struct{}{}:
			default:
			}
		}
	}
}

func (rl *RateLimiter) allow() bool {
	select {
		// Check if channel has a token available
	case <-rl.tokens:
		return true
	default:
		return false
	}
}

func rateLimitingTokenBucketAlgo() {
	rateLimiter := NewRateLimiter(5, time.Second)

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