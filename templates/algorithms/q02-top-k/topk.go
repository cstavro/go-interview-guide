package topk

// TopK tracks the top K frequent items.
type TopK struct {
	// TODO
}

// NewTopK creates a TopK tracker.
func NewTopK(k int) *TopK {
	// TODO
}

// Observe records an item from the stream.
func (tk *TopK) Observe(item string) {
	// TODO
}

// Get returns the top K items and their frequencies.
func (tk *TopK) Get() []Item {
	// TODO
}

// Item represents a ranked element.
type Item struct {
	Value string
	Count int
}
