package main

import (
    "reflect"
    "testing"
)

func TestMergeSort(t *testing.T) {
    tests := []struct {
        in   []int
        want []int
    }{
        {[]int{5, 2, 3, 1, 4}, []int{1, 2, 3, 4, 5}},
        {[]int{}, []int{}},
        {[]int{1, 1, 1}, []int{1, 1, 1}},
        {[]int{9, -3, 0, 7}, []int{-3, 0, 7, 9}},
    }
    for _, tt := range tests {
        got := MergeSort(tt.in)
        if !reflect.DeepEqual(got, tt.want) {
            t.Errorf("MergeSort(%v) = %v, want %v", tt.in, got, tt.want)
        }
    }
}
