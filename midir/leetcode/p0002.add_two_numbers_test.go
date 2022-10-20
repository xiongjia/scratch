package leetcode

import (
	"testing"
)

func Test_p0002(t *testing.T) {
	l := &ListNode{}
	temp := l
	nums := []int{1, 25, 3}
	for _, v := range nums {
		temp.Next = &ListNode{Val: v}
		temp = temp.Next
	}

	n := l.Next
	for n != nil {
		t.Logf("value = %v", n.Val)
		n = n.Next
	}
}
