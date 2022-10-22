package leetcode

func AddTwoNumbersSolution1(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{Val: 0}
	current := root
	n1 := 0
	n2 := 0
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		} else {
			n1 = 0
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		} else {
			n2 = 0
		}

		v := n1 + n2 + carry
		if v >= 10 {
			current.Next = &ListNode{Val: v % 10}
			carry = v / 10
		} else {
			current.Next = &ListNode{Val: v}
			carry = 0
		}
		current = current.Next
	}
	return root.Next
}
