# Problem: Add Validation and Safe Logging to Config

This service loads configuration but has no validation. Invalid configs cause runtime panics or obscure errors. The `String()` method is also missing, making it easy to accidentally log secrets.

Your task is to:
1. Implement `Validate() error` that checks for sensible defaults and returns clear, actionable errors.
2. Implement `String() string` that returns a human-readable representation with the `DBPass` field redacted.

## Using AI

- Ask AI to suggest a validation pattern for Go configs that fails fast with helpful errors.
- Ask AI to review your `String()` implementation to ensure it doesn't accidentally expose the password.
- Do not ask AI to "implement validation" — think about the edge cases first, then use AI to review your approach.

## Expected Behavior

`Validate()` should catch zero port, empty host, and non-positive workers. `String()` should never contain the raw password.
