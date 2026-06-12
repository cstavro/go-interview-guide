package main

import (
	"context"
	"fmt"
	"time"
)

func callService(ctx context.Context, name string) (string, error) {
	// Simulated: takes 100ms normally
	select {
	case <-time.After(100 * time.Millisecond):
		return name + "-ok", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func callAll(ctx context.Context) (map[string]string, error) {
	// TODO: call A, B, C concurrently with 500ms deadline
	// return all results or an error on first failure
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	results, err := callAll(ctx)
	fmt.Println(results, err)
}
