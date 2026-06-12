package interval

import "testing"

func TestOverlapping(t *testing.T) {
	t := NewTree()
	t.Insert(Interval{"a", time(1), time(5)})
	t.Insert(Interval{"b", time(3), time(7)})
	t.Insert(Interval{"c", time(8), time(10)})

	result := t.Overlapping(Interval{"q", time(2), time(4)})
	if len(result) != 2 {
		t.Fatalf("expected 2 overlaps, got %d", len(result))
	}
}
