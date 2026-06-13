package main

import (
	"fmt"
	"sync"
	"time"
)

// Batcher accumulates items and flushes them periodically.
type Batcher struct {
	mu       sync.Mutex
	items    []string
	interval time.Duration
	flushFn  func([]string)
	stop     chan struct{}
}

// NewBatcher creates a new Batcher.
func NewBatcher(interval time.Duration, flushFn func([]string)) *Batcher {
	b := &Batcher{
		interval: interval,
		flushFn:  flushFn,
		stop:     make(chan struct{}),
	}
	go b.loop()
	return b
}

// Add adds an item to the batch.
func (b *Batcher) Add(item string) {
	b.mu.Lock()
	b.items = append(b.items, item)
	b.mu.Unlock()
}

func (b *Batcher) loop() {
	ticker := time.NewTicker(b.interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			b.flush()
		case <-b.stop:
			return // BUG: remaining items are lost
		}
	}
}

func (b *Batcher) flush() {
	b.mu.Lock()
	if len(b.items) == 0 {
		b.mu.Unlock()
		return
	}
	items := b.items
	b.items = nil
	b.mu.Unlock()
	b.flushFn(items)
}

// Stop shuts down the batcher.
func (b *Batcher) Stop() {
	close(b.stop)
}

func main() {
	var flushed []string
	b := NewBatcher(100*time.Millisecond, func(items []string) {
		flushed = append(flushed, items...)
	})
	b.Add("a")
	b.Add("b")
	b.Stop()
	fmt.Printf("flushed: %v\n", flushed)
	// BUG: "a" and "b" may not be flushed
}
