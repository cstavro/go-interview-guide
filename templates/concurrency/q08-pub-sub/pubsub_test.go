package pubsub

import (
	"context"
	"testing"
	"time"
)

func TestPubSub(t *testing.T) {
	ps := New()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ch := ps.Subscribe(ctx, "orders.*")
	ps.Publish(Message{Topic: "orders.created", Data: []byte("hello")})

	select {
	case msg := <-ch:
		if msg.Topic != "orders.created" {
			t.Errorf("got %s, want orders.created", msg.Topic)
		}
	case <-ctx.Done():
		t.Fatal("timeout")
	}
}
