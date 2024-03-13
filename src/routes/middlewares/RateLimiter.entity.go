package Middlewares

import (
	"sync"
	"time"
)

type RateLimiter struct {
	Mu        sync.Mutex
	Counters  map[string]int
	ResetTime time.Time
	MaxLimit  int
	Interval  time.Duration
}

func NewRateLimiter(MaxLimit int, Interval time.Duration) *RateLimiter {
	return &RateLimiter{
		Counters:  make(map[string]int),
		ResetTime: time.Now().Add(Interval),
		MaxLimit:  MaxLimit,
		Interval:  Interval,
	}
}

func (rl *RateLimiter) IncrementCounter(key string) int {
	rl.Mu.Lock()
	defer rl.Mu.Unlock()

	if time.Now().After(rl.ResetTime) {
		rl.ResetTime = time.Now().Add(rl.Interval)
		rl.Counters = make(map[string]int)
	}

	rl.Counters[key]++
	return rl.Counters[key]
}
