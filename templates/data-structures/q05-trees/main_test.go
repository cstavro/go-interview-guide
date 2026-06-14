package main

import "testing"

func TestMaxDepth(t *testing.T) {
    root := &TreeNode{
        Val: 3,
        Left: &TreeNode{Val: 9},
        Right: &TreeNode{
            Val:   20,
            Left:  &TreeNode{Val: 15},
            Right: &TreeNode{Val: 7},
        },
    }
    if got := MaxDepth(root); got != 3 {
        t.Errorf("MaxDepth(tree) = %d, want 3", got)
    }
    if got := MaxDepth(nil); got != 0 {
        t.Errorf("MaxDepth(nil) = %d, want 0", got)
    }
}
