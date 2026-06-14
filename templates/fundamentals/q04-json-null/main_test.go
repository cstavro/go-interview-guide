package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUnmarshalRoundTrip(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"present", `{"name":"Alice","age":30}`, `{"name":"Alice","age":30}`},
		{"null", `{"name":"Bob","age":null}`, `{"name":"Bob","age":null}`},
		{"missing", `{"name":"Charlie"}`, `{"name":"Charlie"}`},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var p Person
			if err := json.Unmarshal([]byte(c.input), &p); err != nil {
				t.Fatal(err)
			}
			got, err := json.Marshal(p)
			if err != nil {
				t.Fatal(err)
			}
			var gotMap, wantMap map[string]any
			if err := json.Unmarshal(got, &gotMap); err != nil {
				t.Fatal(err)
			}
			if err := json.Unmarshal([]byte(c.want), &wantMap); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("marshal = %s, want %s", got, c.want)
			}
		})
	}
}
