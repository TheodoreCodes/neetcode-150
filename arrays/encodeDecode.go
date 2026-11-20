package arrays

import (
	"fmt"
	"strconv"
	"strings"
)

// fixed-length numerical header (e.g. 004 means the next 4 characters form a string)
const (
	header    = "%03d"
	headerLen = 3
)

// Encode and Decode solve NeetCode "Encode and Decode Strings"
// https://neetcode.io/problems/string-encode-and-decode/question
//
// Design an algorithm to encode a list of strings to a single string, and
// decode that string back to the original list.

// Encode turns strs into a single string.
func Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var encoded strings.Builder
	for _, s := range strs {
		fmt.Fprintf(&encoded, header, len(s))
		encoded.WriteString(s)
	}

	return encoded.String()
}

// Decode turns a string produced by encode back into the original list.
func Decode(s string) []string {
	var decoded []string
	i := 0
	for i < len(s) {
		j := i + headerLen
		length, _ := strconv.Atoi(s[i:j])
		decoded = append(decoded, s[j:j+length])
		i = j + length
	}

	return decoded
}
