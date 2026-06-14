package main

import (
    "reflect"
    "testing"
)

func TestMergeSortedArrays(t *testing.T) {
    tests := []struct {
        a, b []int
        want []int
    }{
        {[]int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
        {[]int{}, []int{1, 2, 3}, []int{1, 2, 3}},
        {[]int{7, 8}, []int{}, []int{7, 8}},
        {[]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 1, 2, 2, 3, 3}},
    }
    for _, tt := range tests {
        got := MergeSortedArrays(tt.a, tt.b)
        if !reflect.DeepEqual(got, tt.want) {
            t.Errorf("MergeSortedArrays(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
        }
    }
}
