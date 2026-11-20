package arrays

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Problem constraints:
//
//	0 <= strs.length < 100
//	0 <= strs[i].length < 200
//	strs[i] contains any of the 256 ASCII byte values
const (
	maxStrs   = 99
	maxStrLen = 199
)

func repeatStrs(s string, n int) []string {
	strs := make([]string, n)
	for i := range strs {
		strs[i] = s
	}
	return strs
}

// allByteValueStrs spreads bytes 0-255 over strings within the length limit.
func allByteValueStrs() []string {
	all := make([]byte, 256)
	for i := range all {
		all[i] = byte(i)
	}
	return []string{string(all[:128]), string(all[128:])}
}

func TestEncodeDecode(t *testing.T) {
	tests := []struct {
		name string
		strs []string
	}{
		{"example 1", []string{"neet", "code", "love", "you"}},
		{"example 2", []string{"we", "say", ":", "yes"}},
		{"empty list", []string{}},
		{"single empty string", []string{""}},
		{"multiple empty strings", []string{"", "", ""}},
		{"empty strings mixed in", []string{"a", "", "b", ""}},
		{"single string", []string{"hello"}},
		{"contains delimiter-like chars", []string{"#", "5#", "#5", "a#b#c"}},
		{"contains digits", []string{"12", "3", "456#", "0"}},
		{"looks like encoded output", []string{"4#neet", "4#code"}},
		{"contains commas and colons", []string{"a,b", ":", ",", "::"}},
		{"whitespace and newlines", []string{" ", "\n", "a b", "\t\n"}},
		{"leetcode example 1", []string{"Hello", "World"}},
		{"max length string", []string{strings.Repeat("a", maxStrLen)}},
		{"max strings of max length", repeatStrs(strings.Repeat("z", maxStrLen), maxStrs)},
		{"max strings all empty", repeatStrs("", maxStrs)},
		{"all 256 byte values", allByteValueStrs()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := decode(encode(tt.strs))

			// nil and an empty slice are both valid for an empty list.
			if len(tt.strs) == 0 {
				assert.Empty(t, got)
				return
			}
			assert.Equal(t, tt.strs, got)
		})
	}
}

// An empty list and a list holding one empty string must not collide.
func TestEncodeDistinguishesEmptyListFromEmptyString(t *testing.T) {
	assert.NotEqual(t, encode([]string{}), encode([]string{""}))
}
