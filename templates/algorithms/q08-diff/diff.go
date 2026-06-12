package diff

// Result holds the diff output.
type Result struct {
	Added   []string // only in B
	Removed []string // only in A
	Common  []string // in both
}

// Diff computes the difference between two sorted slices.
func Diff(a, b []string) Result {
	// TODO
}
