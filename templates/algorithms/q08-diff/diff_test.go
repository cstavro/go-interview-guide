package diff

import (
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
	a := []string{"a", "c", "d", "f"}
	b := []string{"b", "c", "e", "f"}

	result := Diff(a, b)
	if !reflect.DeepEqual(result.Added, []string{"b", "e"}) {
		t.Errorf("Added = %v", result.Added)
	}
	if !reflect.DeepEqual(result.Removed, []string{"a", "d"}) {
		t.Errorf("Removed = %v", result.Removed)
	}
	if !reflect.DeepEqual(result.Common, []string{"c", "f"}) {
		t.Errorf("Common = %v", result.Common)
	}
}
