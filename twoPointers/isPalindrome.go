package twoPointers

import (
	"regexp"
	"strings"
)

// isPalindrome solves NeetCode "Is Palindrome"
// https://neetcode.io/problems/is-palindrome/question
//
// Return true if s reads the same forwards and backwards, considering only
// alphanumeric characters and ignoring case.
//
// Constraints:
//
//	1 <= s.length <= 1000
//	s is made of printable ASCII characters

var nonAlphaRegex = regexp.MustCompile("[^a-zA-Z0-9]+")

func isPalindrome(s string) bool {
	s = strings.ToLower(nonAlphaRegex.ReplaceAllString(s, ""))
	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
