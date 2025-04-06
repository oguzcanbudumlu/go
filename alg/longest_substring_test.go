package alg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestSubstringWithoutRepetition(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{{"abcabcbb", 3}, {"bbbbb", 1}, {"pwwkew", 3}}

	for _, tt := range tests {
		result := longestSubstring(tt.input)
		assert.Equal(t, tt.output, result)
	}
}

func longestSubstring(s string) int {
	lastSeen := make(map[rune]int)
	left := 0
	maxLen := 0

	for right, char := range s {

		if prevIdx, ok := lastSeen[char]; ok {
			left = prevIdx + 1
		}

		lastSeen[char] = right

		if currentLen := right - left + 1; currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}
