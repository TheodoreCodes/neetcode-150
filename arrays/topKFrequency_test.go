package arrays

import (
	"slices"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"example 1", []int{1, 2, 2, 3, 3, 3}, 2, []int{2, 3}},
		{"example 2", []int{7, 7}, 1, []int{7}},
		{"single element", []int{5}, 1, []int{5}},
		{"k equals distinct count", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"negative numbers", []int{-1, -1, -2, -2, -2, 0}, 2, []int{-1, -2}},
		{"k is one", []int{4, 4, 4, 1, 1, 2}, 1, []int{4}},
		{"unsorted input", []int{3, 1, 3, 2, 1, 3, 2, 2, 2}, 2, []int{2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.nums, tt.k)

			// The problem allows the answer in any order, so compare sorted.
			slices.Sort(got)
			slices.Sort(tt.want)

			if !slices.Equal(got, tt.want) {
				t.Errorf("topKFrequent(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
