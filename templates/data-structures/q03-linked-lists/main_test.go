package main

import (
    "reflect"
    "testing"
)

func sliceToList(nums []int) *ListNode {
    dummy := &ListNode{}
    cur := dummy
    for _, n := range nums {
        cur.Next = &ListNode{Val: n}
        cur = cur.Next
    }
    return dummy.Next
}

func listToSlice(head *ListNode) []int {
    var out []int
    for head != nil {
        out = append(out, head.Val)
        head = head.Next
    }
    return out
}

func TestMergeTwoLists(t *testing.T) {
    tests := []struct {
        a, b []int
        want []int
    }{
        {[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
        {[]int{}, []int{}, []int{}},
        {[]int{}, []int{0}, []int{0}},
    }
    for _, tt := range tests {
        got := listToSlice(MergeTwoLists(sliceToList(tt.a), sliceToList(tt.b)))
        if !reflect.DeepEqual(got, tt.want) {
            t.Errorf("MergeTwoLists(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
        }
    }
}
