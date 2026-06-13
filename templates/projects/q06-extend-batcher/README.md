# Problem: Add Graceful Shutdown to a Batcher

This batch processor accumulates items and flushes them periodically. It has a bug: when `Stop()` is called, any remaining items are lost because the goroutine exits immediately without flushing.

Your task is to:
1. Fix the bug so that `Stop()` flushes any remaining items before returning.
2. Add a `Flush()` method that forces an immediate flush and can be called safely from any goroutine.

## Using AI

- Ask AI to explain patterns for graceful shutdown in Go with `context.Context`.
- Ask AI to explain how to safely signal a goroutine to flush and exit without losing data.
- Do not ask AI to "implement the shutdown" — reason about the channel and mutex interactions first.

## Expected Behavior

After calling `Stop()`, all items added before the call should be flushed. The `Flush()` method should not race with the background loop.
