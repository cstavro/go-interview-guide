package main

import (
	"testing"
	"time"
)

func TestConfigEqual(t *testing.T) {
	c1 := Config{
		Name:      "svc",
		Timeout:   5 * time.Second,
		Endpoints: []string{"a", "b"},
	}
	c2 := Config{
		Name:      "svc",
		Timeout:   5 * time.Second,
		Endpoints: []string{"a", "b"},
	}

	// Equal values with different backing arrays should be equal
	if !c1.Equal(c2) {
		t.Error("expected c1 and c2 to be equal")
	}

	// Different slice contents should not be equal
	c3 := Config{
		Name:      "svc",
		Timeout:   5 * time.Second,
		Endpoints: []string{"a", "c"},
	}
	if c1.Equal(c3) {
		t.Error("expected c1 and c3 to be not equal")
	}

	// Different Name should not be equal
	c4 := Config{
		Name:      "other",
		Timeout:   5 * time.Second,
		Endpoints: []string{"a", "b"},
	}
	if c1.Equal(c4) {
		t.Error("expected c1 and c4 to be not equal")
	}

	// Different Timeout should not be equal
	c5 := Config{
		Name:      "svc",
		Timeout:   10 * time.Second,
		Endpoints: []string{"a", "b"},
	}
	if c1.Equal(c5) {
		t.Error("expected c1 and c5 to be not equal")
	}

	// Shared backing array should still be equal
	c6 := Config{
		Name:      "svc",
		Timeout:   5 * time.Second,
		Endpoints: c1.Endpoints,
	}
	if !c1.Equal(c6) {
		t.Error("expected c1 and c6 to be equal")
	}
}
