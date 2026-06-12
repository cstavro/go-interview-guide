package middleware

import (
	"net/http"
	"time"
)

// InstrumentedHandler wraps an HTTP handler with logging, metrics, and tracing.
func InstrumentedHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// TODO: wrap ResponseWriter to capture status code

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		// TODO: emit structured log, metric, and trace span
	})
}
