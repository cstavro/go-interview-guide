package main

import (
	"strings"
	"testing"
)

func TestConfigValidation(t *testing.T) {
	tests := []struct {
		name    string
		cfg     Config
		wantErr string
	}{
		{
			name:    "zero port",
			cfg:     Config{Port: 0, DBHost: "localhost", DBPass: "secret", Workers: 4, LogLevel: "info"},
			wantErr: "port",
		},
		{
			name:    "empty host",
			cfg:     Config{Port: 8080, DBHost: "", DBPass: "secret", Workers: 4, LogLevel: "info"},
			wantErr: "host",
		},
		{
			name:    "non-positive workers",
			cfg:     Config{Port: 8080, DBHost: "localhost", DBPass: "secret", Workers: 0, LogLevel: "info"},
			wantErr: "workers",
		},
		{
			name:    "valid config",
			cfg:     Config{Port: 8080, DBHost: "localhost", DBPass: "secret", Workers: 4, LogLevel: "info"},
			wantErr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.Validate()
			if tt.wantErr != "" {
				if err == nil || !strings.Contains(strings.ToLower(err.Error()), tt.wantErr) {
					t.Fatalf("expected error containing %q, got %v", tt.wantErr, err)
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}
		})
	}
}

func TestConfigStringRedactsPassword(t *testing.T) {
	cfg := Config{Port: 8080, DBHost: "localhost", DBPass: "super-secret", Workers: 4, LogLevel: "info"}
	s := cfg.String()
	if strings.Contains(s, "super-secret") {
		t.Fatalf("String() must redact DBPass, got: %s", s)
	}
	if !strings.Contains(s, "localhost") {
		t.Fatalf("String() should include non-sensitive fields, got: %s", s)
	}
}
