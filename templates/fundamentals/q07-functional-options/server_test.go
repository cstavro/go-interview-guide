package server

import (
	"log/slog"
	"testing"
	"time"
)

func TestNewServerDefaults(t *testing.T) {
	s, err := NewServer()
	if err != nil {
		t.Fatalf("NewServer() unexpected error: %v", err)
	}
	if s == nil {
		t.Fatal("NewServer() returned nil Server")
	}
	if s.port <= 0 || s.port > 65535 {
		t.Errorf("port = %d, want a valid non-zero default port", s.port)
	}
	if s.timeout <= 0 {
		t.Errorf("timeout = %v, want a positive default timeout", s.timeout)
	}
	if s.logger == nil {
		t.Error("expected a non-nil default logger")
	}
	if s.tls {
		t.Error("expected TLS to be disabled by default")
	}
}

func TestNewServerWithPort(t *testing.T) {
	s, err := NewServer(WithPort(9000))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s.port != 9000 {
		t.Errorf("port = %d, want 9000", s.port)
	}
}

func TestNewServerWithTimeout(t *testing.T) {
	s, err := NewServer(WithTimeout(5 * time.Second))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s.timeout != 5*time.Second {
		t.Errorf("timeout = %v, want 5s", s.timeout)
	}
}

func TestNewServerWithLogger(t *testing.T) {
	logger := slog.Default()
	s, err := NewServer(WithLogger(logger))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s.logger != logger {
		t.Error("logger was not set")
	}
}

func TestNewServerWithTLS(t *testing.T) {
	s, err := NewServer(WithTLS(true))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !s.tls {
		t.Error("TLS was not enabled")
	}
}

func TestNewServerWithAllOptions(t *testing.T) {
	logger := slog.Default()
	s, err := NewServer(
		WithPort(9000),
		WithTimeout(5*time.Second),
		WithLogger(logger),
		WithTLS(true),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s.port != 9000 {
		t.Errorf("port = %d, want 9000", s.port)
	}
	if s.timeout != 5*time.Second {
		t.Errorf("timeout = %v, want 5s", s.timeout)
	}
	if s.logger != logger {
		t.Error("logger was not set")
	}
	if !s.tls {
		t.Error("TLS was not enabled")
	}
}

func TestNewServerInvalidConfiguration(t *testing.T) {
	cases := []struct {
		name string
		opts []Option
	}{
		{"invalid port", []Option{WithPort(0)}},
		{"negative port", []Option{WithPort(-1)}},
		{"negative timeout", []Option{WithTimeout(-1 * time.Second)}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := NewServer(c.opts...)
			if err == nil {
				t.Error("expected error for invalid configuration, got nil")
			}
		})
	}
}
