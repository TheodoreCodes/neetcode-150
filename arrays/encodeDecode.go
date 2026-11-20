package arrays

import (
	"fmt"
	"strconv"
)

// fixed-length numerical header (e.g. 004 means the next 4 characters form a string)
const (
	header = "%03d"
	headerLen = 3
)

// encode and decode solve NeetCode "Encode and Decode Strings"
// https://neetcode.io/problems/string-encode-and-decode/question
//
// Design an algorithm to encode a list of strings to a single string, and
// decode that string back to the original list.

// encode turns strs into a single string.
func encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var encoded string
	for _, s := range strs {
		encoded += fmt.Sprintf(header, len(s))
		encoded += s
	}

	return encoded
}

// decode turns a string produced by encode back into the original list.
func decode(s string) []string {
	var decoded []string
	i := 0
	for i < len(s) {
		j := i+headerLen
		length, _ := strconv.Atoi(s[i:j])
		decoded = append(decoded, s[j:j+length])
		i = j+length
	}

	return decoded
}
