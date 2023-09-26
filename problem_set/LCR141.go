package problemset

func trainningPlanLCR141(head *ListNode) *ListNode {
	//迭代解决链表反转
	var pre, cur, next *ListNode
	pre, cur, next = nil, head, head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
