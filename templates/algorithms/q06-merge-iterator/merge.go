package merge

// Iterator yields sorted values.
type Iterator interface {
	Next() bool
	Value() int
}

// MergeIterator merges N sorted iterators.
type MergeIterator struct {
	// TODO
}

// NewMergeIterator creates a merge iterator.
func NewMergeIterator(iters []Iterator) *MergeIterator {
	// TODO
}

// Next advances to the next smallest element.
func (m *MergeIterator) Next() bool {
	// TODO
}

// Value returns the current element.
func (m *MergeIterator) Value() int {
	// TODO
}
