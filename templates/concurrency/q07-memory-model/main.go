package main

import (
	"fmt"
	"runtime"
)

func main() {
	var a string
	var done bool

	go func() {
		a = "hello, world"
		done = true
	}()

	for !done {
		runtime.Gosched()
	}
	fmt.Println(a)
}
