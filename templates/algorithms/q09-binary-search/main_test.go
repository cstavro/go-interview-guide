package main

import "testing"

func TestSearch(t *testing.T) {
    nums := []int{-1, 0, 3, 5, 9, 12}
    if got := Search(nums, 9); got != 4 {
        t.Errorf("Search(nums, 9) = %d, want 4", got)
    }
    if got := Search(nums, 2); got != -1 {
        t.Errorf("Search(nums, 2) = %d, want -1", got)
    }
}
