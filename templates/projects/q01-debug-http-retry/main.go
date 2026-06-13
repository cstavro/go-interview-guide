package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"runtime"
	"time"
)

// Client makes HTTP requests with retries.
type Client struct {
	MaxRetries int
	BaseDelay  time.Duration
}

func (c *Client) doOnce(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}

func (c *Client) Do(ctx context.Context, url string) (*http.Response, error) {
	var lastErr error
	for i := 0; i < c.MaxRetries; i++ {
		respCh := make(chan *http.Response)
		errCh := make(chan error)

		go func() {
			resp, err := c.doOnce(ctx, url)
			if err != nil {
				errCh <- err
				return
			}
			respCh <- resp
		}()

		select {
		case resp := <-respCh:
			return resp, nil
		case err := <-errCh:
			lastErr = err
		case <-ctx.Done():
			return nil, ctx.Err() // BUG: the goroutine is now leaked
		}
	}
	return nil, fmt.Errorf("after %d retries: %w", c.MaxRetries, lastErr)
}

func main() {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	before := runtime.NumGoroutine()
	client := &Client{MaxRetries: 3, BaseDelay: 10 * time.Millisecond}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	_, err := client.Do(ctx, server.URL)
	if err == nil {
		fmt.Println("expected error")
	}

	time.Sleep(200 * time.Millisecond)
	after := runtime.NumGoroutine()
	fmt.Printf("goroutines: before=%d, after=%d\n", before, after)
}
