package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"runtime"
	"testing"
	"time"
)

func BenchmarkClientGoroutineLeak(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	before := runtime.NumGoroutine()
	client := &Client{MaxRetries: 3, BaseDelay: 10 * time.Millisecond}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
		client.Do(ctx, server.URL)
		cancel()
	}
	b.StopTimer()

	time.Sleep(500 * time.Millisecond)
	after := runtime.NumGoroutine()
	if after > before+5 {
		b.Fatalf("possible goroutine leak: before=%d, after=%d", before, after)
	}
}

func TestClientDoSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := &Client{MaxRetries: 3, BaseDelay: 10 * time.Millisecond}
	ctx := context.Background()
	resp, err := client.Do(ctx, server.URL)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
}
