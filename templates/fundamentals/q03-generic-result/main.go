package main

import "fmt"

// Result holds either a value of type T or an error.
type Result[T any] struct {
	// TODO: define
}

// Map transforms the value if present.
func (r Result[T]) Map(fn func(T) T) Result[T] {
	// TODO
}

// MapErr transforms the error if present.
func (r Result[T]) MapErr(fn func(error) error) Result[T] {
	// TODO
}

// OrElse returns the value if present, otherwise computes a fallback.
func (r Result[T]) OrElse(fn func() Result[T]) Result[T] {
	// TODO
}

func main() {
	// Chain: parse int, double it, stringify, fallback on error.
}
