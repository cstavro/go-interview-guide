package main

import (
    "reflect"
    "sort"
    "testing"
)

func TestTopKFrequent(t *testing.T) {
    got := TopKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
    sort.Ints(got)
    want := []int{1, 2}
    if !reflect.DeepEqual(got, want) {
        t.Errorf("TopKFrequent([1,1,1,2,2,3], 2) = %v, want %v", got, want)
    }
}
