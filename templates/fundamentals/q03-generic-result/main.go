package main

import "fmt"

// Result holds either a value of type T or an error.
type Result[T any] struct {
	// TODO: define fields
}

// Map transforms the value if present.
func (r Result[T]) Map(fn func(T) T) Result[T] {
	// TODO
	return Result[T]{}
}

// OrElse returns the value if present, otherwise computes a fallback value.
func (r Result[T]) OrElse(fn func() T) Result[T] {
	// TODO
	return Result[T]{}
}

// parseInt attempts to parse a string as an int.
func parseInt(s string) Result[int] {
	// TODO
	return Result[int]{}
}

func main() {
	// Example chain:
	// result := parseInt("42").
	// 	Map(func(n int) int { return n * 2 }).
	// 	Map(func(n int) int { return n + 10 }).
	//  OrElse(func() int { return 0 })
	fmt.Println("Hello, Result!")
}
