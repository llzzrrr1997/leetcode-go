package top100

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	mod := 0
	for l1 != nil || l2 != nil {
		t := mod
		if l1 != nil {
			t += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			t += l2.Val
			l2 = l2.Next
		}
		v := t % 10
		mod = t / 10
		cur.Next = &ListNode{
			Val: v,
		}
		cur = cur.Next
	}
	if mod > 0 {
		cur.Next = &ListNode{
			Val: mod,
		}
	}

	return head.Next
}
