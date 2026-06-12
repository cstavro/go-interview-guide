package breaker

import (
	"errors"
	"testing"
	"time"
)

func TestBreakerOpens(t *testing.T) {
	b := NewBreaker(3, 1*time.Second)
	for i := 0; i < 3; i++ {
		b.Call(func() error { return errors.New("fail") })
	}
	if b.Call(func() error { return nil }) == nil {
		t.Error("expected breaker to be open")
	}
}
