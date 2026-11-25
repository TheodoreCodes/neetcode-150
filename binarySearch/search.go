package binarySearch

// search solves NeetCode "Binary Search"
// https://neetcode.io/problems/binary-search/question
//
// nums is sorted in ascending order and contains distinct values. Return the
// index of target in nums, or -1 if it is not present. Must run in O(log n).
//
// Constraints:
//
//	1 <= nums.length <= 10000
//	-10000 < nums[i], target < 10000
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)/2

		switch {
		case target == nums[mid]:
			return mid
		case target < nums[mid]:
			high = mid - 1
		case target > nums[mid]:
			low = mid + 1
		}
	}

	return -1
}
