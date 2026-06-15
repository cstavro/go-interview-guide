package workerpool

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestPoolProcessesSubmittedJobs(t *testing.T) {
	pool := New(3)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool.Start(ctx)

	var counter atomic.Int64
	const n = 100

	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		if err := pool.Submit(func() {
			defer wg.Done()
			counter.Add(1)
		}); err != nil {
			t.Fatalf("Submit returned unexpected error: %v", err)
		}
	}

	wg.Wait()
	cancel()
	pool.Stop()

	if got := counter.Load(); got != n {
		t.Fatalf("expected %d jobs to complete, got %d", n, got)
	}
}

func TestPoolStopRejectsNewSubmissions(t *testing.T) {
	pool := New(2)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	pool.Start(ctx)

	if err := pool.Submit(func() {}); err != nil {
		t.Fatalf("Submit returned unexpected error before stop: %v", err)
	}

	pool.Stop()

	if err := pool.Submit(func() {}); err == nil {
		t.Fatal("expected Submit to return error after Stop, got nil")
	}
}

// TestPoolStopIsGraceful verifies that Stop waits for in-flight work to finish
// before returning. It does not rely on ctx cancellation; Stop itself is the
// graceful shutdown mechanism.
func TestPoolStopIsGraceful(t *testing.T) {
	pool := New(1)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	pool.Start(ctx)

	started := make(chan struct{})
	done := make(chan struct{})

	pool.Submit(func() {
		close(started)
		time.Sleep(50 * time.Millisecond)
		close(done)
	})

	<-started

	stopFinished := make(chan struct{})
	go func() {
		pool.Stop()
		close(stopFinished)
	}()

	select {
	case <-stopFinished:
		t.Fatal("Stop returned before in-flight job completed")
	case <-time.After(10 * time.Millisecond):
		// Stop is correctly waiting for the job.
	}

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("in-flight job did not complete")
	}

	select {
	case <-stopFinished:
	case <-time.After(time.Second):
		t.Fatal("Stop did not return after job completed")
	}
}

// TestContextCancelAbortsWorkers verifies that ctx cancellation is an external
// abort signal: workers stop accepting new jobs and exit. Unlike Stop, it does
// not need to drain in-flight work gracefully.
func TestContextCancelAbortsWorkers(t *testing.T) {
	pool := New(1)
	ctx, cancel := context.WithCancel(context.Background())
	pool.Start(ctx)

	started := make(chan struct{})
	pool.Submit(func() {
		close(started)
		time.Sleep(100 * time.Millisecond)
	})

	<-started

	cancel()

	stopDone := make(chan struct{})
	go func() {
		pool.Stop()
		close(stopDone)
	}()

	select {
	case <-stopDone:
	case <-time.After(time.Second):
		t.Fatal("workers did not exit after context cancellation")
	}
}

func TestPoolStopIsIdempotent(t *testing.T) {
	pool := New(2)
	ctx, cancel := context.WithCancel(context.Background())
	pool.Start(ctx)

	cancel()
	pool.Stop()
	pool.Stop() // should not panic or deadlock
}

func TestPoolStartStopLifecycle(t *testing.T) {
	pool := New(2)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool.Start(ctx)
	pool.Stop()
}

func TestSubmitBeforeStartReturnsError(t *testing.T) {
	pool := New(2)

	if err := pool.Submit(func() {}); err == nil {
		t.Fatal("expected Submit to return error before Start, got nil")
	}
}
