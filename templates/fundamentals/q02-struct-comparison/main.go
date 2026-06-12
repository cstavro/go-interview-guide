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

func main() {
	c1 := Config{Name: "svc", Timeout: 5 * time.Second, Endpoints: []string{"a", "b"}}
	c2 := Config{Name: "svc", Timeout: 5 * time.Second, Endpoints: []string{"a", "b"}}

	if c1 == c2 {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}
