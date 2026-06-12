package lru

import "testing"

func TestLRUCache(t *testing.T) {
	cache := NewCache(2)
	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("c", 3) // should evict "a"

	if _, ok := cache.Get("a"); ok {
		t.Error("expected 'a' to be evicted")
	}
	if _, ok := cache.Get("b"); !ok {
		t.Error("expected 'b' to be present")
	}
}
