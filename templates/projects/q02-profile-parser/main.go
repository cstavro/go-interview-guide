package main

import (
	"fmt"
	"strings"
)

// LogEntry represents a parsed log line.
type LogEntry struct {
	Level   string
	Message string
}

// ParseLogLine parses a single log line.
// Format: "LEVEL message content here"
func ParseLogLine(line string) (*LogEntry, error) {
	parts := strings.Split(line, " ")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid log line")
	}
	level := strings.TrimSpace(parts[0])
	msg := strings.Join(parts[1:], " ")
	return &LogEntry{Level: level, Message: msg}, nil
}

// ParseLog parses a multi-line log string.
func ParseLog(input string) []LogEntry {
	var entries []LogEntry
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		entry, err := ParseLogLine(line)
		if err != nil {
			continue
		}
		entries = append(entries, *entry)
	}
	return entries
}
