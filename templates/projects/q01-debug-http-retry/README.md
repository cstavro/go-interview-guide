# Problem: Debug HTTP Retry Goroutine Leak

A small HTTP client with retry logic is leaking goroutines. The client has a `Do` method that retries failed requests up to `MaxRetries` times with an exponential backoff.

The tests pass, but `go test` shows a growing goroutine count. Run the benchmark and use `go test -bench` to confirm the leak. Then fix the root cause.

## Using AI

- Ask AI to explain what happens when a goroutine sends on a channel that nobody is receiving.
- Ask AI to suggest patterns for preventing goroutine leaks in retry loops.
- Do not ask AI to "fix the leak for you" — verify the fix yourself.

## Expected Behavior

After the fix, the benchmark should show a stable goroutine count regardless of retries.
