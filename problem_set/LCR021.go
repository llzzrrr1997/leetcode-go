package problemset

func removeNthFromEndLCR021(head *ListNode, n int) *ListNode {
	newHead := &ListNode{Next: head}
	pre := newHead
	for i := 0; i < n+1; i++ {
		pre = pre.Next
	}
	cur := newHead
	for pre != nil {
		cur = cur.Next
		pre = pre.Next
	}
	//cur为newHead倒数第n+1个节点，为head倒数第n+1节点
	cur.Next = cur.Next.Next
	return newHead.Next
}
