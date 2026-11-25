package binarySearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Problem constraints:
//
//	1 <= nums.length <= 10000
//	-10000 < nums[i], target < 10000
//	nums is sorted ascending with distinct values
func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"example 1", []int{-1, 0, 2, 4, 6, 8}, 4, 3},
		{"example 2", []int{-1, 0, 2, 4, 6, 8}, 3, -1},
		{"single element found", []int{5}, 5, 0},
		{"single element not found", []int{5}, 3, -1},
		{"first element", []int{1, 3, 5, 7}, 1, 0},
		{"last element", []int{1, 3, 5, 7}, 7, 3},
		{"middle, odd length", []int{1, 3, 5, 7, 9}, 5, 2},
		{"two elements, first", []int{1, 2}, 1, 0},
		{"two elements, second", []int{1, 2}, 2, 1},
		{"smaller than all", []int{1, 3, 5}, -5, -1},
		{"larger than all", []int{1, 3, 5}, 9, -1},
		{"gap between elements", []int{1, 3, 5, 7}, 4, -1},
		{"negatives", []int{-9, -7, -3, -1}, -3, 2},
		{"extreme values", []int{-9999, 0, 9999}, 9999, 2},
		{"extreme low", []int{-9999, 0, 9999}, -9999, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, search(tt.nums, tt.target))
		})
	}
}

func TestSearchMaxLength(t *testing.T) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i - 5000 // -5000..4999, within (-10000, 10000)
	}
	for _, i := range []int{0, 1, 4999, 5000, 9998, 9999} {
		assert.Equal(t, i, search(nums, nums[i]))
	}
	assert.Equal(t, -1, search(nums, 5000))
	assert.Equal(t, -1, search(nums, -5001))
}
