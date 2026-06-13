package main

import (
	"sync"
	"testing"
	"time"
)

func TestBatcherFlushOnStop(t *testing.T) {
	var flushed []string
	var mu sync.Mutex
	b := NewBatcher(1*time.Hour, func(items []string) {
		mu.Lock()
		flushed = append(flushed, items...)
		mu.Unlock()
	})

	b.Add("a")
	b.Add("b")
	b.Stop()

	time.Sleep(100 * time.Millisecond)
	mu.Lock()
	defer mu.Unlock()

	if len(flushed) != 2 {
		t.Fatalf("expected 2 items flushed, got %d: %v", len(flushed), flushed)
	}
}

func TestBatcherFlushMethod(t *testing.T) {
	var flushed []string
	var mu sync.Mutex
	b := NewBatcher(1*time.Hour, func(items []string) {
		mu.Lock()
		flushed = append(flushed, items...)
		mu.Unlock()
	})

	b.Add("x")
	// TODO: implement and test Flush()
	// b.Flush()
	b.Stop()

	time.Sleep(100 * time.Millisecond)
	mu.Lock()
	defer mu.Unlock()

	// After implementing Flush, uncomment:
	// if len(flushed) != 1 || flushed[0] != "x" {
	// 	t.Fatalf("expected [x], got %v", flushed)
	// }
	t.Skip("Flush() not yet implemented")
}
