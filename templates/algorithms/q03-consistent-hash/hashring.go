package hashring

// Ring is a consistent hash ring.
type Ring struct {
	// TODO
}

// New creates a ring with virtual nodes per physical node.
func New(nodes []string, virtualNodes int) *Ring {
	// TODO
}

// AddNode adds a node to the ring.
func (r *Ring) AddNode(node string) {
	// TODO
}

// RemoveNode removes a node from the ring.
func (r *Ring) RemoveNode(node string) {
	// TODO
}

// GetNode returns the node responsible for key.
func (r *Ring) GetNode(key string) string {
	// TODO
}
