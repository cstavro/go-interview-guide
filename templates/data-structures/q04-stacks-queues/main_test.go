package main

import "testing"

func TestIsValid(t *testing.T) {
    tests := []struct {
        s    string
        want bool
    }{
        {"()", true},
        {"()[]{}", true},
        {"(]", false},
        {"([)]", false},
        {"{[]}", true},
        {"", true},
    }
    for _, tt := range tests {
        if got := IsValid(tt.s); got != tt.want {
            t.Errorf("IsValid(%q) = %v, want %v", tt.s, got, tt.want)
        }
    }
}
