package pool

import "sync"

// Pool is a type-safe wrapper around sync.Pool.
type Pool[T any] struct {
	// TODO
}

// New creates a Pool. The newFn should return a zero-value T.
func New[T any](newFn func() T) *Pool[T] {
	// TODO
}

// Get returns an item from the pool or a new one.
func (p *Pool[T]) Get() T {
	// TODO
}

// Put returns an item to the pool.
func (p *Pool[T]) Put(v T) {
	// TODO
}
