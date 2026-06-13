package main

import (
	"fmt"
	"time"
)

type Config struct {
	Name      string
	Timeout   time.Duration
	Endpoints []string
}

// Equal reports whether c and other represent the same configuration.
// Two configs are equal if all fields match, including the slice contents.
func (c Config) Equal(other Config) bool {
	// TODO: implement
	return false
}

func main() {
	c1 := Config{Name: "svc", Timeout: 5 * time.Second, Endpoints: []string{"a", "b"}}
	c2 := Config{Name: "svc", Timeout: 5 * time.Second, Endpoints: []string{"a", "b"}}

	if c1.Equal(c2) {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}
