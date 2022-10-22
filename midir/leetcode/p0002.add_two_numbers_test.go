package leetcode

import (
	"testing"
)

func Test_p0002(t *testing.T) {
	cases := []struct {
		l1, l2, e []int
	}{
		{
			l1: []int{2, 4, 3}, l2: []int{5, 6, 4}, e: []int{7, 0, 8},
		}, {
			l1: []int{0}, l2: []int{0}, e: []int{0},
		},
	}

	for _, c := range cases {
		t.Logf("l1 = %v, l2 = %v, e = %v", c.l1, c.l2, c.e)
	}
}
