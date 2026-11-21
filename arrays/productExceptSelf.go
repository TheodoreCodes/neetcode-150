package arrays

// productExceptSelf solves NeetCode "Products of Array Except Self"
// https://neetcode.io/problems/products-of-array-discluding-self/question
//
// Given an integer array nums, return an array output where output[i] is the
// product of all the elements of nums except nums[i]. Must run in O(n) time
// and without using division.
//
// Constraints:
//
//	2 <= nums.length <= 1000
//	-20 <= nums[i] <= 20
//	The product of any prefix or suffix of nums fits in a 32-bit integer.
func productExceptSelf(nums []int) []int {
	prefixProducts := make([]int, len(nums))
	prefixProducts[0] = 1

	for i := 1; i < len(nums); i++ {
		prefixProducts[i] = prefixProducts[i-1] * nums[i-1]
	}

	suffixProducts := make([]int, len(nums))
	suffixProducts[len(nums)-1] = 1

	for i := len(nums) - 2; i >= 0; i-- {
		suffixProducts[i] = suffixProducts[i+1] * nums[i+1]
	}

	result := make([]int, len(nums))
	for i := range nums {
		result[i] = prefixProducts[i] * suffixProducts[i]
	}

	return result
}
