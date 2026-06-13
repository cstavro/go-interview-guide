# Problem: Debug Broken Cursor Pagination

This REST API handler uses cursor-based pagination. The cursor encodes a timestamp and an ID so the next page can resume where the previous page left off.

The bug is subtle: the cursor uses `json.Marshal` for an `int64` timestamp. JSON numbers are parsed as `float64` by default, which loses precision for large `int64` values (anything above 2^53). This causes the decoded cursor to have a different timestamp, and pagination returns the wrong results or loops forever.

Your task is to:
1. Write a test that demonstrates the precision loss.
2. Fix the cursor encoding so that round-trip preserves the exact timestamp.

## Using AI

- Ask AI to explain why JSON number precision matters for large `int64` values.
- Ask AI to explain the difference between `json.Marshal` of a struct vs. a custom binary encoding for cursors.
- Do not ask AI to "fix the cursor" — understand the root cause first, then implement the fix.

## Expected Behavior

Cursor round-trip should preserve the exact `int64` timestamp and string ID.
