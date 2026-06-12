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
}
