package main

import (
	"sync"
	"testing"
	"time"
)

// Counter is a thread-safe counter.
type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func TestCounterConcurrent(t *testing.T) {
	c := &Counter{}
	for i := 0; i < 100; i++ {
		go func() {
			c.Inc()
		}()
	}
	// BUG: flaky sleep — may be too short or too long.
	// The real fix is to use sync.WaitGroup.
	// On slower machines or with GOMAXPROCS=1, this may pass some of the time.
	// Run with -count=100 to see failures.
	time.Sleep(0)
	if got := c.Value(); got != 100 {
		t.Fatalf("expected 100, got %d", got)
	}
}
