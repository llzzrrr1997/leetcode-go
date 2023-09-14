package problemset

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	m := 0
	for l1 != nil || l2 != nil || m != 0 {
		if l1 != nil {
			m += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			m += l2.Val
			l2 = l2.Next
		}
		cur := &ListNode{
			Val:  m % 10,
			Next: nil,
		}
		m = m / 10
		if head == nil {
			head = cur
		}
		if tail == nil {
			tail = cur
		} else {
			tail.Next = cur
			tail = tail.Next
		}
	}
	return head
}

func addTwoNumbersN(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	m := 0
	for l1 != nil || l2 != nil {
		val := m
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		m = val / 10
		tmp := &ListNode{
			Val: val % 10,
		}
		cur.Next = tmp
		cur = cur.Next
	}
	if m != 0 {
		tmp := &ListNode{
			Val: m,
		}
		cur.Next = tmp
	}
	return head.Next
}
