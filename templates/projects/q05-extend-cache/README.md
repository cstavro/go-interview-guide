# Problem: Fix a Race and Add Cache Stats

This simple in-memory cache supports TTL but is not safe for concurrent use. Running `go test -race` will fail.

Your task is to:
1. Fix the race condition so that the cache is safe for concurrent access.
2. Add a `Stats()` method that returns the number of hits and misses since the cache was created.

## Using AI

- Ask AI to explain the difference between `sync.Mutex` and `sync.RWMutex` and when to use each.
- Ask AI to review your `Stats` implementation to ensure it doesn't introduce a new race.
- Do not ask AI to "fix the race" — use `go test -race` to understand the issue, then apply the fix.

## Expected Behavior

`go test -race` should pass. The `Stats()` method should return accurate hit and miss counts.
