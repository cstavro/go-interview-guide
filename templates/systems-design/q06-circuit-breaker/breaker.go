package breaker

import (
	"errors"
	"time"
)

// State represents the circuit breaker state.
type State int

const (
	StateClosed State = iota
	StateOpen
	StateHalfOpen
)

// Breaker is a generic circuit breaker.
type Breaker struct {
	// TODO
}

// NewBreaker creates a breaker.
func NewBreaker(failureThreshold int, timeout time.Duration) *Breaker {
	// TODO
}

// Call executes the function if the breaker allows.
func (b *Breaker) Call(fn func() error) error {
	// TODO
}
