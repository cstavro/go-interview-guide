package main

import (
	"reflect"
	"testing"
)

func TestAppendUnique(t *testing.T) {
	original := []int{1, 2, 3}
	result := AppendUnique(original, []int{4, 5, 6, 7, 8})

	want := []int{1, 2, 3, 4, 5, 6, 7, 8}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("AppendUnique() = %v, want %v", result, want)
	}

	// Ensure original was not modified
	if !reflect.DeepEqual(original, []int{1, 2, 3}) {
		t.Errorf("original slice was modified: %v", original)
	}
}
