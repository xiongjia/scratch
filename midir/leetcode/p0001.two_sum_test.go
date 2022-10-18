package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_p0001(t *testing.T) {
	t.Log("P001 Two Sum")

	cases := []struct {
		nums, expected []int
		target         int
	}{
		{[]int{3, 2, 4}, []int{1, 2}, 6},
		{[]int{2, 7, 11, 15}, []int{0, 1}, 9},
	}

	for _, c := range cases {
		s1 := TwoSumSolution1(c.nums, c.target)
		t.Logf("Solution#1: src = %v, target = %v. ret = %v, expected = %v",
			c.nums, c.target, s1, c.expected)
		assert.ElementsMatch(t, s1, c.expected)

		s2 := TwoSumSolution1(c.nums, c.target)
		t.Logf("Solution#2: src = %v, target = %v. ret = %v, expected = %v",
			c.nums, c.target, s2, c.expected)
		assert.ElementsMatch(t, s2, c.expected)
	}
}
