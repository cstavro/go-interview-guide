package main

import (
	"testing"
	"time"
)

func TestSlidingWindow(t *testing.T) {
	rl := NewSlidingWindowRateLimiter(60*time.Second, 3)
	user := "alice"

	if !rl.Allow(user, time.Unix(0, 0)) {
		t.Error("expected first request allowed")
	}
	if !rl.Allow(user, time.Unix(10, 0)) {
		t.Error("expected second request allowed")
	}
	if !rl.Allow(user, time.Unix(20, 0)) {
		t.Error("expected third request allowed")
	}
	if rl.Allow(user, time.Unix(30, 0)) {
		t.Error("expected fourth request blocked")
	}
	if !rl.Allow(user, time.Unix(70, 0)) {
		t.Error("expected request after window slid allowed")
	}
}
