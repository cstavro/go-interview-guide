# Problem: Safe Slice Append

Implement a function `AppendUnique(dest, src []int) []int` that appends elements from `src` to `dest` without causing the caller to accidentally share the backing array with other slices.

Include a demonstration that proves the original `dest` is unaffected if the returned slice grows beyond `dest`'s capacity.

## Expected Behavior

```
original := []int{1, 2, 3}
result := AppendUnique(original, []int{4, 5, 6, 7, 8})
// result should be [1 2 3 4 5 6 7 8]
// original should remain [1 2 3] regardless of capacity
```
