package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	cases := []struct {
		input []string
		want  string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{}, ""},
		{[]string{"alone"}, "alone"},
		{[]string{"", "b"}, ""},
		{[]string{"prefix", "prefix", "prefix"}, "prefix"},
	}
	for _, c := range cases {
		got := LongestCommonPrefix(c.input)
		if got != c.want {
			t.Errorf("LongestCommonPrefix(%v) = %q, want %q", c.input, got, c.want)
		}
	}
}
