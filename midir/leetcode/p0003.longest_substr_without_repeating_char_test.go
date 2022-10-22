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
	}

	for _, c := range cases {
		r := LengthOfLongestSubstringSolution1(c.input)
		t.Logf("input = %s, expected = %d; r = %d\n", c.input, c.expected, r)
		assert.Equal(t, r, c.expected)
	}
}
