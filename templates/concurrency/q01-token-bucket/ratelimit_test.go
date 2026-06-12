package ratelimit

import (
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	bucket := NewTokenBucket(5, 1*time.Second)
	if !bucket.AllowN(5) {
		t.Error("expected burst of 5 to be allowed")
	}
	if bucket.Allow() {
		t.Error("expected throttled after burst")
	}
}
