package main

import (
	"testing"
	"time"
)

func TestCursorRoundTrip(t *testing.T) {
	now := time.Now().UnixNano()
	c := Cursor{Timestamp: now, ID: "item-123"}
	enc := EncodeCursor(c)
	dec, err := DecodeCursor(enc)
	if err != nil {
		t.Fatal(err)
	}
	if dec.Timestamp != c.Timestamp {
		t.Fatalf("timestamp mismatch: got %d, want %d", dec.Timestamp, c.Timestamp)
	}
	if dec.ID != c.ID {
		t.Fatalf("id mismatch: got %s, want %s", dec.ID, c.ID)
	}
}

func TestCursorLargeTimestamp(t *testing.T) {
	// This test should fail before the fix due to float64 precision loss.
	c := Cursor{Timestamp: 1752345678901234567, ID: "edge-case"}
	enc := EncodeCursor(c)
	dec, err := DecodeCursor(enc)
	if err != nil {
		t.Fatal(err)
	}
	if dec.Timestamp != c.Timestamp {
		t.Fatalf("large timestamp mismatch: got %d, want %d", dec.Timestamp, c.Timestamp)
	}
}
