package main

import (
	"sync"
	"testing"
	"time"
)

func TestCacheRace(t *testing.T) {
	c := NewCache()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			c.Set("key", "value", time.Hour)
		}()
		go func() {
			defer wg.Done()
			c.Get("key")
		}()
	}
	wg.Wait()
}

func TestCacheStats(t *testing.T) {
	c := NewCache()
	c.Set("a", "1", time.Hour)
	c.Set("b", "2", time.Hour)

	_, _ = c.Get("a") // hit
	_, _ = c.Get("a") // hit
	_, _ = c.Get("c") // miss

	// TODO: implement Stats() and uncomment
	// stats := c.Stats()
	// if stats.Hits != 2 || stats.Misses != 1 {
	// 	t.Fatalf("expected hits=2, misses=1, got %+v", stats)
	// }
	t.Skip("Stats() not yet implemented")
}
