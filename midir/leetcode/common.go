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

func MakeArray(l *ListNode) []int {
	r := make([]int, 0)
	for l != nil {
		r = append(r, l.Val)
		l = l.Next
	}
	return r
}
