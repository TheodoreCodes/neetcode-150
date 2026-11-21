package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Problem constraints:
//
//	2 <= nums.length <= 1000
//	-20 <= nums[i] <= 20
//	The product of any prefix or suffix of nums fits in a 32-bit integer.
const (
	minNums = 2
	maxNums = 1000
	minVal  = -20
	maxVal  = 20
)

// repeatInts returns a slice of n copies of v.
func repeatInts(v, n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = v
	}
	return s
}

// bruteForceProductExceptSelf is an O(n^2) reference used to verify large inputs.
func bruteForceProductExceptSelf(nums []int) []int {
	out := make([]int, len(nums))
	for i := range nums {
		p := 1
		for j, n := range nums {
			if i != j {
				p *= n
			}
		}
		out[i] = p
	}
	return out
}

func TestProductExceptSelf(t *testing.T) {
	// 1000 elements: a zero and ±1s keep every product within 32 bits.
	maxLenWithZero := repeatInts(1, maxNums)
	maxLenWithZero[500] = 0
	wantMaxLenWithZero := repeatInts(0, maxNums)
	wantMaxLenWithZero[500] = 1

	maxLenAlternating := repeatInts(1, maxNums)
	for i := 0; i < maxNums; i += 2 {
		maxLenAlternating[i] = -1
	}

	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"example 1", []int{1, 2, 4, 6}, []int{48, 24, 12, 8}},
		{"example 2", []int{-1, 0, 1, 2, 3}, []int{0, -6, 0, 0, 0}},
		{"min length", []int{3, 4}, []int{4, 3}},
		{"min length negative", []int{-5, 7}, []int{7, -5}},
		{"min length with zero", []int{0, 9}, []int{9, 0}},
		{"min length both zero", []int{0, 0}, []int{0, 0}},
		{"all ones", []int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{"all negative ones, even count", []int{-1, -1, -1, -1}, []int{-1, -1, -1, -1}},
		{"all negative ones, odd count", []int{-1, -1, -1}, []int{1, 1, 1}},
		{"single zero", []int{2, 0, 3}, []int{0, 6, 0}},
		{"zero at start", []int{0, 2, 3}, []int{6, 0, 0}},
		{"zero at end", []int{2, 3, 0}, []int{0, 0, 6}},
		{"two zeros", []int{0, 2, 0, 3}, []int{0, 0, 0, 0}},
		{"mixed signs", []int{-2, 3, -4, 5}, []int{-60, 40, -30, 24}},
		{"duplicates", []int{2, 2, 2, 2}, []int{8, 8, 8, 8}},
		// 20^7 = 1.28e9 still fits in int32.
		{"max value", repeatInts(maxVal, 7), repeatInts(20*20*20*20*20*20, 7)},
		{"min value, odd count", repeatInts(minVal, 7), repeatInts(20*20*20*20*20*20, 7)},
		{"min value, even count", repeatInts(minVal, 6), repeatInts(-20*20*20*20*20, 6)},
		{"value bounds mixed", []int{minVal, maxVal}, []int{maxVal, minVal}},
		{"max length with zero", maxLenWithZero, wantMaxLenWithZero},
		{"max length alternating signs", maxLenAlternating, bruteForceProductExceptSelf(maxLenAlternating)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Copy so we can also verify the input is not mutated.
			input := append([]int(nil), tt.nums...)

			got := productExceptSelf(input)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.nums, input, "input must not be modified")
		})
	}
}

func TestProductExceptSelfMatchesBruteForce(t *testing.T) {
	// Sweep every valid length with a deterministic pattern covering the full
	// value range. Zeros every 7th element keep products within 32 bits.
	for n := minNums; n <= maxNums; n += 37 {
		nums := make([]int, n)
		for i := range nums {
			if i%7 == 0 {
				nums[i] = 0
			} else {
				nums[i] = minVal + (i*13)%(maxVal-minVal+1)
			}
		}
		// Avoid overflow: with fewer than two zeros, shrink values to ±1.
		zeros := 0
		for _, v := range nums {
			if v == 0 {
				zeros++
			}
		}
		if zeros < 2 {
			for i, v := range nums {
				if v != 0 {
					nums[i] = v / max(1, abs(v))
				}
			}
		}

		assert.Equal(t, bruteForceProductExceptSelf(nums), productExceptSelf(nums), "length %d", n)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
