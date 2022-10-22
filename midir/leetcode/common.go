package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeListNodes(values []int) *ListNode {
	r := &ListNode{}
	n := r
	for _, v := range values {
		n.Next = &ListNode{Val: v}
		n = n.Next
	}
	return r.Next
}
