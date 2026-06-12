package lru

// Cache is an LRU cache with a maximum size.
type Cache struct {
	// TODO
}

// NewCache creates a cache with the given capacity.
func NewCache(capacity int) *Cache {
	// TODO
}

// Get retrieves a value. Returns nil, false if not found.
func (c *Cache) Get(key string) (any, bool) {
	// TODO
}

// Set adds or updates a key.
func (c *Cache) Set(key string, value any) {
	// TODO
}

// Delete removes a key.
func (c *Cache) Delete(key string) {
	// TODO
}
