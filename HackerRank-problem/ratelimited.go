package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	requests map[string]time.Time
	mu       sync.Mutex
	limit    time.Duration
}

func (r *RateLimiter) Allow(ip string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if t, exists := r.requests[ip]; exists && time.Since(t) < r.limit {
		return false
	}
	r.requests[ip] = time.Now()
	return true
}

func main() {
	rl := &RateLimiter{requests: make(map[string]time.Time), limit: time.Second}
	fmt.Println(rl.Allow("127.0.0.1")) // true
	fmt.Println(rl.Allow("127.0.0.1")) // false (within 1 sec)
}
