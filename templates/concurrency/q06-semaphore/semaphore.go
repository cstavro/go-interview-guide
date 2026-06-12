package semaphore

// Semaphore limits concurrent access to a resource.
type Semaphore struct {
	// TODO
}

// New creates a semaphore with capacity n.
func New(n int) *Semaphore {
	// TODO
}

// Acquire blocks until a slot is available.
func (s *Semaphore) Acquire() {
	// TODO
}

// Release frees a slot.
func (s *Semaphore) Release() {
	// TODO
}
