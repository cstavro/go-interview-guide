package main

import (
    "testing"
)

func TestFindOrder(t *testing.T) {
    tests := []struct {
        numCourses    int
        prerequisites [][]int
    }{
        {2, [][]int{{1, 0}}},
        {4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}},
    }
    for _, tt := range tests {
        order := FindOrder(tt.numCourses, tt.prerequisites)
        if len(order) != tt.numCourses {
            t.Errorf("FindOrder(%d, %v) length = %d, want %d", tt.numCourses, tt.prerequisites, len(order), tt.numCourses)
        }
        pos := make([]int, tt.numCourses)
        for i, c := range order {
            pos[c] = i
        }
        for _, p := range tt.prerequisites {
            if pos[p[0]] <= pos[p[1]] {
                t.Errorf("FindOrder(%d, %v) violates prerequisite %v", tt.numCourses, tt.prerequisites, p)
            }
        }
    }

    cyclic := FindOrder(2, [][]int{{1, 0}, {0, 1}})
    if len(cyclic) != 0 {
        t.Errorf("FindOrder cyclic = %v, want empty", cyclic)
    }
}
