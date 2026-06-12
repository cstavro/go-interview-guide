package leaderelection

import (
	"context"
	"time"
)

// LeaderElection manages leadership for a service.
type LeaderElection struct {
	// TODO
}

// NewLeaderElection creates a new election manager.
func NewLeaderElection(locker DistributedLock, ttl time.Duration) *LeaderElection {
	// TODO
}

// Run starts the leader election loop.
// If elected, it runs the provided function. On lock loss, it cancels the context.
func (le *LeaderElection) Run(ctx context.Context, onLeader func(ctx context.Context)) {
	// TODO
}

// DistributedLock is the interface for the lock backend.
type DistributedLock interface {
	Acquire(ctx context.Context, id string, ttl time.Duration) (bool, error)
	Renew(ctx context.Context, id string, ttl time.Duration) (bool, error)
	Release(ctx context.Context, id string) error
}
