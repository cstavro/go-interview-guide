package main

import "testing"

func TestSolveNQueens(t *testing.T) {
    got := SolveNQueens(4)
    if len(got) != 2 {
        t.Errorf("SolveNQueens(4) returned %d solutions, want 2", len(got))
    }
    got1 := SolveNQueens(1)
    if len(got1) != 1 || len(got1[0]) != 1 || got1[0][0] != "Q" {
        t.Errorf("SolveNQueens(1) = %v, want [[Q]]", got1)
    }
}
