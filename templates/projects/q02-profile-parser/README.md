# Problem: Profile and Optimize a Log Parser

A log parser processes a large stream of text lines. The current implementation works correctly but is very slow on large inputs. The benchmark shows excessive allocations.

Your task is to run the benchmark, use `go test -bench -memprofile` to understand where the allocations happen, and then optimize the hot path. Aim for at least a 2x speedup and 2x reduction in allocations.

## Using AI

- Ask AI to help you interpret the CPU and memory profile output.
- Ask AI to explain the difference between `strings.Split`, `strings.Fields`, and `bytes.Fields` for this use case.
- Do not ask AI to "rewrite the parser for me" — compare the suggestions and pick the best approach.

## Expected Behavior

The optimized parser should pass all existing tests while being significantly faster.
