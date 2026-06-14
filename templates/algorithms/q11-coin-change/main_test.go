package main

import "testing"

func TestCoinChange(t *testing.T) {
    if got := CoinChange([]int{1, 2, 5}, 11); got != 3 {
        t.Errorf("CoinChange([1,2,5], 11) = %d, want 3", got)
    }
    if got := CoinChange([]int{2}, 3); got != -1 {
        t.Errorf("CoinChange([2], 3) = %d, want -1", got)
    }
    if got := CoinChange([]int{1}, 0); got != 0 {
        t.Errorf("CoinChange([1], 0) = %d, want 0", got)
    }
}
