package main

import (
	"time"
)

// SlidingWindowRateLimiter limits requests per user.
type SlidingWindowRateLimiter struct {
	// TODO
}

// NewSlidingWindowRateLimiter creates a limiter.
// window: time window (e.g., 1 minute)
// maxRequests: max requests allowed in the window.
func NewSlidingWindowRateLimiter(window time.Duration, maxRequests int) *SlidingWindowRateLimiter {
	// TODO
}

// Allow checks if a request from userID at now is allowed.
func (rl *SlidingWindowRateLimiter) Allow(userID string, now time.Time) bool {
	// TODO
}
