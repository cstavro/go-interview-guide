package main

import "testing"

func TestLongestSubstring(t *testing.T) {
    tests := []struct {
        s    string
        want int
    }{
        {"abcabcbb", 3},
        {"bbbbb", 1},
        {"pwwkew", 3},
        {"", 0},
        {"abcdef", 6},
    }
    for _, tt := range tests {
        if got := LongestSubstring(tt.s); got != tt.want {
            t.Errorf("LongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
        }
    }
}
