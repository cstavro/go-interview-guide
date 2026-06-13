# Problem: Generic Result Type

Go does not have a built-in `Result<T, E>` type like Rust. Implement a generic `Result[T]` type that holds either a value or an error. Provide `Map` and `OrElse` methods so the type supports fluent chaining.

## Requirements

- `Map` applies a transformation to the value when present.
- `OrElse` provides a fallback value of type `T` when the current result holds an error.
- The methods should be chainable.

## Example Chain

```go
result := parseInt("42").
    Map(func(n int) int { return n * 2 }).
    Map(func(n int) int { return n + 10 }).
    OrElse(func() int { return 0 })
```

## Follow-up

Why can't `Map` change the generic type (e.g., `Result[int]` to `Result[string]`) when it is defined as a method?

In Go, methods on a generic type cannot introduce additional type parameters beyond those declared on the receiver. `Map` is defined as:

```go
func (r Result[T]) Map(fn func(T) T) Result[T]
```

To transform `Result[int]` into `Result[string]`, `Map` would need to be a standalone function with two type parameters:

```go
func Map[T, U any](r Result[T], fn func(T) U) Result[U]
```

But this breaks the fluent method-chaining style because it is no longer a method on the receiver. How would you redesign the API to support type-changing transformations while still keeping chaining ergonomic?
