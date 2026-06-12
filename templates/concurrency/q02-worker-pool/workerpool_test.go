package workerpool

import (
	"context"
	"testing"
	"time"
)

func TestPoolGracefulShutdown(t *testing.T) {
	pool := New(3)
	ctx, cancel := context.WithCancel(context.Background())
	pool.Start(ctx)

	for i := 0; i < 10; i++ {
		pool.Submit(func() { time.Sleep(10 * time.Millisecond) })
	}

	cancel()
	pool.Stop()
	// TODO: assert all jobs completed
}
