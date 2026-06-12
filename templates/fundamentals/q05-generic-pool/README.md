# Problem: Generic Type-Safe Pool

Go's `sync.Pool` is not type-safe and can suffer from type assertions. Build a generic, type-safe object pool `Pool[T]` with `Get()` and `Put(T)` methods. It should use `sync.Pool` under the hood but expose a clean API. Include a benchmark comparing it to a naive allocation.
