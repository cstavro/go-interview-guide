package ratelimit

import (
	"sync"
	"testing"
	"time"
)

// manualClock is a fake clock that tests can advance explicitly.
type manualClock struct {
	mu  sync.Mutex
	now time.Time
}

func (m *manualClock) Now() time.Time {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.now
}

func (m *manualClock) Advance(d time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.now = m.now.Add(d)
}

func TestTokenBucket(t *testing.T) {
	clock := &manualClock{now: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}
	bucket := NewTokenBucketWithClock(5, time.Second, clock)

	if !bucket.AllowN(5) {
		t.Error("expected burst of 5 to be allowed")
	}
	if bucket.Allow() {
		t.Error("expected throttled after burst")
	}

	clock.Advance(time.Second)
	if !bucket.Allow() {
		t.Error("expected one token after refill interval")
	}
	if bucket.Allow() {
		t.Error("expected throttle after single refill")
	}
}

func TestTokenBucketConcurrent(t *testing.T) {
	clock := &manualClock{now: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}
	bucket := NewTokenBucketWithClock(100, time.Second, clock)

	var wg sync.WaitGroup
	results := make(chan bool, 200)

	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results <- bucket.Allow()
		}()
	}

	wg.Wait()
	close(results)

	allowed := 0
	for ok := range results {
		if ok {
			allowed++
		}
	}

	if allowed != 100 {
		t.Errorf("expected exactly 100 concurrent allows, got %d", allowed)
	}
}
