package workerpool

import "context"

// Job is a unit of work.
type Job func()

// Pool manages a set of workers.
type Pool struct {
	// TODO
}

// New creates a Pool with n workers.
func New(n int) *Pool {
	// TODO
}

// Start begins processing jobs.
func (p *Pool) Start(ctx context.Context) {
	// TODO
}

// Submit adds a job. Returns error if pool is stopped.
func (p *Pool) Submit(j Job) error {
	// TODO
}

// Stop signals shutdown and waits for in-flight jobs.
func (p *Pool) Stop() {
	// TODO
}
