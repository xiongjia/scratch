package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_p0003(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{input: "abcabcbb", expected: 3},
		{input: "bbbbb", expected: 1},
		{input: "pwwkew", expected: 3},
		{input: "", expected: 0},
	}

	for _, c := range cases {
		r := LengthOfLongestSubstringSolution1(c.input)
		t.Logf("Solution1: I = %s, E = %d; R = %d\n", c.input, c.expected, r)
		assert.Equal(t, r, c.expected)
	}
}
