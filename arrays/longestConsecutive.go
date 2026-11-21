package arrays

import "sort"

// longestConsecutive solves NeetCode "Longest Consecutive Sequence"
// https://neetcode.io/problems/longest-consecutive-sequence/question
//
// Given an unsorted integer array nums, return the length of the longest
// sequence of consecutive integers (e.g. 2,3,4,5). The elements need not be
// adjacent in nums. Must run in O(n) time.
//
// Constraints:
//
//	0 <= nums.length <= 100_000
//	-10^9 <= nums[i] <= 10^9
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	buckets := make(map[int]struct{})
	for _, num := range nums {
		buckets[num] = struct{}{}
	}

	longestSeq := 1
	for num := range buckets {
		if _, exists := buckets[num-1]; exists {
			continue
		}
		currSeq := seq(buckets, num, 1)
		if longestSeq < currSeq {
			longestSeq = currSeq
		}
	}

	return longestSeq
}

func seq(nums map[int]struct{}, currNum, seqLen int) int {
	if _, exists := nums[currNum+1]; !exists {
		return seqLen
	}

	return seq(nums, currNum+1, seqLen+1)
}
