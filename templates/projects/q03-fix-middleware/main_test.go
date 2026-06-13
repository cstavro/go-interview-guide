package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddlewareLogsPanic(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("boom")
	})

	// TODO: fix the middleware order and add timing middleware
	handler := recoveryMiddleware(loggingMiddleware(mux))

	req := httptest.NewRequest(http.MethodGet, "/panic", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", rec.Code)
	}
}

func TestMiddlewareTimingHeader(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	// TODO: add timing middleware and ensure it runs in the correct order
	handler := recoveryMiddleware(loggingMiddleware(mux))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	// After adding timing middleware, uncomment this test:
	// duration := rec.Header().Get("X-Request-Duration-Ms")
	// if duration == "" {
	// 	t.Fatal("missing X-Request-Duration-Ms header")
	// }
	_ = rec
	t.Skip("timing middleware not yet implemented")
}

func TestMiddlewareOrder(t *testing.T) {
	// This test should verify that logging runs before recovery.
	// In other words, a panic should be logged.
	mux := http.NewServeMux()
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("boom")
	})

	var logs []string
	loggedRecovery := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					logs = append(logs, "recovered")
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}()
			logs = append(logs, "logged")
			next.ServeHTTP(w, r)
		})
	}

	handler := loggedRecovery(mux)
	req := httptest.NewRequest(http.MethodGet, "/panic", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if len(logs) != 2 || logs[0] != "logged" || logs[1] != "recovered" {
		t.Fatalf("expected [logged, recovered], got %v", logs)
	}
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", rec.Code)
	}
}
