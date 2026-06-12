package pubsub

import "context"

// Message is a pub/sub message.
type Message struct {
	Topic string
	Data  []byte
}

// PubSub manages subscriptions.
type PubSub struct {
	// TODO
}

// New creates a PubSub.
func New() *PubSub {
	// TODO
}

// Subscribe registers a subscriber for a topic pattern.
func (ps *PubSub) Subscribe(ctx context.Context, pattern string) <-chan Message {
	// TODO
}

// Publish sends a message to matching subscribers.
func (ps *PubSub) Publish(msg Message) {
	// TODO
}
