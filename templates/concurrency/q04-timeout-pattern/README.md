# Problem: Timeout Pattern with Multiple Services

You need to call three services (A, B, C) concurrently. The overall operation must complete within 500ms. If any service fails or times out, return an error immediately. Do not leak goroutines.
