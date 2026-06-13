package main

import (
	"time"
)

// Cache is a simple TTL cache with a race condition.
type Cache struct {
	data map[string]item
}

type item struct {
	value      string
	expiration time.Time
}

// NewCache creates a new cache.
func NewCache() *Cache {
	return &Cache{data: make(map[string]item)}
}

// Set stores a value with a TTL.
func (c *Cache) Set(key, val string, ttl time.Duration) {
	c.data[key] = item{value: val, expiration: time.Now().Add(ttl)}
}

// Get retrieves a value if it has not expired.
func (c *Cache) Get(key string) (string, bool) {
	it, ok := c.data[key]
	if !ok || time.Now().After(it.expiration) {
		return "", false
	}
	return it.value, true
}

func main() {
	c := NewCache()
	c.Set("hello", "world", time.Minute)
	val, ok := c.Get("hello")
	if ok {
		println(val)
	}
}
