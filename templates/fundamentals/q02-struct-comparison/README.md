# Problem: Struct Comparison

Two `Config` structs are considered equal when all their fields match, including the slice of endpoints.

In Go, the `==` operator cannot be used on structs that contain slices because slices are reference types (they contain a pointer to a backing array). Even comparing slice headers with `==` only checks whether they point to the same backing array, not whether the contents are the same.

Implement the `Equal` method on `Config` so that two configs with identical values are considered equal, even if their backing arrays are different.

## Note

`==` on the struct itself will fail to compile:

```go
c1 := Config{Endpoints: []string{"a"}}
c2 := Config{Endpoints: []string{"a"}}
// if c1 == c2 { // compile error: struct containing []string cannot be compared
```
