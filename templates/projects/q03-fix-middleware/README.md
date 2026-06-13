# Problem: Fix Middleware Order and Add Timing

This HTTP server has middleware, but the panic recovery middleware is applied in the wrong order. When a handler panics, the recovery middleware catches it but the logging middleware never sees the request — so panics are invisible in logs.

Your task is to fix the middleware order so that panics are both logged and recovered. Then add a new `TimingMiddleware` that records the request duration in a header (`X-Request-Duration-Ms`).

## Using AI

- Ask AI to explain how middleware wrapping order works in Go (which middleware runs first).
- Ask AI to review your `TimingMiddleware` implementation for correctness.
- Do not ask AI to "fix the middleware order" — reason about it first, then verify.

## Expected Behavior

A panic in a handler should result in a 500 response, and the request should still appear in the log with its duration.
