package hashring

import "testing"

func TestConsistentHash(t *testing.T) {
	ring := New([]string{"node-a", "node-b", "node-c"}, 150)
	node := ring.GetNode("user:12345")
	if node == "" {
		t.Error("expected a node")
	}

	// Removing a node should only move keys mapped to that node
	ring.RemoveNode("node-b")
	// TODO: assert minimal movement
}
