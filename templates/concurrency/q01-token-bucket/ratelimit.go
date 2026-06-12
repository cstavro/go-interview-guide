package ratelimit

import "time"

// TokenBucket implements a token bucket rate limiter.
type TokenBucket struct {
	// TODO
}

// NewTokenBucket creates a limiter with the given capacity and refill rate.
func NewTokenBucket(capacity int, refillRate time.Duration) *TokenBucket {
	// TODO
}

// Allow consumes a single token if available.
func (tb *TokenBucket) Allow() bool {
	return tb.AllowN(1)
}

// AllowN attempts to consume n tokens.
func (tb *TokenBucket) AllowN(n int) bool {
	// TODO
}
