# Problem: Fix a Flaky Concurrent Test

The `Counter` implementation is correct, but the test is flaky. It uses `time.Sleep` to wait for goroutines to finish, which is unreliable on slow machines and wasteful on fast ones.

Your task is to fix the test without changing the `Counter` implementation. Make the test robust and fast. Then run it with `go test -count=100` to prove it is stable.

## Using AI

- Ask AI to explain why `time.Sleep` makes tests flaky.
- Ask AI to explain how `sync.WaitGroup` works and when it is appropriate.
- Do not ask AI to "fix the test" — write the fix yourself and ask AI to review it.

## Expected Behavior

`go test -count=100` should pass consistently without increasing test duration.
