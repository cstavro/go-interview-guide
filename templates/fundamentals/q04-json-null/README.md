# Problem: JSON Missing vs Null

Given the JSON payloads below, design a struct and a helper that can distinguish between missing fields, null fields, and present fields.

- Case 1: `{ "name": "Alice", "age": 30 }` — age is present
- Case 2: `{ "name": "Bob", "age": null }` — age is explicitly null
- Case 3: `{ "name": "Charlie" }` — age is missing

This is a common pattern for PATCH APIs.
