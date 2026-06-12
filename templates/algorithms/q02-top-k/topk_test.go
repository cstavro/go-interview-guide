package topk

import "testing"

func TestTopK(t *testing.T) {
	tk := NewTopK(2)
	for _, v := range []string{"a", "b", "a", "c", "b", "a"} {
		tk.Observe(v)
	}
	result := tk.Get()
	if len(result) != 2 {
		t.Fatalf("expected 2 results, got %d", len(result))
	}
	// TODO: assert correct ordering
}
