package arrays

// topKFrequent solves NeetCode "Top K Elements in List"
// https://neetcode.io/problems/top-k-elements-in-list/question
//
// Given an integer array nums and an integer k, return the k most frequent
// elements. The answer is guaranteed to be unique and may be in any order.
//
// Approach: bucket sort by frequency.
// A value can appear at most len(nums) times, so frequencies are bounded and
// we can use the frequency as an index into a slice of buckets instead of
// sorting. This gives O(n) time and O(n) space.
func topKFrequent(nums []int, k int) []int {
	// return early in case k has an invalid value
	if k <= 0 {
		return nil
	}

	// numsFreq maps each number to how many times it appears in nums.
	numsFreq := make(map[int]int)
	// buckets[i] holds all numbers that appear exactly i times.
	// Size is len(nums)+1 because a frequency can range from 1 to len(nums).
	buckets := make([][]int, len(nums)+1)

	// Step 1: count occurrences of every number.
	for _, num := range nums {
		numsFreq[num]++
	}

	// Step 2: place each number into the bucket matching its frequency.
	for num, freq := range numsFreq {
		buckets[freq] = append(buckets[freq], num)
	}

	// Step 3: walk buckets from the highest frequency down to 1, collecting
	// numbers until we have at least k. Since the answer is guaranteed to be
	// unique, the final bucket won't push us past k with ambiguous ties.
	result := make([]int, 0, k)
	for i := len(buckets) - 1; i > 0; i-- {
		result = append(result, buckets[i]...)

		if len(result) >= k {
			return result[:k]
		}
	}

	// Unreachable for valid input (k <= number of distinct elements).
	return nil
}
