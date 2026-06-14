package ratelimit

import "time"

// Clock abstracts time so tests can advance time manually.
type Clock interface {
	Now() time.Time
}

type realClock struct{}

func (realClock) Now() time.Time { return time.Now() }

// TokenBucket implements a token bucket rate limiter.
type TokenBucket struct {
	// TODO
}

// NewTokenBucket creates a limiter with the given capacity and refill rate.
func NewTokenBucket(capacity int, refillRate time.Duration) *TokenBucket {
	return NewTokenBucketWithClock(capacity, refillRate, realClock{})
}

// NewTokenBucketWithClock creates a limiter using the provided clock.
func NewTokenBucketWithClock(capacity int, refillRate time.Duration, clock Clock) *TokenBucket {
	// TODO
	return nil
}

// Allow consumes a single token if available.
func (tb *TokenBucket) Allow() bool {
	return tb.AllowN(1)
}

// AllowN attempts to consume n tokens.
func (tb *TokenBucket) AllowN(n int) bool {
	// TODO
	return false
}
