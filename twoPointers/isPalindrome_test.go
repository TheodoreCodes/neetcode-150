package twoPointers

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Problem constraints:
//
//	1 <= s.length <= 1000
//	s is made of printable ASCII characters
func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"example 1", "Was it a car or a cat I saw?", true},
		{"example 2", "tab a cat", false},
		{"single char", "a", true},
		{"only punctuation", ".,!? ", true},
		{"only spaces", "   ", true},
		{"mixed case", "Aa", true},
		{"digits palindrome", "12321", true},
		{"digits not palindrome", "12345", false},
		{"letter and digit differ", "0P", false},
		{"punctuation ignored", "a.b,a", true},
		{"leading and trailing noise", "!!a b a!!", true},
		{"even length", "abba", true},
		{"even length not palindrome", "abca", false},
		{"long palindrome", strings.Repeat("a", 712), true},
		{"long non-palindrome", strings.Repeat("a", 999) + "b", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isPalindrome(tt.s))
		})
	}
}
