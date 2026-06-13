package main

import (
	"fmt"
	"strings"
	"testing"
)

func generateLogLines(n int) string {
	var b strings.Builder
	for i := 0; i < n; i++ {
		fmt.Fprintf(&b, "INFO this is a log message with some content %d\n", i)
		fmt.Fprintf(&b, "WARN this is a warning message %d\n", i)
		fmt.Fprintf(&b, "ERROR something went wrong in module %d\n", i)
	}
	return b.String()
}

func BenchmarkParseLog(b *testing.B) {
	input := generateLogLines(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ParseLog(input)
	}
}

func TestParseLog(t *testing.T) {
	input := "INFO hello world\nWARN watch out\n"
	entries := ParseLog(input)
	if len(entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(entries))
	}
	if entries[0].Level != "INFO" || entries[0].Message != "hello world" {
		t.Fatalf("unexpected first entry: %+v", entries[0])
	}
	if entries[1].Level != "WARN" || entries[1].Message != "watch out" {
		t.Fatalf("unexpected second entry: %+v", entries[1])
	}
}
