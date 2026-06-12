package cache

import "time"

type Cache struct {
	data map[string]string
}

func (c *Cache) Set(key, val string) {
	c.data[key] = val
}

func (c *Cache) Get(key string) string {
	return c.data[key]
}

func main() {
	c := &Cache{data: make(map[string]string)}
	go func() { for { c.Set("x", "y") } }()
	go func() { for { _ = c.Get("x") } }()
	time.Sleep(1 * time.Second)
}
