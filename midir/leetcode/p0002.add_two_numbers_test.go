package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_p0002(t *testing.T) {
	cases := []struct {
		l1, l2, e []int
	}{
		{
			l1: []int{2, 4, 3}, l2: []int{5, 6, 4}, e: []int{7, 0, 8},
		}, {
			l1: []int{0}, l2: []int{0}, e: []int{0},
		}, {
			l1: []int{9, 9, 9, 9, 9, 9, 9},
			l2: []int{9, 9, 9, 9},
			e:  []int{8, 9, 9, 9, 0, 0, 0, 1},
		}, {
			l1: []int{9, 9, 9, 9}, l2: []int{1}, e: []int{0, 0, 0, 0, 1},
		}, {
			l1: []int{}, l2: []int{}, e: []int{},
		},
	}

	for _, c := range cases {
		output := MakeArray(AddTwoNumbersSolution1(MakeListNodes(c.l1), MakeListNodes(c.l2)))
		t.Logf("l1 = %v, l2 = %v, e = %v; o = %v\n", c.l1, c.l2, c.e, output)
		assert.Equal(t, output, c.e)
	}
}
