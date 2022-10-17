package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_p0001(t *testing.T) {
	t.Log("P001 Two Sum")

	cases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}
	for _, c := range cases {
		r := TwoSumSolution1(c.nums, c.target)
		t.Logf("Source = %v, Target = %v. Result = %v, Expected = %v",
			c.nums, c.target, r, c.expected)
		assert.ElementsMatch(t, r, c.expected)
	}
}
