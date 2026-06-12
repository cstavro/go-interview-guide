package main

import (
	"encoding/json"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		want   string
	}{
		{"present", `{ "name": "Alice", "age": 30 }`, "present"},
		{"null", `{ "name": "Bob", "age": null }`, "null"},
		{"missing", `{ "name": "Charlie" }`, "missing"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var p Person
			if err := json.Unmarshal([]byte(c.input), &p); err != nil {
				t.Fatal(err)
			}
			// TODO: assert state
		})
	}
}
