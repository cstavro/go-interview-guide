package main

import (
	"testing"
)

func TestResultMap(t *testing.T) {
	var got int
	parseInt("42").Map(func(n int) int { return n * 2 }).Map(func(n int) int {
		got = n
		return n
	})
	if got != 84 {
		t.Errorf("got %d, want 84", got)
	}
}

func TestResultMapOnError(t *testing.T) {
	var got int
	parseInt("bad").Map(func(n int) int { return n * 2 }).OrElse(func() int {
		return 1
	}).Map(func(n int) int {
		got = n
		return n
	})
	if got != 1 {
		t.Errorf("got %d, want 1", got)
	}
}

func TestResultOrElse(t *testing.T) {
	var got int
	parseInt("bad").OrElse(func() int {
		return 1
	}).Map(func(n int) int {
		got = n
		return n
	})
	if got != 1 {
		t.Errorf("got %d, want 1", got)
	}
}

func TestResultOrElseOnSuccess(t *testing.T) {
	var got int
	parseInt("42").OrElse(func() int {
		return 1
	}).Map(func(n int) int {
		got = n
		return n
	})
	if got != 42 {
		t.Errorf("got %d, want 42", got)
	}
}

func TestResultChain(t *testing.T) {
	var got int
	parseInt("42").
		Map(func(n int) int { return n * 2 }).
		Map(func(n int) int { return n + 10 }).
		Map(func(n int) int {
			got = n
			return n
		})
	if got != 94 {
		t.Errorf("got %d, want 94", got)
	}
}

func TestResultChainWithError(t *testing.T) {
	var got int
	parseInt("not-a-number").
		Map(func(n int) int { return n * 2 }).
		Map(func(n int) int { return n + 10 }).
		OrElse(func() int {
			return 1
		}).
		Map(func(n int) int {
			got = n
			return n
		})
	if got != 1 {
		t.Errorf("got %d, want 1", got)
	}
}


