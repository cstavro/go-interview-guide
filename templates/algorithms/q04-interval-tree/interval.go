package interval

import (
	"time"
)

// Interval represents a time range.
type Interval struct {
	ID    string
	Start time.Time
	End   time.Time
}

// Tree is an interval tree.
type Tree struct {
	// TODO
}

// NewTree creates an empty tree.
func NewTree() *Tree {
	// TODO
}

// Insert adds an interval to the tree.
func (t *Tree) Insert(iv Interval) {
	// TODO
}

// Delete removes an interval by ID.
func (t *Tree) Delete(id string) {
	// TODO
}

// Overlapping returns all intervals that overlap with q.
func (t *Tree) Overlapping(q Interval) []Interval {
	// TODO
}
