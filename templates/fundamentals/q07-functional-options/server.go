package server

import (
	"log/slog"
	"time"
)

type Server struct {
	port    int
	timeout time.Duration
	logger  *slog.Logger
	tls     bool
}

// Option configures a Server.
type Option func(*Server) error

// NewServer creates a Server with the given options.
func NewServer(opts ...Option) (*Server, error) {
	// TODO: implement with defaults and validation
	return &Server{}, nil
}

// WithPort sets the server's port.
func WithPort(port int) Option {
	// TODO
	return func(*Server) error { return nil }
}

// WithTimeout sets the server's request timeout.
func WithTimeout(timeout time.Duration) Option {
	// TODO
	return func(*Server) error { return nil }
}

// WithLogger sets the server's logger.
func WithLogger(logger *slog.Logger) Option {
	// TODO
	return func(*Server) error { return nil }
}

// WithTLS enables or disables TLS.
func WithTLS(tls bool) Option {
	// TODO
	return func(*Server) error { return nil }
}
